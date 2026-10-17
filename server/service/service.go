package service

import (
	"petproj/model"
	"time"
)

type Result struct {
	Path    []int
	Versh   []int
	Perehod int
	PC      int
	Time    time.Duration
}

func DFS(tree map[int]*model.Node, target int) ([]int, []int, int, time.Duration) {
	start := time.Now()
	put_po_vershinam := []int{}
	stack := []int{1}
	parent := make(map[int]int)
	path := []int{target}
	counter := 0
	for len(stack) != 0 {
		curent := stack[len(stack)-1]
		put_po_vershinam = append(put_po_vershinam, curent)
		counter += 1
		stack = stack[:len(stack)-1]
		if curent != target {
			_, ok := tree[curent]
			if ok {
				for i := 0; i < len(tree[curent].Children); i++ {
					parent[tree[curent].Children[i]] = curent
					stack = append(stack, tree[curent].Children[i])
				}
			}
		} else {
			for curent != 1 {
				path = append(path, parent[curent])
				curent = parent[curent]
			}
			left := 0
			right := len(path) - 1
			for left < right {
				buf := path[left]
				path[left] = path[right]
				path[right] = buf
				left++
				right--
			}
			elapsed := time.Since(start)
			return path, put_po_vershinam, counter, elapsed
		}
	}
	elapsed := time.Since(start)
	return []int{}, put_po_vershinam, counter, elapsed
}

func BFS(tree map[int]*model.Node, target int) ([]int, []int, int, time.Duration) {
	start := time.Now()
	put_po_vershinam := []int{}
	queue := []int{1}
	parent := make(map[int]int)
	path := []int{target}
	counter := 0
	for len(queue) != 0 {
		curent := queue[0]
		queue = queue[1:]
		counter++
		put_po_vershinam = append(put_po_vershinam, curent)
		if curent != target {
			_, ok := tree[curent]
			if ok {
				for i := 0; i < len(tree[curent].Children); i++ {
					parent[tree[curent].Children[i]] = curent
					queue = append(queue, tree[curent].Children[i])
				}
			}
		} else {
			for curent != 1 {
				path = append(path, parent[curent])
				curent = parent[curent]
			}
			left := 0
			right := len(path) - 1
			for left < right {
				buf := path[left]
				path[left] = path[right]
				path[right] = buf
				left++
				right--
			}
			elapsed := time.Since(start)
			return path, put_po_vershinam, counter, elapsed
		}
	}
	elapsed := time.Since(start)
	return []int{}, put_po_vershinam, counter, elapsed
}

func PathCost(tree map[int]*model.Node, path []int) int {
	var pathcost int

	for i := 0; i < len(path)-1; i++ {
		for _, value := range tree[path[i]].Edge_cost {
			if value.To == path[i+1] {
				pathcost += value.Cost
				break
			}
		}
	}
	return pathcost
}

func Helper(req model.TreeTarget) map[string]Result {
	results := make(map[string]Result)
	if req.Algo == "BFS" || req.Algo == "BOTH" {
		Bfs_Path, Bfs_Versh, Bfs_Perehod, Bfs_vremya := BFS(req.Tree, req.Target)
		PC := PathCost(req.Tree, Bfs_Path)

		results["BFS"] = Result{
			Path:    Bfs_Path,
			Versh:   Bfs_Versh,
			Perehod: Bfs_Perehod,
			PC:      PC,
			Time:    Bfs_vremya,
		}
	}

	if req.Algo == "DFS" || req.Algo == "BOTH" {
		Dfs_Path, Dfs_Versh, Dfs_Perehod, Dfs_vremya := DFS(req.Tree, req.Target)
		PC := PathCost(req.Tree, Dfs_Path)

		results["DFS"] = Result{
			Path:    Dfs_Path,
			Versh:   Dfs_Versh,
			Perehod: Dfs_Perehod,
			PC:      PC,
			Time:    Dfs_vremya,
		}
	}
	return results

}
