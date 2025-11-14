package skill

// 思路:找规律并模拟,如果有三行,那么索引i的坐标应移动方式是 0->1->2 ->1 ->0->1->2 ->1  ...这种规律走的

func convert(s string, numRows int) string {
	res := ""
	queue := make([]string, numRows)
	if len(s) < 2 {
		return s
	}
	if numRows < 2 {
		return s
	}
	for i := 0; i < numRows; i++ {
		queue[i] = ""
	}
	//坐标:如果是3行,则可以是0,1,2
	j := 0
	flag := -1
	for i := 0; i < len(s); i++ {
		queue[j] += string(s[i])
		if j == 0 || j == numRows-1 {
			flag = -flag
		}
		j += flag
	}
	for _, v := range queue {
		res += v
	}
	return res
}
