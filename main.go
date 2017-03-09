package main

import (
	"runtime"
	"net/http"
	"fmt"
	"strings"
	"log"
)

const (
	//下载进程数
	numDowner = 10
	//通道缓存
	chanBufferSize = 1000
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func sayHello(response http.ResponseWriter, request *http.Request) {
	request.ParseForm() //解析参数，默认不解析
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url.long"])
	for k, v := range request.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(response, "Hello Roy")
}

func main() {
	// 设置访问路由
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
