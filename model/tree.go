package model

type Edge struct {
	To   int
	Cost int
}

type Node struct {
	Id        int
	Depth     int
	Children  []int
	Edge_cost []Edge
}

type TreeTarget struct {
	Tree   map[int]*Node `json:"tree"`
	Target int           `json:"target"`
	Algo   string        `json:"algo_name"`
}
