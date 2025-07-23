package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//货币单位换算
/*
*
题目描述
记账本上记录了若干条多国货币金额，需要转换成人民币分（fen），汇总后输出。
每行记录一条金额，金额带有货币单位，格式为数字+单位，可能是单独元，或者单独分，或者元与分的组合。
要求将这些货币全部换算成人民币分（fen）后进行汇总，汇总结果仅保留整数，小数部分舍弃。
元和分的换算关系都是1:100，如下：

1CNY=100fen（1元=100分）
1HKD=100cents（1港元=100港分）
1JPY=100sen（1日元=100仙）
1EUR=100eurocents（1欧元=100欧分）
1GBP=100pence（1英镑=100便士）
汇率表如下：

即：100CNY = 1825JPY = 123HKD = 14EUR = 12GBP

CNY	JPY	HKD	EUR	GBP
100	1825	123	14	12
输入描述
第一行输入为N，N表示记录数。0<N<100

之后N行，每行表示一条货币记录，且该行只会是一种货币。

输出描述
将每行货币转换成人民币分（fen）后汇总求和，只保留整数部分。
输出格式只有整数数字，不带小数，不带单位。s
*/

/*
*
输入
2
20CNY53fen
53HKD87cents

输出
6432
*/
func main11() {
	//思路：1、定义方法：通过货币单位获取人民币分的汇率
	//2、逐行解析字符串s：如果是数子：c>='0'&&c<='9' 则current=current*10+int(c-'0')，-》得出金额count
	//3、如果不是数字，则从字符串解析出来(以字符串结尾或者后面又出现数字为截止)，->得出汇率 rate
	//4、换算，将count*rate，然后加到总和结果中
	var exChange func(unit string) float64
	exChange = func(unit string) float64 {
		switch unit {
		case "CNY":
			return 100.0 // 人民币
		case "JPY":
			return 100.0 / 1825 * 100 // 日元
		case "HKD":
			return 100.0 / 123 * 100 // 港元
		case "EUR":
			return 100.0 / 14 * 100 // 欧元
		case "GBP":
			return 100.0 / 12 * 100 // 英镑
		case "fen":
			return 1.0 // 人民币分
		case "cents":
			return 100.0 / 123 // 港元分
		case "sen":
			return 100.0 / 1825 // 日元分
		case "eurocents":
			return 100.0 / 14 // 欧元分
		case "pence":
			return 100.0 / 12 // 英镑分
		default:
			return 0.0 // 无效单位
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	totalFen := 0.0

	for i := 0; i < n; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		currentFen := 0
		//临时存储金币单位字符串
		currentUnit := ""
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c >= '0' && c <= '9' {
				currentFen = currentFen*10 + int((c - '0'))
			} else {
				//输入
				//2
				//20CNY53fen
				//53HKD87cents
				currentUnit = currentUnit + string(c)
				//检查是否到了结尾，或者是否是数字
				if i+1 == len(s) || (s[i+1] >= '0' && s[i+1] <= '9') {
					rate := exChange(currentUnit)
					f := float64(currentFen) * rate
					totalFen += f
					currentUnit = ""
					currentFen = 0.0
				}
			}

		}

	}
	fmt.Println(int(totalFen))
}
