package mongo

import (
	"gopkg.in/mgo.v2"
)

//const URL = "47.95.231.56:27020" //mogon分片集
const URL = "39.97.229.11:27017" //公网
//const URL = "122.152.203.189:27017"
//const URL = "139.155.46.242:27017" //研发网

var (
	mgoSession *mgo.Session
	//dataBase   = "yotta"
)

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

//公共方法，获取collection对象
func WitchCollection(dataBase string, collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}
