package practice

/*
*

过程演练:
digits=“23” -》 abc def
a  b  c

a -> ad ae af
b -> bd be bf
c -> cd ce cf

思路:使用回溯算法
1、定义dfs(i,path),i从1~开始到n结束,表示对第一个数字对应的字符串遍历(比如abc)
a ,  b  , c
2、拼接a,拼接b,拼接c ,这三个分支在分别对第二个数组对应的字符串(比如def)进行遍历
例如:以a分支为起始分别会和d、e、f进行拼接得到 ad,ad,af
以b分支为起始分别会和d、e、f进行拼接得到 bd,bd,bf
以此类推进行递归
3、直到i==n,说明有符合结果预期的,则进行数据收集
4、最后别忘了dfs方法的调用
5、因为使用了append(path, byte(c),所有不需要恢复
*/
func letterCombinations(digits string) []string {
	n := len(digits)
	var mapping = [10]string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var dfs func(int, []byte)
	ans := []string{}
	dfs = func(i int, path []byte) {
		// 终止条件和结果收集
		if i == n {
			ans = append(ans, string(path))
			return
		}
		//digits[i]表示当前的数字
		//digits[i]-'0'表示当前数字的索引
		//mapping[digits[i]-'0']表示当前数字对应的字母,比如2对应的字母是abc
		for _, c := range mapping[digits[i]-'0'] {
			dfs(i+1, append(path, byte(c)))
		}
	}
	dfs(0, []byte{})
	return ans
}
