package bedown

import (
	"DownloadNew/data"
	"DownloadNew/gorotine"
	"DownloadNew/mongo"
	"encoding/binary"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

var datas []mongo.Shard

//获取数据前准备
func SharNodeData(limit int, num int, skip int, retrynum int) {
	fl, _ := os.OpenFile("shard.json", os.O_RDONLY, 0644)
	dc := json.NewDecoder(fl)
	dc.Decode(&datas)
	//q := gorotine.NewLoopQueue()
	//for i := 0; i < limit; i++ {
	//	q.Enqueue(data[i])
	//}
	//gorotine.R.Notice = retrynum
	//gorotine.MakeGorotinesForQueue(q, num)

	var shardData []mongo.Shard
	if limit < 140000 {
		for i := 0; i < limit; i++ {
			shardData = append(shardData, datas[i])
		}
	} else {
		n := limit / 100000
		for i := 0; i < n; i++ {
			for i := 0; i < len(datas); i++ {
				shardData = append(shardData, datas[i])
			}
		}
	}
	//gorotine.R.ShardCount = limit
	gorotine.R.Notice = retrynum
	//gorotine.MakeGorotinesForData(&shardData, num)
}

//从mongo里获取数据未过滤
func GetSharNodeDataK(skip int, limit int) bool {
	query := bson.M{"_id": bson.M{"$gt": getMongoId(skip)}}
	shard, err := mongo.SearchShard("metabase", "shards", query, "", nil, skip, limit)
	if err != nil {
		panic(err)
	}
	for _, v := range shard {
		sn := data.ShardAndNodeId{ShardVHF: v.VHF, NodeId: v.NodeId}
		gorotine.MakeGetTokenMsg <- sn
	}
	close(gorotine.MakeGetTokenMsg)
	return true
}

//从mongo里获取数据
func GetSharNodeData(skip int, limit int, num int) bool {
	query := bson.M{"_id": bson.M{"$gt": getMongoId(skip)}}
	shard, err := mongo.SearchShard("metabase", "shards", query, "", nil, skip, limit)
	if err != nil {
		panic(err)
	}
	gorotine.MakeGorotinesForData(&shard, num)
	return true
}

func getMongoId(skip int) int64 {
	timeLayout := "2006-01-02 15:04:05"
	oldTime := time.Now().AddDate(0, 0, -skip)
	timesTwo, _ := time.Parse(timeLayout, oldTime.Format("2006-01-02 15:04:05"))
	b := timeForInt64(timesTwo.Unix())
	return b
}

func timeForInt64(timeUnix int64) int64 {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, uint32(int32(timeUnix)))
	mongoId := int64(binary.BigEndian.Uint64(buf))
	return mongoId
}
