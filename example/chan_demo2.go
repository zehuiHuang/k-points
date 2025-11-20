package example

import (
	"fmt"
	"time"
)

/*
*
模拟生产者和消费者

主流程被done通道阻塞(可以利用sync.GroupWait),等待生产者都发送完之后再close掉jobs通道
消费者异步实时监听jobs通道,若检测到关闭,则消费者也主动退出,推出时想done发送信号,主流程被done通道阻塞被放开,这个执行完毕

根据执行结果判定:失败了三个,被丢弃了
*/
func producerConsumer() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 生产者
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case jobs <- i:
				fmt.Printf("Produced job %d\n", i)
			case <-time.After(100 * time.Millisecond):
				fmt.Println("Producer timeout")
			}
		}
		close(jobs)
	}()

	// 消费者
	go func() {
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					done <- true
					return
				}
				fmt.Printf("Consumed job %d\n", job)
				time.Sleep(200 * time.Millisecond) // 模拟处理时间
			case <-time.After(300 * time.Millisecond):
				fmt.Println("Consumer waiting...")
			}
		}
	}()

	<-done
}

/*
*执行结果:
Produced job 0
Produced job 1
Produced job 2
Produced job 3
Produced job 4
Produced job 5
Consumed job 0
Producer timeout
Consumed job 1
Produced job 7
Producer timeout
Producer timeout
Consumed job 2
Consumed job 3
Consumed job 4
Consumed job 5
Consumed job 7
*/
