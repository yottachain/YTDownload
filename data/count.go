package data

import "DownloadNew/mongo"

type CountForAll struct {
	BucketCount      int
	FileCount        int
	BlockCount       int
	ShardCount       int
	SDKShardCount    int
	ChanShardCount   int
	ChanSendMsgCount int
}

func (c *CountForAll) ToBucketCount() {
	lock.Lock()
	defer lock.Unlock()
	c.BucketCount++
}

func (c *CountForAll) ToFileCount() {
	lock.Lock()
	defer lock.Unlock()
	c.FileCount++
}

func (c *CountForAll) ToBlockCount() {
	lock.Lock()
	defer lock.Unlock()
	c.BlockCount++
}

func (c *CountForAll) ToShardCount() {
	lock.Lock()
	defer lock.Unlock()
	c.ShardCount++
}

func (c *CountForAll) ToSDKShardCount() {
	lock.Lock()
	defer lock.Unlock()
	c.SDKShardCount++
}

func (c *CountForAll) ToChanShardCount() {
	lock.Lock()
	defer lock.Unlock()
	c.ChanShardCount++
}

func (c *CountForAll) ToChanShardCountJ() {
	lock.Lock()
	defer lock.Unlock()
	c.ChanShardCount--
}

func (c *CountForAll) ToChanSendMsgCount() {
	lock.Lock()
	defer lock.Unlock()
	c.ChanSendMsgCount++
}

func (c *CountForAll) ToChanSendMsgCountJ() {
	lock.Lock()
	defer lock.Unlock()
	c.ChanSendMsgCount--
}

type SendMsgGo struct {
	SendMsgGoLive int32
	SendMsgGoConn int32
	SendMsgGoFish int32
}

func (s *SendMsgGo) ToSendMsgGoLive() {
	lock.Lock()
	defer lock.Unlock()
	s.SendMsgGoLive++
}

func (s *SendMsgGo) ToSendMsgGoConn() {
	lock.Lock()
	defer lock.Unlock()
	s.SendMsgGoConn++
}

func (s *SendMsgGo) ToSendMsgGoFish() {
	lock.Lock()
	defer lock.Unlock()
	s.SendMsgGoLive--
	s.SendMsgGoConn--
	s.SendMsgGoFish++
}

func (s *SendMsgGo) ToSendMsgGoConnDie() {
	lock.Lock()
	defer lock.Unlock()
	s.SendMsgGoLive--
	s.SendMsgGoConn--
	s.SendMsgGoFish++
}

type GetTokenGo struct {
	GetTokenGoLive int32
	GetByGoConn    int32
	GetTokenGoConn int32
	GetTokenGoFish int32
	GetTokenRe     int32
}

func (g *GetTokenGo) ToGetTokenGoLive() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoLive++
	g.GetTokenGoFish--
}

func (g *GetTokenGo) ToGetByGoConn() {
	lock.Lock()
	defer lock.Unlock()
	g.GetByGoConn++
}

func (g *GetTokenGo) ToGetByGoConnDie() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoLive--
	g.GetByGoConn--
}

func (g *GetTokenGo) ToGetTokenGoConn() {
	lock.Lock()
	defer lock.Unlock()
	g.GetByGoConn--
	g.GetTokenGoConn++
}

func (g *GetTokenGo) ToGetTokenGoFish() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoLive--
	g.GetTokenGoConn--
	g.GetTokenGoFish++
}

func (g *GetTokenGo) ToGetTokenGoConnDie() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoLive--
	g.GetTokenGoConn--
}

func (g *GetTokenGo) ToGetTokenGoFishJ() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoFish--
}

func (g *GetTokenGo) ToGetTokenGoFishRe() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoLive--
	g.GetTokenGoFish++
	g.GetTokenRe--
}

func (g *GetTokenGo) ToGetTokenRe() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoConn--
	g.GetTokenRe++
}

func (g *GetTokenGo) ToGetTokenReDie() {
	lock.Lock()
	defer lock.Unlock()
	g.GetTokenGoLive--
	g.GetTokenRe--
}

type GetCount struct {
	Adderr               int
	Senderr              int
	Success              int
	Sendfail             int
	GetTokenErr          int64
	Used                 int64
	UsedTotal            float64
	Count                int64
	Latency              int64
	OldCount             int64
	Retry                int64
	ShardCount           int
	Notice               int
	File                 []mongo.Shard
	FileCount            int
	Live                 int64
	LiveMsg              int64
	DieConn              int64
	DieMsg               int64
	MakeConn             int64
	MakeMsg              int64
	Finish               int64
	Receive              int64
	GetTokenFail         int64
	GetTokenSuccess      int64
	GetTokenSuccessCount int64
	SendChanNs           []SendChan
	ChannelCount         int64
	ConCount             int64
	ShardCountTotal      int64
	Net                  int
}

func (gc *GetCount) ToShardCountTotal() {
	lock.Lock()
	defer lock.Unlock()
	gc.ShardCountTotal++
}

func (gc *GetCount) ToConCount() {
	lock.Lock()
	defer lock.Unlock()
	gc.ConCount++
}

func (gc *GetCount) ToSuccess() {
	lock.Lock()
	defer lock.Unlock()
	gc.Success++
}

func (gc *GetCount) ToSenderr() {
	lock.Lock()
	defer lock.Unlock()
	gc.Senderr++
}

func (gc *GetCount) ToSendfail() {
	lock.Lock()
	defer lock.Unlock()
	gc.Sendfail++
}

func (gc *GetCount) ToAdderr() {
	lock.Lock()
	defer lock.Unlock()
	gc.Adderr++
}

func (gc *GetCount) ToGetTokenSuccess() {
	lock.Lock()
	defer lock.Unlock()
	gc.GetTokenSuccess++
}

func (gc *GetCount) ToGetTokenFail() {
	lock.Lock()
	defer lock.Unlock()
	gc.GetTokenFail++
}

func (gc *GetCount) ToCount() {
	lock.Lock()
	defer lock.Unlock()
	gc.Count++
}

func (gc *GetCount) ToUsed(bt int64) {
	lock.Lock()
	defer lock.Unlock()
	gc.Used += bt
}

func (gc *GetCount) ToUsedTotal(bt float64) {
	lock.Lock()
	defer lock.Unlock()
	gc.UsedTotal += bt
}

type DownLoadRate struct {
	TokenSuccessRate  float64
	ShardSuccessRate  float64
	ConSuccessRate    float64
	DownloadRate      float64
	ConCount          int64 //连接总数
	TokenCount        int64 //token总数 and 连接成功数
	GetTokenSuccess   int64 //成功的token数量
	ShardCountTotal   int64 //真实的分片
	SuccessShardCount int64 //成功的分片数
	ShardDownLoadFail int64 //下载失败的分片数
}
