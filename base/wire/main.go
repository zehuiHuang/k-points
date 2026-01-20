package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

/**
wire中的provider 和 injector的概念
provider是创建依赖的函数,比如本案例中的NewMessage、NewGreeter和NewEvent的函数
injector一般是在wire.go中,需要增加如下代码:
//go:build wireinject
// +build wireinject

具体查看wire.go代码

通过命令 wire gen .  在当前目录下自动生成wire_gen.go

在wire_gen.go可以查看生成的代码里,自动生成了依赖

*/

// 依赖关系:Event->Greeter->Message

type Message string

// NewMessage wire中的Provider,带参数的函数
func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Message Message
}

// NewGreeter wire中的Provider
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

// NewEvent wire中的Provider,可以添加错误返回值
func NewEvent(g Greeter) (Event, error) {
	if time.Now().Unix()%2 == 0 {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

//func main() {
//	message := NewMessage()
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

//func main() {
//	event := InitializeEvent()
//	event.Start()
//}

//func main() {
//	e, err := InitializeEvent()
//	if err != nil {
//		fmt.Printf("failed to create event: %s\n", err)
//		os.Exit(2)
//	}
//	e.Start()
//}

func main() {
	event, err := InitializeEvent("Hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	event.Start()
}
