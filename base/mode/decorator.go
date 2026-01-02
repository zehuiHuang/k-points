package mode

import (
	"context"
	"fmt"
	"go-learn/base/mode/common"
)

// 装饰器模式，经典案例：grpc拦截器实现，经过层层装饰，将整个执行链路压缩成一个拦截器UnaryServerInterceptor

// 装饰器模式的另外一种实现方案

type Weather2 struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func (w Weather2) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...common.Option2) (string, error) {
	fmt.Println("---------------call Weather tool---------------------")
	return "{\"temperature\": 25.0, \"humidity\": 60.0}", nil
}

// Invoke 定义 Invoke类型的函数
type Invoke[I, O, TOption any] func(ctx context.Context, input I, opts ...TOption) (output O, err error)

func wrapToolCall2(it common.InvokableTool, middlewares []common.InvokableToolMiddleware, needCallback bool) common.InvokableToolEndpoint {
	if needCallback {
		it = &invokableToolWithCallback{it: it}
	}
	middleware := func(next common.InvokableToolEndpoint) common.InvokableToolEndpoint {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}

	return middleware(func(ctx context.Context, input *string) (*string, error) {
		fmt.Println("---------------call tool---------------------")
		result, err := it.InvokableRun(ctx, *input)
		fmt.Println("---------------call tool---------------------")
		if err != nil {
			return nil, err
		}
		return &result, nil
	})
}

type invokableToolWithCallback struct {
	it common.InvokableTool
}

func (i *invokableToolWithCallback) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...common.Option2) (string, error) {
	//相当于又做了一次函数增强,强执行函数的前后进行处理
	f := invokeWithCallbacks(i.it.InvokableRun)
	return f(ctx, argumentsInJSON)
}

func invokeWithCallbacks[I, O, TOption any](i Invoke[I, O, TOption]) Invoke[I, O, TOption] {
	return runWithCallbacks(i, onStart[I], onEnd[O], onError)
}

func runWithCallbacks[I, O, TOption any](r func(context.Context, I, ...TOption) (O, error),
	onStart on[I], onEnd on[O], onError on[error]) func(context.Context, I, ...TOption) (O, error) {
	return func(ctx context.Context, input I, opts ...TOption) (output O, err error) {
		//开始
		ctx, input = onStart(ctx, input)
		output, err = r(ctx, input, opts...)
		if err != nil {
			//错误
			ctx, err = onError(ctx, err)
			return output, err
		}
		//执行后
		ctx, output = onEnd(ctx, output)
		return output, nil
	}
}

// 定义函数类型(带泛型)
type on[T any] func(context.Context, T) (context.Context, T)

func onStart[T any](ctx context.Context, input T) (context.Context, T) {
	//return icb.On(ctx, input, icb.OnStartHandle[T], callbacks.TimingOnStart, true)
	fmt.Println("开始---------,入参为:", input)
	return context.Background(), input
}

func onEnd[T any](ctx context.Context, output T) (context.Context, T) {
	//return icb.On(ctx, output, icb.OnEndHandle[T], callbacks.TimingOnEnd, false)
	fmt.Println("结束---------,出参为:", output)
	return context.Background(), output
}
func onError(ctx context.Context, err error) (context.Context, error) {
	//return icb.On(ctx, err, icb.OnErrorHandle, callbacks.TimingOnError, false)
	fmt.Println("错误---------,错误信息为:", err)
	return context.Background(), nil
}

//----------------------------------

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
