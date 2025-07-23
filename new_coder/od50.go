package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//整理扑克牌
/*
*
题目
给定一组数字，表示扑克牌的牌面数字，忽略扑克牌的花色，请按如下规则对这一组扑克牌进行整理：步骤1. 对扑克牌进行分组，形成组合牌，规则如下：当牌面数字相同张数大于等于4时，组合牌为“炸弹”； 3张相同牌面数字 + 2张相同牌面数字，且3张牌与2张牌不相同时，组合牌为“葫芦”； 3张相同牌面数字，组合牌为“三张”； 2张相同牌面数字，组合牌为“对子”； 剩余没有相同的牌，则为“单张”； 步骤2. 对上述组合牌进行由大到小排列，规则如下：不同类型组合牌之间由大到小排列规则：“炸弹” > “葫芦” > “三张” > “对子” > “单张”； 相同类型组合牌之间，除“葫芦”外，按组合牌全部牌面数字加总由大到小排列； “葫芦”则先按3张相同牌面数字加总由大到小排列，3张相同牌面数字加总相同时，再按另外2张牌面数字加总由大到小排列； 由于“葫芦”>“三张”，因此如果能形成更大的组合牌，也可以将“三张”拆分为2张和1张，其中的2张可以和其它“三张”重新组合成“葫芦”，剩下的1张为“单张” 步骤3. ​当存在多个可能组合方案时，按如下规则排序取最大的一个组合方案：依次对组合方案中的组合牌进行大小比较，规则同上； 当组合方案A中的第n个组合牌大于组合方案B中的第n个组合牌时，组合方案A大于组合方案B；

输入描述
第一行为空格分隔的N个正整数，每个整数取值范围[1,13]，N的取值范围[1,1000]

输出描述
经重新排列后的扑克牌数字列表，每个数字以空格分隔
*/

/*
*输入：
1 3 3 3 2 1 5
输出：
3 3 3 1 1 5 2
*/
func main() {
	/**思路：
	1、先通过map结构将各个卡牌上的数出现的次数进行统计
	2、根据1次、2次、3次、4次以上的四个维度进行分组，将牌放到四个不同的组内
	3、每个组按照顺序排序：
	   1）出现四次以上的先按照出现次数进行排序，出现次数一样的则按照牌面大小进行排序（生序排列）
	   2）对出现1、2、3张的都按照牌面进行生序排序
	4、创建一个结果集，将牌按照一定规则放进去
	   1）先将大于等于4的按照排号的放进去
	   2）在将牌数为3的放进去，并且尽量和出现次数为2的进行配对，
	      配对规则为：若后面还有3张的牌的前提下，且没有出现2张的（或者有2张的但是小于出现3张的牌面），则对3张的牌进行拆分（AA、A），将AA和3张的牌配对形成葫芦，另外的A放到单张去
	                若后面没有3张的，则顺便把两张的放入结果集中
	   3）将2张的按照排好的顺序进行放入
	   4）将1张的按照排好的顺序放入
	*/

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	tables := make([]int, len(input))

	mp := map[int]int{}
	for i := range tables {
		tables[i], _ = strconv.Atoi(input[i])
		mp[tables[i]] = mp[tables[i]] + 1
	}

	z1 := []int{}
	z2 := []int{}
	z3 := []int{}
	z4 := []int{}

	for k, v := range mp {
		if v == 1 {
			z1 = append(z1, k)
		} else if v == 2 {
			z2 = append(z2, k)
		} else if v == 3 {
			z3 = append(z3, k)
		} else {
			z4 = append(z4, k)
		}
	}

	//排序

	sort.Slice(z1, func(i, j int) bool {
		return z1[i] < z1[j]
	})
	sort.Slice(z2, func(i, j int) bool {
		return z2[i] < z2[j]
	})
	sort.Slice(z3, func(i, j int) bool {
		return z3[i] < z3[j]
	})
	sort.Slice(z4, func(i, j int) bool {
		//张数相同，则比牌面大小
		if mp[z4[i]] == mp[z4[j]] {
			return z4[i] < z4[j]
		} else {
			//张数不同则比牌数量s
			return mp[z4[i]] < mp[z4[j]]
		}
	})
	ans := []int{}
	//按照规则，先放入>=4张的牌
	for i := range z4 {
		count := mp[z4[i]]
		for j := 0; j < count; j++ {
			ans = append(ans, z4[i])
		}
	}
	//放入3张的
	for len(z3) > 0 {
		v := z3[len(z3)-1]
		z3 = z3[:len(z3)-1]
		ans = append(ans, v, v, v)
		//对下一个3张的进行拆分,比如：AAA-> AA+A
		if len(z3) > 0 && (len(z2) == 0 || z3[len(z3)-1] > z2[len(z2)-1]) {
			v2 := z3[len(z3)-1]
			z3 = z3[:len(z3)-1] //移除掉3张牌
			ans = append(ans, v2, v2)
			//单张的放入单张的集合
			z1 = append(z1, v2)
		} else {
			if len(z2) > 0 {
				v2 := z2[len(z2)-1]
				z2 = z2[:len(z2)-1]
				ans = append(ans, v2, v2)
			}
		}
	}

	for len(z2) > 0 {
		v2 := z2[len(z2)-1]
		z2 = z2[:len(z2)-1]
		ans = append(ans, v2, v2)
	}
	//对z1在进行升序排序（因为可能有从3张拆出来的放入了z1）
	sort.Slice(z1, func(i, j int) bool {
		return z1[i] < z1[j]
	})
	for len(z1) > 0 {
		v1 := z1[len(z1)-1]
		z1 = z1[:len(z1)-1]
		ans = append(ans, v1)
	}
	ret := make([]string, len(ans))
	for i := range ans {
		ret[i] = strconv.Itoa(ans[i])
	}
	fmt.Println(strings.Join(ret, " "))
}
