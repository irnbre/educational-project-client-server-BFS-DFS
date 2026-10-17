package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Path    []int
	Versh   []int
	Perehod int
	PC      int
	Time    time.Duration
}

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

func main() {
	var target int
	var algo int

	fmt.Println("Введите целевую вершину")
	fmt.Scan(&target)
	fmt.Println(" ")
	fmt.Println("Выберите алгоритм: 1 - BFS\n2 - DFS\n3 - ОБА\n\n")
	fmt.Scan(&algo)
	var algo_name string
	switch algo {
	case 1:
		algo_name = "BFS"
	case 2:
		algo_name = "DFS"
	case 3:
		algo_name = "BOTH"
	}
	PATH := "client/data.txt"
	file, err := os.Open(PATH)
	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var col_vershin int
	_, err = fmt.Fscan(reader, &col_vershin)
	if err != nil {
		fmt.Println("Ошибка чтения")
		return
	}
	mp_derevo := make(map[int]*Node)
	for i := 1; i <= col_vershin; i++ {
		mp_derevo[i] = &Node{Id: i}
	}
	var versh, child, cost int

	for {
		count, err := fmt.Fscan(reader, &versh, &child, &cost)
		if err == io.EOF && count == 0 {
			break

		}
		if versh < 1 || versh > col_vershin || child < 1 || child > col_vershin || cost <= 0 {
			fmt.Println("Некорректные данные в входном фалйе")
			return
		}
		mp_derevo[versh].Children = append(mp_derevo[versh].Children, child)
		mp_derevo[versh].Edge_cost = append(mp_derevo[versh].Edge_cost, Edge{To: child, Cost: cost})
	}

	req := TreeTarget{Tree: mp_derevo, Target: target, Algo: algo_name}
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Ошибка кодирования дерева в json")
		return
	}
	resp, err := http.Post("http://127.0.0.1:8080/search", "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Println("Ошибка запроса к серверу")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	
		fmt.Printf("Ошибка сервера")
		return
	}

	var results map[string]Result

	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		fmt.Println("Ошибка декодирования ответа сервера")
		return
    }
	for algorithm, result := range results {
	fmt.Println("\nАлгоc:", algorithm)
	fmt.Println("Найденный путь:", result.Path)
	fmt.Println("Порядок обхода:", result.Versh)
	fmt.Println("Количество переходов:", result.Perehod)
	fmt.Println("Стоимость пути:", result.PC)
	fmt.Println("Время выполнения:", result.Time)
}

}
