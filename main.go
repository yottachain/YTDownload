package main

import (
	"DownloadNew/bedown"
	"DownloadNew/client"
	"DownloadNew/data"
	"DownloadNew/downlog"
	"DownloadNew/gorotine"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var skip int
	var limit int
	var num int
	var tnum int
	var retrynum int
	var Stime int
	var second int
	var mdata int
	var net int
	flag.IntVar(&skip, "s", 0, "起始位置，默认为0")
	flag.IntVar(&limit, "l", 500, "分片数量，默认为500")
	flag.IntVar(&num, "n", 500, "并发数量，默认为500")
	flag.IntVar(&tnum, "tn", 500, "token并发数量，默认为500")
	flag.IntVar(&retrynum, "r", 5, "重试次数，默认为5")
	flag.IntVar(&Stime, "t", 1, "时间范围，默认为1")
	flag.IntVar(&second, "sc", 10, "间隔秒数，默认为10")
	flag.IntVar(&mdata, "m", 0, "数据来源，默认为0")
	flag.IntVar(&net, "net", 0, "网络环境0研发网1主网，默认为0")
	flag.Parse() //解析命令行参数
	DownLoadO(skip, limit, num, retrynum, Stime, second, mdata, tnum, net)
}

//下载准备命令行模式
func DownLoadO(skip int, limit int, num int, retrynum int, Stime int, Ssecond int, mdata int, tnum int, net int) {
	go gorotine.Performance(Stime, Ssecond)
	gorotine.R.ShardCount = limit
	gorotine.R.Notice = retrynum
	logtotal := downlog.GetLog("total")
	bT := time.Now()
	if mdata == 1 {
		bedown.GetSharNodeDataK(skip, limit)
		fmt.Println("downloading...", len(gorotine.MakeGetTokenMsg), len(client.N.NodeData))
		gorotine.MakeGorotinesForShard(limit, num, tnum)
	} else {
		dataPrepare(skip, limit)
		fmt.Println("downloading...", len(gorotine.MakeGetTokenMsg))
		gorotine.MakeGorotinesForShard(limit, num, tnum)
	}
	eT := time.Since(bT)
	logtotal.Printf("msg=%v sharcount=%v success=%v sendfail=%v senderr=%v tcperr=%v GetTokenErr=%v GetTokenSuccess=%v concount===%v",
		"下载分片统计信息", gorotine.R.ShardCountTotal, gorotine.R.Success, gorotine.R.Sendfail, gorotine.R.Senderr, gorotine.R.Adderr, gorotine.R.GetTokenFail, gorotine.R.GetTokenSuccess, gorotine.R.ConCount)
	logtotal.Printf("msg=%v 下载总耗时=%v 下载成功率%v 下载速率%v 下载成功速率%v",
		"下载速率统计", eT, float64(gorotine.R.Success)/float64(gorotine.R.ShardCountTotal), float64(gorotine.R.ShardCountTotal*16)/gorotine.R.UsedTotal/1024, float64(gorotine.R.Success*16)/gorotine.R.UsedTotal/1024)
	fmt.Printf("%T %v\n", gorotine.R.UsedTotal, gorotine.R.UsedTotal)
	fmt.Printf("%T %v\n", gorotine.R.GetTokenErr, gorotine.R.GetTokenErr)
	gorotine.T.PrintNode()
	writeRedisToMtrics()
}

func dataPrepare(skip int, limit int) {
	fmt.Println("data prepare...about 25s")
	bedown.GetSharNodeData(skip, limit*5, 200000)
}

var D data.DownLoadRate

func writeRedisToMtrics() {
	D.ConSuccessRate = float64(gorotine.R.ConCount) / float64(gorotine.R.ShardCount)
	D.DownloadRate = float64(gorotine.R.Success*16) / gorotine.R.UsedTotal / 1024
	D.ShardSuccessRate = float64(gorotine.R.Success) / float64(gorotine.R.ShardCountTotal)
	D.TokenSuccessRate = float64(gorotine.R.GetTokenSuccess) / float64(gorotine.R.ConCount)
	D.ShardCountTotal = gorotine.R.ShardCountTotal
	D.ConCount = int64(gorotine.R.ShardCount)
	D.TokenCount = gorotine.R.ConCount
	D.GetTokenSuccess = gorotine.R.GetTokenSuccess
	D.ShardDownLoadFail = int64(gorotine.R.Senderr)
	D.SuccessShardCount = int64(gorotine.R.Success)
	D.TimeTotal = gorotine.R.UsedTotal
	Post("http://192.168.6.137:8081/downloadRate", D, "application/json")
	//c := mongo.Pool.Get()
	//defer c.Close()
	//
	//data, err := json.Marshal(D)
	//t := time.Now()
	//filePtr, err := os.Create("/var/tmp/" + t.Format("2006-01-02 15:04:05"))
	//if err != nil {
	//	fmt.Println("创建文件失败:", err)
	//}
	//writeString, err := filePtr.WriteString(string(data))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("------", writeString)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//_, errs := c.Do("Set", "DownLoadRate", data)
	//if errs != nil {
	//	fmt.Println(errs)
	//	return
	//}
	//r, err := redis.Bytes(c.Do("Get", "DownLoadRate"))
	//if err != nil {
	//	fmt.Println("get abc failed,", err)
	//	return
	//}
	//err = json.Unmarshal(r, &D)
	//fmt.Println("redis", D)
}

func Post(url string, data interface{}, contentType string) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	_, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
}
