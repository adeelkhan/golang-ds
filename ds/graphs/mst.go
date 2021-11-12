package main

import (
	"fmt"
	"math"
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

func find_min_cost_vertex(g *graph, src uint64, bucket []uint64, visit []uint64) (uint64, uint64, bool) {
	var cost uint64 = math.MaxUint64
	var i, v uint64
	var found = false

	for i = 0; i < g.size; i++ {
		if visit[i] != 1 && g.IsConnected(src, i) == true {
			if g.g[src][i] < cost {
				v, cost, found = i, g.g[src][i], true
			}
		}
	}
	return v, cost, found
}
func compute_mst(g *graph, bucket []uint64, visit []uint64, path map[Path]uint64) {
	var vert, src uint64

	for uint64(len(bucket)) < g.size {
		var min_cost uint64 = math.MaxUint64
		for _, val := range bucket {
			v, cost, found := find_min_cost_vertex(g, val, bucket, visit)
			if found == true && min_cost > cost {
				min_cost, vert, src = cost, v, val
			}
		}
		bucket = append(bucket, vert)
		visit[vert] = 1

		path[Path{from: src, to: vert}] = min_cost
		vert = src
	}
}

func mst(g *graph, path map[Path]uint64) {
	bucket := make([]uint64, 1)
	visit := make([]uint64, g.size)
	bucket[0], visit[0] = 0, 1
	compute_mst(g, bucket, visit, path)
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

	mst(&g, path)
	fmt.Println(path)
}
