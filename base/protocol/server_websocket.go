package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 参考： https://zhuanlan.zhihu.com/p/669547021
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", wsUpGrader)
	err := http.ListenAndServe("localhost:8088", nil)
	if err != nil {
		log.Println("server start err", err)
	}
}

func wsUpGrader(w http.ResponseWriter, r *http.Request) {
	//转换为升级为websocket
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//释放连接
	defer conn.Close()

	for {
		//接收消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("server receive messageType", messageType, "message", string(message))
		//发送消息
		err = conn.WriteMessage(messageType, []byte("pong"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
