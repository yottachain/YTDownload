package data

import (
	"DownloadNew/downlog"
	"sync"
)

type NodeData struct {
	ConnSuccess        int
	ConnErrs           int
	GtSuccess          int
	GtErrs             int
	GtAvgDelay         int
	GtMaxDelay         int
	SendSuccess        int
	SendErrs           int
	SendAvgDelay       int
	SendMaxDelay       int
	SuccessNode        int64
	SenderrNode        int64
	SendfailNode       int64 //发送
	AdderrNode         int64 //矿机地址错误
	RetryNode          int64 //重试矿机次数
	GetTokenErrNode    int64 //获取token错误
	GetTokenTcpErrNode int64 //获取token通信错误
	RetryToken         int64 //重试token次数
}

func (n *NodeData) ToRetryToken() {
	lock.Lock()
	defer lock.Unlock()
	n.RetryToken++
}

func (n *NodeData) ToGetTokenErrNode() {
	lock.Lock()
	defer lock.Unlock()
	n.GetTokenErrNode++
}

func (n *NodeData) ToGetTokenTcpErrNode() {
	lock.Lock()
	defer lock.Unlock()
	n.GetTokenTcpErrNode++
}

func (n *NodeData) ToRetryNode() {
	lock.Lock()
	defer lock.Unlock()
	n.RetryNode++
}

func (n *NodeData) ToAdderrNode() {
	lock.Lock()
	defer lock.Unlock()
	n.AdderrNode++
}

func (n *NodeData) ToSendfailNode() {
	lock.Lock()
	defer lock.Unlock()
	n.SendfailNode++
}

func (n *NodeData) ToSenderrNode() {
	lock.Lock()
	defer lock.Unlock()
	n.SenderrNode++
}

func (n *NodeData) ToSuccessNode() {
	lock.Lock()
	defer lock.Unlock()
	n.SuccessNode++
}

type NodeOneData struct {
	sync.Map
}

func (ns *NodeOneData) NoAdderr(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToAdderrNode()
}

func (ns *NodeOneData) NoSenderr(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToSenderrNode()
}

func (ns *NodeOneData) NoSuccess(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToSuccessNode()
}

func (ns *NodeOneData) NoSendfail(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToSendfailNode()
}

func (ns *NodeOneData) NoRetry(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToRetryNode()
}

func (ns *NodeOneData) NoGetTokenErrNode(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToGetTokenErrNode()
}

func (ns *NodeOneData) NoGetTokenTcpErrNode(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToGetTokenTcpErrNode()
}

func (ns *NodeOneData) NoRetryToken(nodeid int32) {
	st, ok := ns.Map.Load(nodeid)
	if !ok {
		st = &NodeData{RetryNode: 0, AdderrNode: 0, SendfailNode: 0, SenderrNode: 0, SuccessNode: 0, GetTokenErrNode: 0, GetTokenTcpErrNode: 0, RetryToken: 0}
		ns.Map.LoadOrStore(nodeid, st)
	}
	st.(*NodeData).ToRetryToken()
}

func (ns *NodeOneData) PrintNode() {
	getLog := downlog.GetLog("oneNode")
	f := func(k, v interface{}) bool {
		getLog.Printf("nodeid=%v Adderr=%v Senderr=%v Success=%v Sendfail=%v Retry=%v GetTokenErrNode=%v GetTokenTcpErrNode=%v RetryToken=%v\n",
			k, v.(*NodeData).AdderrNode, v.(*NodeData).SenderrNode, v.(*NodeData).SuccessNode, v.(*NodeData).SendfailNode, v.(*NodeData).RetryNode, v.(*NodeData).GetTokenErrNode, v.(*NodeData).GetTokenTcpErrNode, v.(*NodeData).RetryToken)
		return true
	}
	ns.Map.Range(f)
	return
}
