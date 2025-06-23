package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	//服务器地址 websocket 统一使用 ws://
	url := "ws://localhost:8088/ws"
	//使用默认拨号器，向服务器发送连接请求
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//关闭连接
	defer ws.Close()
	//发送消息
	go func() {
		for {
			err := ws.WriteMessage(websocket.BinaryMessage, []byte("ping"))
			if err != nil {
				log.Fatal(err)
			}
			//休眠两秒
			time.Sleep(time.Second * 2)
		}
	}()

	//接收消息
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("client receive message: ", string(data))
	}
}
