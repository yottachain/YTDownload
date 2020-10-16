package data

import (
	"github.com/yottachain/YTHost/client"
)

//type YTShard struct {
//	Length     int64
//	UClient    *api.Client
//	REFS       map[int32]*api.BlockInfo
//	ADDR       map[int32][]string
//	ShardCount int32
//}

//type YTShard struct {
//	NodeId int32
//	NodeAddr []string
//	ShardVHF []byte
//	ShardCount int
//}

type ShardAndNodeId struct {
	ShardVHF []byte
	NodeId   int32
}

type SendChan struct {
	Ns    ShardAndNodeId
	Clt   *client.YTHostClient
	Token string
}
