package aigc_retry

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learn/tools/aigc_retry/business"
	"go-learn/tools/aigc_retry/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

const (
	AttemptMetadataKey = "x-retry-attempty"
)

// 退避策略下的等待
func waitRetryBackoff(attempt uint, parentCtx context.Context, callOpts *options) error {
	var waitTime time.Duration = 0
	if attempt > 0 {
		waitTime = callOpts.backoffFunc(parentCtx, attempt)
	}
	if waitTime > 0 {
		fmt.Sprintf("model attempt: %d, backoff for %v", attempt, waitTime)
		timer := time.NewTimer(waitTime)
		select {
		case <-parentCtx.Done():
			timer.Stop()
			//return contextErrToGrpcErr(parentCtx.Err())
			return parentCtx.Err()
		case <-timer.C:
		}
	}
	return nil
}

//func contextErrToGrpcErr(err error) error {
//	switch err {
//	case context.DeadlineExceeded:
//		return status.Error(codes.DeadlineExceeded, err.Error())
//	case context.Canceled:
//		return status.Error(codes.Canceled, err.Error())
//	default:
//		return status.Error(codes.Unknown, err.Error())
//	}
//}

func UnaryClientInterceptor(optFuncs ...CallOption) interceptor.UnaryClientInterceptor {
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)
	return func(parentCtx context.Context, req, reply interface{}, invoker interceptor.UnaryInvoker) error {
		if intOpts.max == 0 {
			return invoker(parentCtx, req, reply)
		}
		var lastErr error
		for attempt := uint(0); attempt < intOpts.max; attempt++ {
			if err := waitRetryBackoff(attempt, parentCtx, intOpts); err != nil {
				return err
			}
			callCtx := perCallContext(parentCtx, intOpts, attempt)
			lastErr = invoker(callCtx, req, reply)
			if lastErr == nil {
				return nil
			}
			fmt.Sprintf("attempt: %d, got err: %v", attempt, lastErr)
			if isContextError(lastErr) {
				if parentCtx.Err() != nil {
					fmt.Sprintf("grpc_retry attempt: %d, parent context error: %v", attempt, parentCtx.Err())
					return lastErr
				} else if intOpts.perCallTimeout != 0 {
					// We have set a perCallTimeout in the retry middleware, which would result in a context error if
					// the deadline was exceeded, in which case try again.
					fmt.Sprintf("attempt: %d, context error from retry call", attempt)
					continue
				}
			}
			if !isRetriable(lastErr, intOpts) {
				return lastErr
			}
		}
		return lastErr
	}
}

func AigcInterceptor(optFuncs ...CallOption) interceptor.HandlerFunc {
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)
	return func(parentCtx context.Context, c *gin.Context, taskInfo *business.Info, ret bool, invoker interceptor.Invoker) error {
		var err error
		if intOpts.max == 0 {
			err = invoker(parentCtx, c, taskInfo, ret)
			fmt.Println(ret)
			return err
		}
		var lastErr error
		for attempt := uint(0); attempt < intOpts.max; attempt++ {
			if err := waitRetryBackoff(attempt, parentCtx, intOpts); err != nil {
				return err
			}
			callCtx := perCallContext(parentCtx, intOpts, attempt)
			lastErr = invoker(callCtx, c, taskInfo, ret)
			if lastErr == nil {
				return nil
			}
			fmt.Sprintf("attempt: %d, got err: %v", attempt, lastErr)
			if isContextError(lastErr) {
				if parentCtx.Err() != nil {
					fmt.Sprintf("model attempt: %d, parent context error: %v", attempt, parentCtx.Err())
					return lastErr
				} else if intOpts.perCallTimeout != 0 {
					fmt.Sprintf("model attempt: %d, context error from retry call", attempt)
					continue
				}
			}
			if !isRetriable(lastErr, intOpts) {
				return lastErr
			}
		}
		return lastErr
	}
}

//func AigcInterceptor(optFuncs ...CallOption) interceptor.AigcInterceptor {
//	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)
//	return func(ctx context.Context, c *gin.Context, taskInfo business.Info, invoker interceptor.Invoker) error {
//		var ret bool
//		var err error
//		if intOpts.max == 0 {
//			ret, err = invoker(c, taskInfo)
//			fmt.Println(ret)
//			return err
//		}
//		var lastErr error
//		for attempt := uint(0); attempt < intOpts.max; attempt++ {
//			if err := waitRetryBackoff(attempt, c, intOpts); err != nil {
//				return err
//			}
//			//callCtx := perCallContext(c, intOpts, attempt)
//			ret, lastErr = invoker(c, taskInfo)
//			if lastErr == nil {
//				return nil
//			}
//			fmt.Sprintf("attempt: %d, got err: %v", attempt, lastErr)
//			if isContextError(lastErr) {
//				if parentCtx.Err() != nil {
//					fmt.Sprintf("grpc_retry attempt: %d, parent context error: %v", attempt, parentCtx.Err())
//					return lastErr
//				} else if intOpts.perCallTimeout != 0 {
//					// We have set a perCallTimeout in the retry middleware, which would result in a context error if
//					// the deadline was exceeded, in which case try again.
//					fmt.Sprintf("attempt: %d, context error from retry call", attempt)
//					continue
//				}
//			}
//			if !isRetriable(lastErr, intOpts) {
//				return lastErr
//			}
//		}
//		return lastErr
//	}
//}

