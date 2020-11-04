package main

import (
	"DownloadNew/data"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var D data.DownLoadRate

func main() {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	fmt.Println("**************************")
	goPath := os.Getenv("GOPATH")
	fmt.Printf("GOPATH is %s\n", goPath)
	//t := time.Now()
	//time.Sleep(time.Second *10)
	//b := time.Since(t)
	//fmt.Printf("%T %v",b.Seconds(),b.Seconds())
}

//func main() {
//	D.ConSuccessRate = 0.9
//	D.DownloadRate = 0.8
//	D.ShardSuccessRate = 0.7
//	D.TokenSuccessRate = 0.6
//	Post("http://192.168.6.137:8081/downloadRate",D,"application/json")
//	//c := mongo.Pool.Get()
//	//defer c.Close()
//	//data, err := json.Marshal(D)
//	//t := time.Now()
//	//filePtr, err := os.Create(t.Format("2006-01-02 15:04:05"))
//	//if err != nil {
//	//	fmt.Println("创建文件失败:", err)
//	//}
//	//writeString, err := filePtr.WriteString(string(data))
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//fmt.Println("------", writeString)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//_, errs := c.Do("Set", "DownLoadRate", data)
//	//if errs != nil {
//	//	fmt.Println(errs)
//	//	return
//	//}
//	//r, err := redis.Bytes(c.Do("Get", "DownLoadRate"))
//	//if err != nil {
//	//	fmt.Println("get abc failed,", err)
//	//	return
//	//}
//	//err = json.Unmarshal(r, &D)
//	//fmt.Println("redis", D)
//}

func Post(url string, data interface{}, contentType string) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	fmt.Println("lalalala", resp.Status)
}
