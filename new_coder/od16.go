package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//数组二叉树
/**
题目描述
二叉树也可以用数组来存储，给定一个数组，树的根节点的值存储在下标1，对于存储在下标N的节点，它的左子节点和右子节点分别存储在下标2*N和2*N+1，并且我们用值-1代表一个节点为空。

给定一个数组存储的二叉树，试求从根节点到最小的叶子节点的路径，路径由节点的值组成。

输入描述
输入一行为数组的内容，数组的每个元素都是正整数，元素间用空格分隔。

注意第一个元素即为根节点的值，即数组的第N个元素对应下标N，下标0在树的表示中没有使用，所以我们省略了。

输入的树最多为7层。

输出描述
输出从根节点到最小叶子节点的路径上，各个节点的值，由空格分隔，用例保证最小叶子节点只有一个。
*/

/*
*
输入：
3 5 7 -1 -1 2 4
输出：
3 7 2
*/

func main16() {
	//思路，找到最小子节点，然后倒推到根节点
	//1、循环遍历数组，找到是有效的页子节点（首先它不等于-1，其次它没有左右子节点，那么它一定是页子节点），相比较并选择最小的一个叶子节点对应的数组下标
	//2、找到叶子节点后，进行循环遍历for idx!=0 ,根据idx计算出父节点[由z=2*i+1=》(z-1)/2即为父节点下索引坐标]，直到父节点的下标为0
	//3、将切片进行排序反转
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	arr := make([]int, len(input))
	for i := range input {
		arr[i], _ = strconv.Atoi(input[i])
	}
	n := len(arr)
	idx := -1
	minV := math.MaxInt64
	for i := 0; i < n; i++ {
		//过滤掉不是子节点的
		if arr[i] == -1 {
			continue
		}
		// 计算左右子节点索引
		left := 2*i + 1
		right := 2*i + 2

		// 检查是否为叶子节点：左右子节点均不存在或为空
		isLeftExist := left < n && arr[left] != -1
		isRightExist := right < n && arr[right] != -1
		if !isLeftExist && !isRightExist {
			if minV > arr[i] {
				minV = arr[i]
				idx = i
			}
		}
	}
	ans := []int{arr[idx]}
	//只要不是root节点则继续遍历
	for idx != 0 {
		//根据idx=2*p+1 推到出p=(idx-1)/2 左右子节点都符合条件
		idx = (idx - 1) / 2
		ans = append(ans, arr[idx])
	}
	strArr := make([]string, len(ans))
	k := len(ans)
	for i := range ans {
		strArr[k-i-1] = strconv.Itoa(ans[i])
	}

	fmt.Println(strings.Join(strArr, " "))
}
