package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//定义student结构,变量大写
type Shard struct {
	Id     int64  `bson:"_id" form:"_id" json:"_id"`
	NodeId int32  `bson:"nodeId" form:"nodeId" json:"nodeId"`
	VHF    []byte `bson:"VHF" form:"VHF" json:"VHF"`
}

var shards []Shard

func SearchShard(dataBase string, collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (shards []Shard, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Skip(skip).Limit(limit).All(&shards)
	}
	err = WitchCollection(dataBase, collectionName, exop)
	return
}

func SearchShardOne(dataBase string, collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (shards []Shard, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Limit(limit).All(&shards)
	}
	err = WitchCollection(dataBase, collectionName, exop)
	return
}

func SearchShardTime(dataBase string, collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (shards []Shard, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Limit(limit).All(&shards)
	}
	err = WitchCollection(dataBase, collectionName, exop)
	return
}

//查询矿机表
type Node struct {
	ID     int32  `bson:"_id" form:"id" json:"_id"`
	NodeID string `bson:"nodeid" form:"nodeId" json:"nodeid"`
	PubKey string `bson:"pubkey" form:"pubKey" json:"pubKey"`
	//Owner           string      `bson:"owner" form:"owner" json:"owner"`
	//ProfitAcc       string      `bson:"profitAcc" form:"profitAcc" json:"profitAcc"`
	//PoolID          string      `bson:"poolID" form:"poolID" json:"poolID"`
	//PoolOwner       string      `bson:"poolOwner" form:"poolOwner" json:"poolOwner"`
	//Quota           int64       `bson:"quota" form:"quota" json:"quota"`
	Addrs []string `bson:"addrs" form:"addrs" json:"addrs"`
	//CPU             int32       `bson:"cpu" form:"cpu" json:"cpu"`
	//Memory          int32       `bson:"memory" form:"memory" json:"memory"`
	//Bandwidth       int32       `bson:"bandwidth" form:"bandwidth" json:"bandwidth"`
	//MaxDataSpace    int64       `bson:"maxDataSpace" form:"maxDataSpace" json:"maxDataSpace"`
	//AssignedSpace   int64       `bson:"assignedSpace" form:"assignedSpace" json:"assignedSpace"`
	//ProductiveSpace int64       `bson:"productiveSpace" form:"productiveSpace" json:"productiveSpace"`
	//UsedSpace       int64       `bson:"usedSpace" form:"usedSpace" json:"usedSpace"`
	//Uspaces interface{} `bson:"uspaces" form:"uspaces" json:"uspaces"`
	//Weight          float64     `bson:"weight" form:"weight" json:"weight"`
	//Valid           int32       `bson:"valid" form:"valid" json:"valid"`
	//Relay           int32       `bson:"relay" form:"relay" json:"relay"`
	//Status          int32       `bson:"status" form:"status" json:"status"`
	Timestamp int64 `bson:"timestamp" form:"timestamp" json:"timestamp"`
	//Version         int32       `bson:"version" form:"version" json:"version"`
	//Rebuilding      int8        `bson:"rebuilding" form:"rebuilding" json:"rebuilding"`
	//RebuildingTask  uint32      `bson:"rebuildingtask" form:"rebuildingTask" json:"rebuildingTask"`
	//RealSpace       int64       `bson:"realSpace" form:"realSpace" json:"realSpace"`
	//Tx              int64       `bson:"tx" form:"tx" json:"tx"`
	//Rx              int64       `bson:"rx" form:"rx" json:"rx"`
	//Other           [][]struct {
	//	Key   string      `bson:"Key" form:"Key" json:"Key"`
	//	Value interface{} `bson:"Value" form:"Value" json:"Value"`
	//} `bson:"Other" form:"Other" json:"Other"`
	//ManualWeight int32 `bson:"ManualWeight" form:"ManualWeight" json:"ManualWeight"`
}

// Nodes .
type Nodes []Node

func SearchNode(dataBase string, collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (Nodes []Node, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).All(&Nodes)
	}
	err = WitchCollection(dataBase, collectionName, exop)
	return
}

func SearchNodeOne(dataBase string, collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (Nodes []Node, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Skip(skip).Limit(limit).All(&Nodes)
	}
	err = WitchCollection(dataBase, collectionName, exop)
	return
}

type ShardNode struct {
	Shard Shard
	Node  Node
}
