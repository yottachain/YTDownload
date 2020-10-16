package main

import (
	"DownloadNew/data"
	"DownloadNew/mongo"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"time"
)

var D data.DownLoadRate

func main() {
	D.ConSuccessRate = 0.9
	D.DownloadRate = 0.8
	D.ShardSuccessRate = 0.7
	D.TokenSuccessRate = 0.6
	c := mongo.Pool.Get()
	defer c.Close()
	data, err := json.Marshal(D)
	t := time.Now()
	filePtr, err := os.Create(t.Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("创建文件失败:", err)
	}
	writeString, err := filePtr.WriteString(string(data))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------", writeString)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, errs := c.Do("Set", "DownLoadRate", data)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	r, err := redis.Bytes(c.Do("Get", "DownLoadRate"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	err = json.Unmarshal(r, &D)
	fmt.Println("redis", D)
}