// TODO 待改造
func isContextError(err error) bool {
	code := status.Code(err)
	return code == codes.DeadlineExceeded || code == codes.Canceled
}

func isRetriable(err error, callOpts *options) bool {
	var respErr ErrorCode
	if isContextError(err) {
		return false
	}
	switch err.(type) {
	case ErrorCode:
		errors.As(err, &respErr)
	default:
		return false
	}
	for _, code := range callOpts.codes {
		if code.Code == respErr.Code {
			return true
		}
	}
	return false
}

type UnaryServerInfo struct {
}

func ChainUnaryInterceptors(interceptors []interceptor.UnaryClientInterceptor) interceptor.UnaryClientInterceptor {
	return func(ctx context.Context, req, reply any, handler interceptor.UnaryInvoker) error {
		return interceptors[0](ctx, req, reply, getChainUnaryHandler(interceptors, 0, reply, handler))
	}
}

// ctx context.Context, req, reply any, invoker UnaryInvoker, opts ...aigc_retry.CallOption
func getChainUnaryHandler(interceptors []interceptor.UnaryClientInterceptor, curr int, reply any, finalHandler interceptor.UnaryInvoker) interceptor.UnaryInvoker {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req, resp any) error {
		return interceptors[curr+1](ctx, req, reply, getChainUnaryHandler(interceptors, curr+1, reply, finalHandler))
	}
}

func ChainInterceptors(interceptors []interceptor.HandlerFunc) interceptor.HandlerFunc {
	return func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool, handler interceptor.Invoker) error {
		return interceptors[0](ctx, c, taskInfo, ret, getChainHandler(interceptors, 0, handler))
	}
}

func getChainHandler(interceptors []interceptor.HandlerFunc, curr int, finalHandler interceptor.Invoker) interceptor.Invoker {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool) error {
		return interceptors[curr+1](ctx, c, taskInfo, ret, getChainHandler(interceptors, curr+1, finalHandler))
	}
}

type HandlersChain []interceptor.HandlerFunc

func Recovery() interceptor.HandlerFunc {
	return RecoveryWithWriter()
}

func defaultHandleRecovery(c *gin.Context, _ any) {
	c.AbortWithStatus(http.StatusInternalServerError)
}
func RecoveryWithWriter(recovery ...RecoveryFunc) interceptor.HandlerFunc {
	if len(recovery) > 0 {
		return CustomRecoveryWithWriter(recovery[0])
	}
	return CustomRecoveryWithWriter(defaultHandleRecovery)
}

type RecoveryFunc func(c *gin.Context, err any)

func CustomRecoveryWithWriter(handle RecoveryFunc) interceptor.HandlerFunc {
	return func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool, invoker interceptor.Invoker) error {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				//if logger != nil {
				//stack := stack(3)
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				//headersToStr := strings.Join(headers, "\r\n")
				//if brokenPipe {
				//	logger.Printf("%s\n%s%s", err, headersToStr, reset)
				//} else if IsDebugging() {
				//	logger.Printf("[Recovery] %s panic recovered:\n%s\n%s\n%s%s",
				//		timeFormat(time.Now()), headersToStr, err, stack, reset)
				//} else {
				//	logger.Printf("[Recovery] %s panic recovered:\n%s\n%s%s",
				//		timeFormat(time.Now()), err, stack, reset)
				//}
				//}
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
				} else {
					handle(c, err)
				}
			}
		}()
		return nil
	}
}

func New(c *gin.Context) *AIGCEngine {
	return &AIGCEngine{
		AIGCConfig{
			c:        c,
			ctx:      context.Background(),
			taskInfo: &business.Info{},
		},
		nil,
	}
}

type AIGCEngine struct {
	AIGCConfig
	Handler interceptor.HandlerFunc
}

type AIGCConfig struct {
	Handlers HandlersChain
	ctx      context.Context
	c        *gin.Context
	taskInfo *business.Info
	ret      bool
	invoker  interceptor.Invoker
}

func (ai *AIGCConfig) Use(middleware ...interceptor.HandlerFunc) {
	ai.Handlers = append(ai.Handlers, middleware...)
}
func (engine *AIGCEngine) Use(middleware ...interceptor.HandlerFunc) {
	engine.AIGCConfig.Use(middleware...)
}
