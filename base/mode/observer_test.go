package mode

import (
	"context"
	"testing"
	"time"
)

// 同步模式
func Test_syncEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")

	sbus := NewSyncEventBus()
	topic := "order_finish"
	sbus.Subscribe(topic, observerA)
	sbus.Subscribe(topic, observerB)
	sbus.Subscribe(topic, observerC)
	sbus.Subscribe(topic, observerD)

	sbus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})
}

// 异步模式
func Test_asyncEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")

	abus := NewAsyncEventBus()
	defer abus.Stop()

	topic := "order_finish"
	abus.Subscribe(topic, observerA)
	abus.Subscribe(topic, observerB)
	abus.Subscribe(topic, observerC)
	abus.Subscribe(topic, observerD)

	abus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})
	<-time.After(time.Second)
}
