package gorotine

import (
	"DownloadNew/client"
	"DownloadNew/data"
	"DownloadNew/downlog"
	"DownloadNew/message"
	"DownloadNew/mongo"
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	hst "github.com/yottachain/YTHost/client"
	"log"
	_ "net/http/pprof"
	"sync/atomic"
	"time"
)

var C data.CountForAll
var F data.YTFile
var M data.YTFileMap
var S data.SendMsgGo
var R data.GetCount
var G data.GetTokenGo
var T data.NodeOneData
var I int64

//var ChanFile = make(chan data.YTFile, 20000)
//var ChanShard = make(chan data.YTShard,20000)
var MakeGetTokenMsg = make(chan data.ShardAndNodeId, 1000000)
var ChanSendMsg = make(chan data.SendChan, 200000)

//拨号和获取token
func DownLoadByAdress(n data.ShardAndNodeId) {
	G.ToGetTokenGoLive()
	C.ToShardCount()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	G.ToGetByGoConn()
	clt, err := client.Hst.ClientStore().GetByAddrString(ctx, client.N.NodeData[n.NodeId].NodeID, client.N.NodeData[n.NodeId].Addrs)
	if err != nil {
		G.ToGetByGoConnDie()
		R.ToAdderr()
	} else {
		R.ToConCount()
		var getToken message.NodeCapacityRequest
		var resGetToken message.NodeCapacityResponse
		getToken.RequestMsgID = message.MsgIDDownloadShardRequest.Value()
		getTokenData, _ := proto.Marshal(&getToken)
		ctxto, cancels := context.WithTimeout(context.Background(), time.Second*5)
		defer cancels()
		G.ToGetTokenGoConn()
		tok, err := clt.SendMsg(ctxto, message.MsgIDNodeCapacityRequest.Value(), getTokenData)
		if err != nil {
			log.Println("通信错误", err, n.NodeId, client.N.NodeData[n.NodeId].Addrs)
			//G.ToGetTokenGoConnDie()
			R.ToGetTokenFail()
			ReTokenAgain(n)
			T.NoGetTokenTcpErrNode(n.NodeId)
		} else {
			proto.Unmarshal(tok[2:], &resGetToken)
			if !resGetToken.Writable {
				log.Println("获取token失败", resGetToken.Writable, n.NodeId, client.N.NodeData[n.NodeId].Addrs)
				//G.ToGetTokenGoConnDie()
				R.ToGetTokenFail()
				RetryToken(clt, n)
				T.NoGetTokenErrNode(n.NodeId)
			} else {
				T.NoSuccess(n.NodeId)
				sn := data.SendChan{Ns: n, Clt: clt, Token: resGetToken.AllocId}
				ChanSendMsg <- sn
				C.ToChanSendMsgCount()
				G.ToGetTokenGoFish()
				R.ToGetTokenSuccess()
				//fmt.Println("生产数据",n.NodeId,len(ChanSendMsg),resGetToken)
			}
		}
	}
}

//下载分片数据
func DownLoadSendMsg(n data.SendChan) {
	R.ToShardCountTotal()
	sendStartTime := time.Now()
	S.ToSendMsgGoLive()
	var msg message.DownloadShardRequest
	ctxs, cancels := context.WithTimeout(context.Background(), time.Second*5)
	defer cancels()
	msg.VHF = n.Ns.ShardVHF
	msg.AllocId = n.Token
	buf, err := proto.Marshal(&msg)
	if err != nil {
		panic(err)
	}
	S.ToSendMsgGoConn()
	msgg, err := n.Clt.SendMsg(ctxs, message.MsgIDDownloadShardRequest.Value(), buf)
	sendEndTime := time.Now()
	timeSend := sendEndTime.Sub(sendStartTime)
	if err != nil {
		//S.ToSendMsgGoConnDie()
		//R.ToSenderr()
		log.Printf("sendmsg:err=%v node:id=%v shard:id=%v\n", err, n.Ns.NodeId, n.Ns.ShardVHF)
		RetryDownLoadSendMsg(n)
	} else if len(msgg) > 22 {
		R.ToSuccess()
		S.ToSendMsgGoFish()
		fmt.Printf("sendmsg:success=%v node:id=%v shard:id=%v\n", "success", n.Ns.NodeId, n.Ns.ShardVHF)
		R.ToCount()
		R.ToUsed(timeSend.Milliseconds())
	} else {
		//S.ToSendMsgGoConnDie()
		R.ToSendfail()
		log.Printf("sendmsg:kong=%v node:id=%v shard:id=%v\n", msgg, n.Ns.NodeId, n.Ns.ShardVHF)
		RetryDownLoadSendMsg(n)
	}
}

//获取token失败后重试
func RetryToken(clt *hst.YTHostClient, n data.ShardAndNodeId) {
	var b int
	//go func() {
	//	http.ListenAndServe("0.0.0.0:8899", nil)
	//}()
	G.ToGetTokenRe()
	var getToken message.NodeCapacityRequest
	var resGetToken message.NodeCapacityResponse
	getTokenData, _ := proto.Marshal(&getToken)
	//i:=0;i<10;i++
	for i := 0; i < R.Notice; i++ {
		R.ToTokenRe()
		T.NoRetryToken(n.NodeId)
		rectxto, recancels := context.WithTimeout(context.Background(), time.Second*5)
		tok, err := clt.SendMsg(rectxto, message.MsgIDNodeCapacityRequest.Value(), getTokenData)
		recancels()
		if err != nil {
			log.Println("通信错误 retry", err, n.NodeId, client.N.NodeData[n.NodeId].Addrs)
		} else {
			proto.Unmarshal(tok[2:], &resGetToken)
			if !resGetToken.Writable {
				log.Println("获取token失败 retry", resGetToken.Writable, n.NodeId, client.N.NodeData[n.NodeId].Addrs)
			} else {
				sn := data.SendChan{Ns: n, Clt: clt, Token: resGetToken.AllocId}
				ChanSendMsg <- sn
				C.ToChanSendMsgCount()
				G.ToGetTokenGoFishRe()
				R.ToGetTokenSuccess()
				//fmt.Println("生产数据retry",n.NodeId,len(ChanSendMsg),resGetToken)
				b = 1
				break
			}
		}
	}
	if b != 1 {
		G.ToGetTokenReDie()
	}
}

