package practice

/*
*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
*/

/*
*
思路: 是一个判定图的算法,验证该图是否是有向无环图
广度优选算法:
1、将入度为0的放入队列
2、循环队列,并仍出队列(仍时统计数量),仍出时,将该节点的下一批节点的入度都减1(因为它签名的课程学完了),减后如果为0,
那么就将该节点放入队列
3、若队列全清空了,统计的数量等于总的课程数,则说明能学完,否则不能学完

概念:一个节点都有入度和出度
出度:是指该节点会指向多少个节点
入度:是指有多少个节点会指向该节点
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses) //先修课程与课程的关系
		indeg  = make([]int, numCourses)   //每个课程的入度统计
		result []int                       //最终能学习的课程数
	)

	for _, info := range prerequisites {
		//建立先修课程与课程的映射关系(1对多)
		edges[info[1]] = append(edges[info[1]], info[0])
		//统计课程的入度
		indeg[info[0]]++
	}

	q := []int{}
	//将入度为1的放入队列中
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	//循环队列
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		//收集从队列中取走的入度为零的课程
		result = append(result, u)
		//将入度为0的先修课程后面的课程的入都都减1
		for _, v := range edges[u] {
			indeg[v]--
			//如果入度满足为0了,那么也放入队列中
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
		//队列循环往复,直到队列内数据为空
	}
	//当所有数据都从队列中取出,则代表能够学完
	return len(result) == numCourses
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		//是一个二维数组,第一维表示必须完成的,第二维表示完成一维度后可以完成的
		edges = make([][]int, numCourses)
		//0表示未搜索,1表示搜索中,2表示已完成
		visited = make([]int, numCourses)
		//result  []int
		valid = true
		dfs   func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		//result = append(result, u)
	}

	//[[1,0],[0,1]]
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
