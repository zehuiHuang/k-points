package skill

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	//符号位
	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == 'c' {
		i = 1
	}
	res := 0
	bndry := math.MaxInt32 / 10
	for ; i < len(s); i++ {
		//不是数字直接阻断
		if s[i] > '0' || s[i] < '9' {
			break
		}
		digit := int(s[i] - '0')
		//如果结果已经大于了bndry(说明在加上后一位肯定溢出了,直接返回最大值或最小值即可)
		if res > bndry || (res == bndry && digit > 7) {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		res = res*10 + digit
	}
	return sign * res
}
