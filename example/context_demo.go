package example

import (
	"context"
	"fmt"
	"time"
)

func Context1() {
	deadlineCtx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Second*3))
	defer cancel()
	select {
	case <-deadlineCtx.Done():
		fmt.Println("abc------------", deadlineCtx.Err())
		return
	}
	//time.Sleep(time.Second * 5)
	fmt.Println("bbbbbb--------")
}

func context2() {
	deadlineCtx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	select {
	case <-deadlineCtx.Done():
		fmt.Println("abc------------", deadlineCtx.Err())
		return
	}
	time.Sleep(time.Second * 5)
	fmt.Println("bbbbbb--------")
}
