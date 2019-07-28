package graph

import (
	"container/list"
	"fmt"
)

//adjacency table, 无向图
type Graph struct {
	adj []*list.List
	v   int
}

// NewGraph init graphh according to capacity
func NewGraph(v int) *Graph {
	graph := &Graph{}
	graph.v = v
	graph.adj = make([]*list.List, v)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

// AddEdge insert as add edge，一条边存2次
func (g *Graph) AddEdge(s int, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// BFS search path by BFS
func (g *Graph) BFS(s int, t int) {

	//todo
	if s == t {
		return
	}

	//init prev
	prev := make([]int, g.v)
	for index := range prev {
		prev[index] = -1
	}

	//search by queue
	var queue []int
	visited := make([]bool, g.v)
	queue = append(queue, s)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedlist := g.adj[top]
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}

}

// DFS search by DFS
func (g *Graph) DFS(s int, t int) {

	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}

	visited := make([]bool, g.v)
	visited[s] = true

	isFound := false
	g.recurse(s, t, prev, visited, isFound)

	printPrev(prev, s, t)
}

// recurse recursivly find path
func (g *Graph) recurse(s int, t int, prev []int, visited []bool, isFound bool) {

	if isFound {
		return
	}

	visited[s] = true

	if s == t {
		isFound = true
		return
	}

	linkedlist := g.adj[s]
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k, t, prev, visited, false)
		}
	}

}

// printPrev print path recursively
func printPrev(prev []int, s int, t int) {

	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}
