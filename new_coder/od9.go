package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*题目
让我们来模拟一个消息队列的运作，有一个发布者和若干消费者，发布者会在给定的时刻向消息队列发送消息，若此时消息队列有消费者订阅，这个消息会被发送到订阅的消费者中优先级最高（输入中消费者按优先级升序排列）的一个； 若此时没有订阅的消费者，该消息被消息队列丢弃。 消费者则会在给定的时刻订阅消息队列或取消订阅。当消息发送和订阅发生在同一时刻时，先处理订阅操作，即同一时刻订阅的消费者成为消息发送的候选。 当消息发送和取消订阅发生在同一时刻时，先处理取消订阅操作，即消息不会被发送到同一时刻取消订阅的消费者。

输入描述
输入为两行。

第一行为2N个正整数，代表发布者发送的N个消息的时刻和内容（为方便解折，消息内容也用正整数表示）。第一个数字是第一个消息的发送时刻，第二个数字是第一个消息的内容，以此类推。用例保证发送时刻不会重复，但注意消息并没有按照发送时刻排列。

第二行为2M个正整数，代表M个消费者订阅和取消订阅的时刻。第一个数字是第一个消费者订阅的时刻，第二个数字是第一个消费者取消订阅的时刻，以此类推。用例保证每个消费者的取消订阅时刻大于订阅时刻，消费者按优先级升序排列。

两行的数字都由空格分隔。N不超过100，M不超过10，每行的长度不超过1000字符

输出描述
输出为M行，依次为M个消费者收到的消息内容，消息内容按收到的顺序排列，且由空格分隔；

若某个消费者没有收到任何消息，则对应的行输出-1
*/

/*
*输入：
2 22 1 11 4 44 5 55 3 33
1 7 2 3
输出：
11 33 44 55
22
*/
func main9() {
	//思路：1、将生产者和消费者组装成二位数组(生产者)，同时对生产者按照消息发送时间进行升序排序
	//2、按照顺序遍历生产者，拿到生产者的发送时间和内容后，开始对消费者倒叙遍历，若生产者发送时间在消费者时间范围内，则将消息发送的内容存储起来
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	producer := strings.Fields(scanner.Text())
	scanner.Scan()
	consumer := strings.Fields(scanner.Text())
	//生产者队列
	pArr := make([][]int, len(producer)/2)
	pIdx := 0
	i := 0
	for i < len(producer) {
		time, _ := strconv.Atoi(producer[i])
		content, _ := strconv.Atoi(producer[i+1])
		pArr[pIdx] = []int{time, content}
		pIdx++
		i += 2
	}

	//消费者队列
	cArr := make([][]int, len(consumer)/2)

	cIdx := 0
	j := 0
	for j < len(consumer) {
		start, _ := strconv.Atoi(consumer[j])
		end, _ := strconv.Atoi(consumer[j+1])
		cArr[cIdx] = []int{start, end}
		cIdx++
		j += 2
	}
	//生产者按照发送时间升序排序，模拟消息发送的顺序性
	sort.Slice(pArr, func(i, j int) bool {
		return pArr[i][0] < pArr[j][0]
	})
	//定义消费者消费的消息
	consumerArr := make([][]int, len(cArr))
	for i := 0; i < len(consumerArr); i++ {
		consumerArr[i] = make([]int, 0)
	}

	for i := 0; i < len(pArr); i++ {
		//时间
		time := pArr[i][0]
		//消息内容
		content := pArr[i][1]
		//优先级最高的消费者(包含消费开始时间和截止时间)的下标索引
		end := len(cArr) - 1

		for end >= 0 {
			statTime := cArr[end][0]
			endTime := cArr[end][1]
			if time >= statTime && time < endTime {
				consumerArr[end] = append(consumerArr[end], content)
				break
			}
			end--
		}
	}

	for i := 0; i < len(consumerArr); i++ {
		a := consumerArr[i]
		if len(a) == 0 {
			fmt.Println(-1)
		} else {
			strs := make([]string, len(a))
			for i, msg := range a {
				strs[i] = strconv.Itoa(msg)
			}
			fmt.Println(strings.Join(strs, " "))
		}
	}

}
