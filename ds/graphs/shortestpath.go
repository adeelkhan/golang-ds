package main

import (
	"fmt"
)

type graph struct {
	g    [][]uint64
	size uint64
}

func (gr *graph) SetConnection(u, v, cost uint64) {
	if u < gr.size && v < gr.size {
		gr.g[u][v] = cost
	}
}

func (gr *graph) IsConnected(u, v uint64) bool {
	if u < gr.size && v < gr.size {
		if gr.g[u][v] > 0 {
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
	gr.g = make([][]uint64, gr.size)
	var i uint64
	for i = 0; i < gr.size; i++ {
		gr.g[i] = make([]uint64, gr.size)
	}
}

type Path struct {
	from uint64
	to   uint64
}

func shortest_path(g *graph, path map[Path]uint64) {

}

var path map[Path]uint64

func main() {
	g := graph{}
	path = make(map[Path]uint64)
	g.Init(5)

	g.SetConnection(0, 1, 2)
	g.SetConnection(0, 2, 1)
	g.SetConnection(1, 4, 3)
	g.SetConnection(1, 3, 1)
	g.SetConnection(1, 2, 2)
	g.SetConnection(2, 3, 5)
	g.SetConnection(2, 4, 4)
	g.print()
}
