package gorotine

import (
	"DownloadNew/data"
	"DownloadNew/mongo"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

//并发下载分片数据过滤
func doWorkerOne(i interface{}) {
	n := i.(data.ShardAndNodeId)
	DownLoadByAdress(n)
}

func MakeGorotinesOne(num int) {
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(num, func(i interface{}) {
		doWorkerOne(i)
		wg.Done()
	})
	defer p.Release()
	for true {
		s, ok := <-MakeGetTokenMsg
		if ok {
			wg.Add(1)
			_ = p.Invoke(s)
		} else {
			break
		}
	}
	//for i := 0; i < limit; i++ {
	//	wg.Add(1)
	//	_ = p.Invoke(<-MakeGetTokenMsg)
	//}
	wg.Wait()
	close(ChanSendMsg)
	//fmt.Printf("ChanSendMsg running goroutines: %d\n", p.Running())
}

func doWorkerTwo(i interface{}) {
	n := i.(data.SendChan)
	DownLoadSendMsg(n)
}

func MakeGorotinesForShard(limit int, num int, tnum int) {
	go MakeGorotinesOne(tnum)
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(num, func(i interface{}) {
		doWorkerTwo(i)
		wg.Done()
	})
	defer p.Release()
	sendStartTime := time.Now()
	for true {
		s, ok := <-ChanSendMsg
		if ok {
			wg.Add(1)
			_ = p.Invoke(s)
			C.ToChanSendMsgCountJ()
		} else {
			break
		}
	}
	wg.Wait()
	sendEndTime := time.Now()
	timeSend := sendEndTime.Sub(sendStartTime)
	R.ToUsedTotal(timeSend.Seconds())
}

//---------------
func doWorkerForData(i interface{}) {
	n := i.(mongo.Shard)
	DoDownLoadForData(&n)
}

func MakeGorotinesForData(n *[]mongo.Shard, num int) {
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(num, func(i interface{}) {
		doWorkerForData(i)
		wg.Done()
	})
	defer p.Release()
	for _, s := range *n {
		wg.Add(1)
		_ = p.Invoke(s)
	}
	wg.Wait()
	close(MakeGetTokenMsg)
}

//----------------
