package aigc_retry

import (
	"context"
	"fmt"
	"time"
)

var (
	ErrorNeedRetry = ErrorCode{
		Message: "api need retry",
		Code:    "Retry",
	}
)

type ErrorCode struct {
	Code       string `json:"code"`
	HttpStatus int    `json:"http_code"`
	Message    string `json:"message"`
}

func (e ErrorCode) Error() string {
	return fmt.Sprintf("[%v]%v", e.Code, e.Message)
}

var (
	DefaultRetrieAbleErrorCode = []ErrorCode{ErrorNeedRetry}
	defaultOptions             = &options{
		max:            0, // disabled
		perCallTimeout: 0, // disabled
		codes:          DefaultRetrieAbleErrorCode,
		backoffFunc: BackoffFuncContext(func(ctx context.Context, attempt uint) time.Duration {
			return BackoffLinearWithJitter(50*time.Millisecond, 0.10)(attempt)
		}),
	}
)

type BackoffFuncContext func(ctx context.Context, attempt uint) time.Duration

type BackoffFunc func(attempt uint) time.Duration

type options struct {
	max            uint               //最大重试次数
	perCallTimeout time.Duration      //每次调用的超时时间限制
	codes          []ErrorCode        //重试错误
	backoffFunc    BackoffFuncContext //退避策略 backoff

	Routing string //智能路由规则（比如重试失败后开启异步重试，重新放入队列）
}

type CallOption struct {
	EmptyCallOption // make sure we implement private after() and before() fields so we don't panic.
	applyFunc       func(opt *options)
}

type EmptyCallOption struct{}

func Disable() CallOption {
	return WithMax(0)
}

func WithMax(maxRetries uint) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.max = maxRetries
	}}
}

func WithBackoff(bf BackoffFunc) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.backoffFunc = BackoffFuncContext(func(ctx context.Context, attempt uint) time.Duration {
			return bf(attempt)
		})
	}}
}

// WithBackoffContext 设置用于控制重试之间的时间
func WithBackoffContext(bf BackoffFuncContext) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.backoffFunc = bf
	}}
}

func WithCodes(retryCodes ...ErrorCode) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.codes = retryCodes
	}}
}

func WithPerRetryTimeout(timeout time.Duration) CallOption {
	return CallOption{applyFunc: func(o *options) {
		o.perCallTimeout = timeout
	}}
}

//func filterCallOptions(callOptions []CallOption) (retryOptions []CallOption) {
//	for _, opt := range callOptions {
//		if co, ok := opt.(CallOption); ok {
//			retryOptions = append(retryOptions, co)
//		}
//	}
//	return retryOptions
//}

func reuseOrNewWithCallOptions(opt *options, callOptions []CallOption) *options {
	if len(callOptions) == 0 {
		return opt
	}
	optCopy := &options{}
	*optCopy = *opt
	for _, f := range callOptions {
		f.applyFunc(optCopy)
	}
	return optCopy
}
func perCallContext(parentCtx context.Context, callOpts *options, attempt uint) context.Context {
	ctx := parentCtx
	if callOpts.perCallTimeout != 0 {
		ctx, _ = context.WithTimeout(ctx, callOpts.perCallTimeout)
	}
	//if attempt > 0 {
	//	mdClone := metautils.ExtractOutgoing(ctx).Clone().Set(AttemptMetadataKey, fmt.Sprintf("%d", attempt))
	//	ctx = mdClone.ToOutgoing(ctx)
	//}
	return ctx
}
