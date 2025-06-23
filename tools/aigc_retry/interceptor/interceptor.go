package interceptor

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-learn/tools/aigc_retry/business"
)

type UnaryClientInterceptor func(ctx context.Context, req, reply any, invoker UnaryInvoker) error

type UnaryInvoker func(ctx context.Context, req, reply any) error

type HandlerFunc func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool, invoker Invoker) error
type Invoker func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool) error
