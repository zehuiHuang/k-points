package main

import (
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"time"
)

func a() {
	retryOps := []grpc_retry.CallOption{
		grpc_retry.WithMax(2),
		grpc_retry.WithPerRetryTimeout(time.Second * 2),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinearWithJitter(time.Second/2, 0.2)),
	}
	grpc_retry.UnaryClientInterceptor(retryOps...)

}