//获取token通信失败后重试
func ReTokenAgain(n data.ShardAndNodeId) {
	G.ToGetTokenRe()
	var a int
	for i := 0; i < R.Notice; i++ {
		R.ToTokenRe()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		clt, err := client.Hst.ClientStore().GetByAddrString(ctx, client.N.NodeData[n.NodeId].NodeID, client.N.NodeData[n.NodeId].Addrs)
		cancel()
		if err != nil {
		} else {
			var getToken message.NodeCapacityRequest
			var resGetToken message.NodeCapacityResponse
			getToken.RequestMsgID = message.MsgIDDownloadShardRequest.Value()
			getTokenData, _ := proto.Marshal(&getToken)
			ctxto, cancels := context.WithTimeout(context.Background(), time.Second*5)
			tok, err := clt.SendMsg(ctxto, message.MsgIDNodeCapacityRequest.Value(), getTokenData)
			cancels()
			if err != nil {
				log.Println("通信错误 retry-two", err, n.NodeId, client.N.NodeData[n.NodeId].Addrs)
			} else {
				proto.Unmarshal(tok[2:], &resGetToken)
				if !resGetToken.Writable {
					log.Println("获取token失败 retry-two", resGetToken.Writable, n.NodeId, client.N.NodeData[n.NodeId].Addrs)
				} else {
					sn := data.SendChan{Ns: n, Clt: clt, Token: resGetToken.AllocId}
					ChanSendMsg <- sn
					C.ToChanSendMsgCount()
					G.ToGetTokenGoFishRe()
					R.ToGetTokenSuccess()
					a = 1
					break
				}
			}
		}
	}
	if a != 1 {
		G.ToGetTokenReDie()
	}

}

//下载分片失败后重试下载分片默认次数5
func RetryDownLoadSendMsg(n data.SendChan) {
	var c int
	for i := 0; i < R.Notice; i++ {
		var msg message.DownloadShardRequest
		ctxs, cancels := context.WithTimeout(context.Background(), time.Second*5)
		msg.VHF = n.Ns.ShardVHF
		msg.AllocId = n.Token
		buf, err := proto.Marshal(&msg)
		if err != nil {
			panic(err)
		}
		msgg, err := n.Clt.SendMsg(ctxs, message.MsgIDDownloadShardRequest.Value(), buf)
		cancels()
		if err != nil {
			log.Printf("sendmsg:err=%v node:id=%v shard:id=%v\n", err, n.Ns.NodeId, n.Ns.ShardVHF)
		} else if len(msgg) > 22 {
			R.ToSuccess()
			S.ToSendMsgGoFish()
			fmt.Printf("sendmsg:success=%v node:id=%v shard:id=%v\n", "success", n.Ns.NodeId, n.Ns.ShardVHF)
			c = 1
			break
		} else {
			log.Printf("sendmsg:kong=%v node:id=%v shard:id=%v\n", msgg, n.Ns.NodeId, n.Ns.ShardVHF)
		}
	}
	if c != 1 {
		R.ToSenderr()
		S.ToSendMsgGoConnDie()
	}
}

//统计延时等信息
func Performance(Stime int, Ssecond int) {
	fmt.Println("ing...")
	m := Stime * 60 / Ssecond
	logger := downlog.GetLog("performance")
	for i := 0; i < m; i++ {
		time.Sleep(time.Duration(Ssecond) * time.Second)
		if R.Count == int64(0) {
		} else {
			R.OldCount = R.Count
			R.Latency = R.Used / R.Count
			logger.Printf("Used %v Rate %v Latency %v\n",
				R.Used, float64(R.Count*16)/float64(Ssecond)/float64(1024), R.Latency)
			logger.Printf("TokenLive %v ByConn %v TokenConn %v TokenRe %v LiveMsg %v MakeMsg %v Shard %v\n",
				G.GetTokenGoLive, G.GetByGoConn, G.GetTokenGoConn, G.GetTokenRe, S.SendMsgGoLive, S.SendMsgGoConn, C.ChanSendMsgCount)
			logger.Println()
			R.Used = R.Latency
			R.Count = 1
		}
	}
	//R.Notice = 1
}

func DoDownLoadForData(n *mongo.Shard) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := client.Hst.ClientStore().GetByAddrString(ctx, client.N.NodeData[n.NodeId].NodeID, client.N.NodeData[n.NodeId].Addrs)
	if err != nil {
		//fmt.Println("DoDownLoadForData", err)
	} else {
		if int(I) < R.ShardCount {
			sn := data.ShardAndNodeId{ShardVHF: n.VHF, NodeId: n.NodeId}
			MakeGetTokenMsg <- sn
			atomic.AddInt64(&I, 1)
		}
	}
}
