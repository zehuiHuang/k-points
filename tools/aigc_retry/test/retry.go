package main

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/hashicorp/go-retryablehttp"
	"log"
	"math"
	"time"
)

// retryable function
func testFunc() (any, error) {
	return "lee", nil
}

func ab(attempt uint) time.Duration {
	return 60 * time.Millisecond * time.Duration(math.Pow(3, float64(attempt)))
}

func main2() {
	//// retry call
	//result := retry.DoWithDefault(testFunc)
	//
	//// result
	//fmt.Println("result:", result.Data())
	//fmt.Println("tryError:", result.TryError())
	//fmt.Println("execErrors:", result.ExecErrors())
	//fmt.Println("isSuccess:", result.IsSuccess())
	for attempt := uint(0); attempt < 4; attempt++ {
		fmt.Println(ab(attempt))
	}

	retryPolicy := `{
        "methodConfig":[{
           "name":[{"service": "grpc-tutorial.05retry.hello.hello"}],
           "waitForReady": true,
           "retryPolicy": {
                "MaxAttempts": 4,
                "InitialBakckoff": ".01s",
                "MaxBackoff": ".01s",
                "BackoffMultiplier": 1.0,
                 "RetryableStatusCodes": ["UNAVAILABLE"]
            }
    }]}`
	//conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
	fmt.Println(retryPolicy)

}

func main3() {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 100 * time.Millisecond
	b.MaxInterval = 5 * time.Second
	b.MaxElapsedTime = 30 * time.Second
	err := backoff.Retry(func() error {
		result, err := doSomething()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println("Result:", result)
		return nil
	}, b)

	if err != nil {
		fmt.Println("Exceeded maximum number of retries")
	}
}

func doSomething() (string, error) {
	// 模拟需要重试的操作
	// 这里使用一个简单的计数器来模拟操作失败的情况
	counter++
	if counter < 3 {
		fmt.Printf("-------------------%v\n", counter)
		return "", fmt.Errorf("Operation failed")
	}
	return "Operation successful", nil
}

var counter int

func main4() {
	// 创建一个重试客户端
	client := retryablehttp.NewClient()
	// 设置重试策略
	client.RetryWaitMin = 1 * time.Second
	client.RetryWaitMax = 5 * time.Second
	client.RetryMax = 3
	// 也可以设置其他的 HTTP 客户端选项，比如 Timeout
	client.HTTPClient.Timeout = 10 * time.Second
	// 发起一个 GET 请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatalf("Error making HTTP request: %s", err)
	}
	defer resp.Body.Close()
	// 检查响应状态码
	if resp.StatusCode != 200 {
		log.Fatalf("Server returned non-200 status code: %d", resp.StatusCode)
	}
	// 读取并打印响应体（这里只是打印了状态，没有读取响应体内容）
	fmt.Println("Response status:", resp.Status)
}
func main7() {
	// 创建一个重试器
	r := retrier.New(retrier.ConstantBackoff(3, 100*time.Millisecond), nil)
	// 定义一个重试任务函数
	task := func() error {
		// 模拟一个失败的任务
		fmt.Println("Performing task...")
		return errors.New("Some error occurred")
	}
	// 使用重试器执行任务
	err := r.Run(task)
	if err != nil {
		fmt.Println("Task failed:", err)
	} else {
		fmt.Println("Task completed successfully")
	}
}

func main() {

}
