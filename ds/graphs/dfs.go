package main

import "fmt"

type graph struct {
	g    [][]int
	size uint64
}

func (gr *graph) SetConnection(u, v uint64) {
	if u < gr.size && v < gr.size {
		gr.g[u][v] = 1
	}
}

func (gr *graph) IsConnected(u, v uint64) bool {
	if u < gr.size && v < gr.size {
		if gr.g[u][v] == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}
func (gr *graph) print() {
	var i, j uint64
	for i = 0; i < gr.size; i++ {
		for j = 0; j < gr.size; j++ {
			fmt.Printf("%d", gr.g[i][j])
		}
		fmt.Println()
	}
}

func (gr *graph) Init(size uint64) {
	gr.size = size
	gr.g = make([][]int, gr.size)
	var i uint64
	for i = 0; i < gr.size; i++ {
		gr.g[i] = make([]int, gr.size)
	}
}

func dfs_walk(gr *graph, src uint64, visit []uint64) {
	fmt.Println(src)
	var i uint64
	for i = 0; i < gr.size; i++ {
		if visit[i] != 1 && gr.IsConnected(src, i) == true {
			visit[i] = 1
			dfs_walk(gr, i, visit)
		}
	}
}

func dfs(gr *graph) {
	visit := make([]uint64, gr.size)
	var src uint64 = 1
	visit[src] = 1
	dfs_walk(gr, src, visit)
}

func main() {
	g := graph{}
	g.Init(6)

	g.SetConnection(1, 2)
	g.SetConnection(1, 3)
	g.SetConnection(2, 3)
	g.SetConnection(2, 4)
	g.SetConnection(3, 5)
	g.SetConnection(4, 5)
	g.print()
	dfs(&g)

}
