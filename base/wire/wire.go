//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

//	func InitializeEvent() Event {
//		wire.Build(NewEvent, NewGreeter, NewMessage)
//		return Event{}
//	}

// 有错误返回值的

var wireSet = wire.NewSet(NewGreeter, NewMessage)

func InitializeEvent(phrase string) (Event, error) {
	panic(wire.Build(wireSet, NewEvent))
}
