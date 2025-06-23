package mode

import (
	"context"
	"fmt"
	"sync"
)

// 三个角色：观察者Observer、事件Event(事物的变更事件)、事件总线EventBus
// 首先，定义事件，主要包括事件类型和事件信息；第二，定义观察者接口，主要是接收事件变更的信息；第三，定义事件总线，主要包括订阅，取消订阅、事件发布等动作

type Event struct {
	Topic string
	Val   interface{}
}

type Observer interface {
	OnChange(ctx context.Context, e *Event) error
}

type EventBus interface {
	Subscribe(topic string, o Observer)
	Unsubscribe(topic string, o Observer)
	Publish(ctx context.Context, e *Event)
}

//--------观察者实现

type BaseObserver struct {
	name string
}

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{
		name: name,
	}
}
func (b *BaseObserver) OnChange(ctx context.Context, e *Event) error {
	fmt.Printf("observer: %s, event key: %s, event val: %v\n", b.name, e.Topic, e.Val)
	// ...
	return nil
}

//------------------事件总线基本实现

type BaseEventBus struct {
	mux       sync.RWMutex
	observers map[string]map[Observer]struct{}
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]struct{}),
	}
}
func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	_, ok := b.observers[topic]
	if !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}
	b.observers[topic][o] = struct{}{}
}
func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	delete(b.observers[topic], o)
}

// -----------------------------事件总线具体实现

type SyncEventBus struct {
	BaseEventBus
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
}

//类比mq的消息发送

func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	subscribers := s.observers[e.Topic]

	errs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.OnChange(ctx, e); err != nil {
			errs[subscriber] = err
		}
	}

	s.handleErr(ctx, errs)
}

func (s *SyncEventBus) handleErr(ctx context.Context, errs map[Observer]error) {
	for o, err := range errs {
		// 处理 publish 失败的 observer
		fmt.Printf("observer: %v, err: %v", o, err)
	}
}

//--------------------------------------------异步模式

type observerWithErr struct {
	o   Observer
	err error
}

type AsyncEventBus struct {
	BaseEventBus
	errC chan *observerWithErr
	ctx  context.Context
	stop context.CancelFunc
}

func NewAsyncEventBus() *AsyncEventBus {
	aBus := AsyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
	aBus.ctx, aBus.stop = context.WithCancel(context.Background())
	// 处理处理错误的异步守护协程
	go aBus.handleErr()
	return &aBus
}
func (a *AsyncEventBus) Stop() {
	a.stop()
}

func (a *AsyncEventBus) Publish(ctx context.Context, e *Event) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	subscribers := a.observers[e.Topic]
	for subscriber := range subscribers {
		// shadow
		subscriber := subscriber
		go func() {
			if err := subscriber.OnChange(ctx, e); err != nil {
				select {
				case <-a.ctx.Done():
				case a.errC <- &observerWithErr{
					o:   subscriber,
					err: err,
				}:
				}
			}
		}()
	}
}

func (a *AsyncEventBus) handleErr() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case resp := <-a.errC:
			// 处理 publish 失败的 observer
			fmt.Printf("observer: %v, err: %v", resp.o, resp.err)
		}
	}
}
