package example

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
	"time"
)

/**
相对sync.waiGroup，errgroup.Group 提供了更丰富的功能，如：
1. 错误返回(首个错误)
2. 上下文取消(其中有一个错误,则其他协程自动取消)
3.并发数控制
4.尝试启动(在并发数配置的前提下,可以感知协程启动成功与否)
*/
// 使用errgroup 进行并发控制

var urls = []string{
	"https://www.baidu.com/",
	"https://www.baidu.com/",
	"http://www.somestupidname.com/", // 这是一个错误的 URL，会导致任务失败
}

// 1、错误返回
func errGroupExample() {
	// 使用 errgroup 创建一个新的 goroutine 组
	var g errgroup.Group // 零值可用，不必显式初始化

	for _, url := range urls {
		// 使用 errgroup 启动一个 goroutine 来获取 URL
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err // 发生错误，返回该错误
			}
			defer resp.Body.Close()
			fmt.Printf("fetch url %s status %s\n", url, resp.Status)
			return nil // 返回 nil 表示成功
		})
	}

	// 等待所有 goroutine 完成并返回第一个错误（如果有）
	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

// 2、上下文取消
func errGroupExample2() {
	// 创建一个带有 context 的 errgroup
	// 任何一个 goroutine 返回非 nil 的错误，或 Wait() 等待所有 goroutine 完成后，context 都会被取消
	g, ctx := errgroup.WithContext(context.Background())

	// 创建一个 map 来保存结果
	var result sync.Map

	for _, url := range urls {
		// 使用 errgroup 启动一个 goroutine 来获取 URL
		g.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return err // 发生错误，返回该错误
			}

			// 发起请求
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err // 发生错误，返回该错误
			}
			defer resp.Body.Close()

			// 保存每个 URL 的响应状态码
			result.Store(url, resp.Status)
			return nil // 返回 nil 表示成功
		})
	}

	// 等待所有 goroutine 完成并返回第一个错误（如果有）
	if err := g.Wait(); err != nil {
		fmt.Println("Error: ", err)
	}

	// 所有 goroutine 都执行完成，遍历并打印成功的结果
	result.Range(func(key, value any) bool {
		fmt.Printf("fetch url %s status %s\n", key, value)
		return true
	})
}

// 3、限制并发数
func errGroupExample3() {
	// 创建一个 errgroup.Group
	var g errgroup.Group
	// 设置最大并发限制为 3
	g.SetLimit(3)

	// 启动 10 个 goroutine
	for i := 1; i <= 10; i++ {
		g.Go(func() error {
			// 打印正在运行的 goroutine
			fmt.Printf("Goroutine %d is starting\n", i)
			time.Sleep(2 * time.Second) // 模拟任务耗时
			fmt.Printf("Goroutine %d is done\n", i)
			return nil
		})
	}

	// 等待所有 goroutine 完成
	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered an error: %v\n", err)
	}

	fmt.Println("All goroutines complete.")
}

// 4、尝试启动
func errGroupExample4() {
	// 创建一个 errgroup.Group
	var g errgroup.Group
	// 设置最大并发限制为 3
	g.SetLimit(3)

	// 启动 10 个 goroutine
	for i := 1; i <= 10; i++ {
		if g.TryGo(func() error {
			// 打印正在运行的 goroutine
			fmt.Printf("Goroutine %d is starting\n", i)
			time.Sleep(2 * time.Second) // 模拟工作
			fmt.Printf("Goroutine %d is done\n", i)
			return nil
		}) {
			// 如果成功启动，打印提示
			fmt.Printf("Goroutine %d started successfully\n", i)
		} else {
			// 如果达到并发限制，打印提示
			fmt.Printf("Goroutine %d could not start (limit reached)\n", i)
		}
	}

	// 等待所有 goroutine 完成
	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered an error: %v\n", err)
	}

	fmt.Println("All goroutines complete.")
}
