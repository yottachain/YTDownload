package data

import (
	"sync"
)

type YTFile struct {
	FileName      string
	FileNameCount int
	BucketName    string
	//Version       primitive.ObjectID
}

type YTFileMap struct {
	sync.Map
}

var lock sync.Mutex

func (f *YTFile) ToFileName(n string) {
	lock.Lock()
	defer lock.Unlock()
	f.FileName = n
}

func (f *YTFile) ToFileNameCount(n int) {
	lock.Lock()
	defer lock.Unlock()
	f.FileNameCount += n
}

func (ns *YTFileMap) ToFileMap(n string, s string) {
	ns.Map.LoadOrStore(n, s)
}

func (ns *YTFileMap) GetFileMap(n string) interface{} {
	if v, ok := ns.Map.Load(n); ok {
		return v
	}
	return ""
}

func (ns *YTFileMap) DeFileMap(n string) {
	ns.Map.Delete(n)
}
