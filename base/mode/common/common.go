package common

import "context"

type InvokableTool interface {
	InvokableRun(ctx context.Context, argumentsInJSON string, opts ...Option2) (string, error)
}

type Option2 struct {
	implSpecificOptFn any
}

// InvokableToolEndpoint 定义 函数类型 :最终要执行的函数类型
type InvokableToolEndpoint func(ctx context.Context, input *string) (*string, error)

// InvokableToolMiddleware 定义函数类型,中间件的函数,要求:入参和出参都是最终要执行的函数类型
type InvokableToolMiddleware func(InvokableToolEndpoint) InvokableToolEndpoint
