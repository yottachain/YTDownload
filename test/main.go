package main

import (
	"DownloadNew/data"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var D data.DownLoadRate

func main() {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	skip := os.Getenv("SKIP")
	SKIP, err := strconv.Atoi(skip)
	if err != nil {
		log.Println(err)
		SKIP = 0
	}
	limit := os.Getenv("LIMIT")
	LIMIT, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
		LIMIT = 500
	}
	num := os.Getenv("NUM")
	NUM, err := strconv.Atoi(num)
	if err != nil {
		log.Println(err)
		NUM = 500
	}
	retrynum := os.Getenv("RETRY_NUM")
	RETRY_NUM, err := strconv.Atoi(retrynum)
	if err != nil {
		log.Println(err)
		RETRY_NUM = 5
	}
	Stime := os.Getenv("S_TIME")
	S_TIME, err := strconv.Atoi(Stime)
	if err != nil {
		log.Println(err)
		S_TIME = 3
	}
	second := os.Getenv("SECOND")
	SECOND, err := strconv.Atoi(second)
	if err != nil {
		log.Println(err)
		SECOND = 10
	}
	mdata := os.Getenv("M_DATA")
	M_DATA, err := strconv.Atoi(mdata)
	if err != nil {
		log.Println(err)
		M_DATA = 0
	}
	tnum := os.Getenv("T_NUM")
	T_NUM, err := strconv.Atoi(tnum)
	if err != nil {
		log.Println(err)
		T_NUM = 500
	}
	net := os.Getenv("NET")
	NET, err := strconv.Atoi(net)
	if err != nil {
		log.Println(err)
		NET = 1
	}
	fmt.Println(SKIP, LIMIT, NUM, RETRY_NUM, S_TIME, SECOND, M_DATA, T_NUM, NET)
}

//func main() {
//	environ := os.Environ()
//	for i := range environ {
//		fmt.Println(environ[i])
//	}
//	fmt.Println("**************************")
//	goPath := os.Getenv("GOPATH")
//	fmt.Printf("GOPATH is %s\n", goPath)
//	//t := time.Now()
//	//time.Sleep(time.Second *10)
//	//b := time.Since(t)
//	//fmt.Printf("%T %v",b.Seconds(),b.Seconds())
//}

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
