算法题的方法套路
1、前缀和

2、滚动数组

3、双指针(快慢指针、滑动窗口)
快慢指针规律:案例:寻找重复数、环形链表等

滑动窗口规律
```
left:=0
for ( right:=0;i<上限;right++){
   //right加入窗口
   for(满足条件){
    //根据情况收集答案
    //left移除窗口
    left++
   }
   //不满足条件（根据情况收集答案）
}
```

4、栈和队列:字符匹配、
```
stack:=[]int{}
//1放入栈中,可以根据一定条件来抵消或弹出栈中数据(比如当前值比栈顶元素大的,和栈顶元素相等的等)
//2将不符合条件的移除
//3收集符合结果的值
```
5、贪心

6、动态规划
背包问题:/Users/huangzehui/learn/learn/k-points/leetcode/dynamic_programming/d.md

7、单调栈:场景一般是找到右边(或左边)第一个比他大(或小)的值
```
	//定义单调栈
	stack := []int{}
	for i := 0; i < 上限; i++ {
		//获取当前遍历的值,并和栈顶的值进行对比（栈里存储的是数组下标）
		for 单调栈>0 && 当前值 > 栈顶元素 {
			//获取栈顶的数据
			index := stack[len(stack)-1]
			//弹出栈顶数据
			stack = stack[:len(stack)-1]
			//根据问题收集结果
			//ans...
		}
		//入栈
		stack = append(stack, i)
	}
```
8、回溯:决策树、扩散问题
```
void backtracking(路径，选择列表) {
    if (终止条件) {
        存放结果;
        return;
    }
    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
```

9、图论:
```
unions := make([]int, m*n)
	for i := range unions {
		unions[i] = i
	}
	var find func(parent []int, x int) int
	find = func(parent []int, x int) int {
		if parent[x] != x {
			parent[x] = find(parent, parent[x])
		}
		return parent[x]
	}
	var union func(parent []int, x, y int)
	union = func(parent []int, x, y int) {
		xx := find(parent, x)
		yy := find(parent, y)
		if xx != yy {
			parent[xx] = yy
		}
	}
```