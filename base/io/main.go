package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"unsafe"
)

var strSlash = []byte("/")

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
func countSections(path string) uint16 {
	s := StringToBytes(path)
	return uint16(bytes.Count(s, strSlash))
}
func main() {
	fmt.Println(countSections("/api/"))
}

// 启动一个 tcp 服务端代码示例
func main11() {
	// 创建一个 tcp 端口监听器
	l, _ := net.Listen("tcp", ":8080")
	// 主动轮询模型
	for {
		// 等待 tcp 连接到达
		conn, _ := l.Accept()
		// 开启一个 goroutine 负责一笔客户端请求的处理
		go serve(conn)
	}
}

// 处理一笔 tcp 连接
func serve(conn net.Conn) {
	defer conn.Close()
	var buf []byte
	// 读取连接中的数据
	_, _ = conn.Read(buf)
	// ...
}

// http标准库实现
func main2() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.HandleFunc("/ping22", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong22"))
	})
	http.ListenAndServe(":8091", nil)

	reqBody, _ := json.Marshal(map[string]string{"key1": "val1", "key2": "val2"})

	resp, _ := http.Post(":8091", "application/json", bytes.NewReader(reqBody))
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("resp: %s", respBody)
}

func main3() {
	// 创建一个 gin Engine，本质上是一个 http Handler
	var mux = gin.Default() // 注册中间件
	mux.Use(myMiddleWare)
	// 注册一个 path 为 /ping 的处理函数
	mux.POST("/ping", myHandleFunc)
	// 运行 http 服务
	if err := mux.Run(":8080"); err != nil {
		panic(err)
	}
}
func myMiddleWare(c *gin.Context) {
}
func myHandleFunc(c *gin.Context) {
	// 前处理
	preHandle()
	c.JSON(http.StatusOK, "pone")
	c.Next()
	// 后处理
	postHandle()
}

func preHandle() {
	fmt.Println("preHandle------------")
}

func postHandle() {
	fmt.Println("postHandle------------")
}
