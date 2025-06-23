package law

import (
	"container/heap"
	"math"
)

// Item 结构体用于优先队列中的元素
type Item struct {
	vertex   string  // 当前顶点的名称
	distance float64 // 从起点到该顶点的当前最短距离
	index    int     // 该元素在优先队列中的索引（用于更新队列时维护堆结构）
}

// 优先队列的实现，基于 container/heap
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// 按距离从小到大排序，实现小顶堆
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1] //返回要删除的元素
	old[n-1] = nil   // 避免内存泄漏
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra 算法实现
func Dijkstra(graph map[string]map[string]float64, start string) (map[string]float64, map[string]string) {
	//记录从起点到每个顶点的最短距离
	shortestPath := make(map[string]float64)
	//记录每个顶点的前驱节点，用于回溯最短路径
	predecessors := make(map[string]string)
	for vertex := range graph {
		shortestPath[vertex] = math.Inf(1) // 初始距离为无穷大
		predecessors[vertex] = ""          // 前驱节点初始化为空
	}
	shortestPath[start] = 0 // 起点到自身距离为0

	// 初始化优先队列
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{vertex: start, distance: 0})

	for pq.Len() > 0 {
		// 取出当前距离最小的顶点
		item := heap.Pop(&pq).(*Item)
		currentVertex := item.vertex
		currentDistance := item.distance

		// 如果当前距离大于已知最短路径，则说明已经处理处理过了，跳过
		if currentDistance > shortestPath[currentVertex] {
			continue
		}

		// 遍历所有邻居
		for neighbor, weight := range graph[currentVertex] {
			distance := currentDistance + weight
			// 如果找到更短的路径，则更新
			if distance < shortestPath[neighbor] {
				shortestPath[neighbor] = distance
				predecessors[neighbor] = currentVertex
				heap.Push(&pq, &Item{vertex: neighbor, distance: distance})
			}
		}
	}
	return shortestPath, predecessors
}
