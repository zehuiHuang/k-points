package law

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	graph := map[string]map[string]float64{
		"A": {"B": 4, "C": 5},
		"B": {"A": 4, "C": 11, "D": 9, "E": 7},
		"C": {"A": 5, "B": 11, "E": 3},
		"D": {"B": 9, "E": 13, "F": 2},
		"E": {"C": 3, "B": 7, "D": 13, "F": 6},
		"F": {"D": 2, "E": 6},
	}
	shortestPath, predecessors := Dijkstra(graph, "A")
	fmt.Println("最短路径：", shortestPath)
	fmt.Println("前驱节点：", predecessors)
}
