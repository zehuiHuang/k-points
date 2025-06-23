package mode

import (
	"context"
	"fmt"
)

// 装饰器模式，经典案例：grpc拦截器实现，经过层层装饰，将整个执行链路压缩成一个拦截器UnaryServerInterceptor

type Food interface {
	Eat() string
	Cost() float32
}

// Rice 米饭
type Rice struct {
}

func NewRice() Food {
	return Rice{}
}
func (r Rice) Eat() string {
	return "rice"
}
func (r Rice) Cost() float32 {
	return 0.1
}

type Noodle struct {
}

func NewNoodle() Noodle {
	return Noodle{}
}

func (n Noodle) Eat() string {
	return "Noodle"
}
func (n Noodle) Cost() float32 {
	return 0.2
}

type Decorator Food

func NewDecorator(f Food) Decorator {
	return f
}

type LaoGanMaDecorator struct {
	Decorator
}

func NewLaoGanMaDecorator(d Decorator) Decorator {
	return &LaoGanMaDecorator{
		d,
	}
}

func (l *LaoGanMaDecorator) Eat() string {
	// 加入老干妈配料
	return "加入一份老干妈~..." + l.Decorator.Eat()
}

func (l *LaoGanMaDecorator) Cost() float32 {
	// 价格增加 0.5 元
	return 0.5 + l.Decorator.Cost()
}

type HamSausageDecorator struct {
	Decorator
}

func (h *HamSausageDecorator) Eat() string {
	// 加入老干妈配料
	return "加入一份老干妈~..." + h.Decorator.Eat()
}

func (h *HamSausageDecorator) Cost() float32 {
	// 价格增加 0.5 元
	return 0.5 + h.Decorator.Cost()
}

//函数增强：闭包实现

type handleFunc func(ctx context.Context, param map[string]interface{}) error

// Decorate 函数增强
func Decorate(fn handleFunc) handleFunc {
	return func(ctx context.Context, param map[string]interface{}) error {
		// 前处理
		fmt.Println("preprocess...")
		err := fn(ctx, param)
		fmt.Println("postprocess...")
		return err
	}
}

// 案例----------------------------grpc

type UnaryHandler func(ctx context.Context, req any) (any, error)

func chainUnaryInterceptors(interceptors []UnaryServerInterceptor) UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (any, error) {
		return interceptors[0](ctx, req, info, getChainUnaryHandler(interceptors, 0, info, handler))
	}
}

type UnaryServerInfo struct {
}
type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)

func getChainUnaryHandler(interceptors []UnaryServerInterceptor, curr int, info *UnaryServerInfo, finalHandler UnaryHandler) UnaryHandler {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req any) (any, error) {
		return interceptors[curr+1](ctx, req, info, getChainUnaryHandler(interceptors, curr+1, info, finalHandler))
	}
}

/**
理解：其实就是函数增强，chainUnaryInterceptors 函数执行时，首先执行第一个Interceptor的方法，然后通过getChainUnaryHandler方法的执行，继续对handler进行包装

拦截器链中的每一个拦截器 UnaryServerInterceptor 可以理解为一个装饰器
*/
