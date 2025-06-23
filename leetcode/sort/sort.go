package sort

/*
插入排序：从第二个数开始，往前找，找到比自己大的数，插入到这个数的后面
*/

func insertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	/*
				1 6 3 9 0 6 3 2

			第一轮  从第二位开始，往前找：6 比 1 大，不用管 已排序1 6 未排序 3 9 0 6 3 2
			第二轮 从第三位开始，往前找： 3 比 6 小，交换位置，继续往前找：3 比 1 大，交换位置，已排序1 3 6 未排序 9 0 6 3 2
		    以此类推
	*/
	//0 1
	//0 2
	//0 3
	//0 n-1
	N := len(arr) - 1
	//下标从1开始
	for i := 1; i <= N; i++ {
		//从i开始往前找，找到比自己大的数，插入到这个数的后面
		currentIndex := i
		for currentIndex == 0 || arr[currentIndex-1] > arr[currentIndex] {
			arr[currentIndex-1], arr[currentIndex] = arr[currentIndex], arr[currentIndex-1]
			currentIndex--
		}
	}
	return arr

}

/*
*
冒泡排序 最大的从最左边冒泡到最右边（右边的就不用管了），继续找第二大的，以此类推
*/
func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	//1 6 3 9 0 6 3 2
	/*
		0 n-1
		0 n-2
		....
		0 end
	*/
	N := len(arr)
	//下标
	for end := N - 1; end >= 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

/*
* 找到最小的值，放到最左边（最左边的第一位就不用管了），继续从第二位～n位，找到最小的放到最左边的第二位（最左边的第二位就不用管了），以此类推
 */
func selectSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	//1 6 3 9 0 6 3 2
	//0 1 2 3 4 5 6 7
	/**

	0 n-1
	1 n-1
	...
	n-2 n-1

	0  1 6 3 9 6 3 2
	0  1 6 3 9 6 3 2
	0  1 2 6 3 9 6 3
	*/
	N := len(arr)
	for i := 0; i <= N-1; i++ {
		//找到最小值的下标，然后和minIndex交换
		minIndex := i
		for j := i; j <= N-1; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}

/*
*
要求对arr做到左边的都大于k，右边的都小于k，中看的等于k
*/
func sortTest(arr []int, k int) []int {
	N := len(arr)
	left := -1
	right := N
	for i := 0; i <= N-1; i++ {
		if arr[i] < k {
			left++
			arr[left], arr[i] = arr[i], arr[left]
		} else if arr[i] > k {
			right--
			arr[right], arr[i] = arr[i], arr[right]
		}
	}
	return []int{left, right}
}

/**
快速排序
*/

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := partition(arr)
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
	return arr
}

// 分区函数:找到基准元素的正确位置
func partition(arr []int) int {
	high := len(arr) - 1
	pivot := arr[high] // 选择最后一个元素作为基准
	i := -1            // 记录小于基准的位置

	// 遍历数组，将小于等于基准的元素移到左侧
	for j := 0; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准元素放到正确位置
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
