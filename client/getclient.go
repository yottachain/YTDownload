package client

import (
	"DownloadNew/mongo"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"

	//"github.com/yottachain/YTCoreService/api"
	host "github.com/yottachain/YTHost"
	inter "github.com/yottachain/YTHost/hostInterface"
)

//var C *api.Client
var Hst inter.Host
var N mongo.NodeData

func init() {
	//api.StartApi()
	////c, err := api.NewClient("ianmooneyy11", "5JXtdc52ARN3922wFJDobx6ozwEvcDKg8aQsxJ9Dknd6U4NfALY")
	//c, err := api.NewClient("pollytestdev", "5JykGxNccDmFX2wgaLkX6GkqTpntn4VEQxDR99jgP9KRK8f4zya")
	////c, err := api.NewClient("pollytestde1", "5JsohFvnt2qhkKxzConrJSxU2ti4qGifjJ9dGCxhpup4EYw1es8")
	//if err != nil {
	//	fmt.Println("NewClient-err:", err)
	//} else {
	//	C = c
	//}
	//query := bson.M{"valid": 1, "status": 1, "assignedSpace": bson.M{"$gt": 0}, "quota": bson.M{"$gt": 0}, "weight": bson.M{"$gt": 0}, "version": bson.M{"$gte": 48}}
	now := time.Now()
	t := now.Add(time.Minute * -3)
	query := bson.M{"timestamp": bson.M{"$gt": int(t.Unix())}}
	node, _ := mongo.SearchNode("yotta", "Node", query, "", nil, 0, 0)
	M := make(map[int32]mongo.Node)
	for _, v := range node {
		M[v.ID] = v
	}
	N.NodeData = M
	hste, err := host.NewHost()
	if err != nil {
		fmt.Println("newHost", err)
	}
	Hst = hste
}
