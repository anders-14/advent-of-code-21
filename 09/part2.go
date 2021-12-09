package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ymax = 100
const xmax = 100

func main() {
	var input [][]int

	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := strings.Split(scanner.Text(), "")
			var nums []int
			for _, v := range line {
				n, _ := strconv.Atoi(v)
				nums = append(nums, n)
			}
			input = append(input, nums)
		}
	}

  a, b, c := 0, 0, 0
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			if isLowpoint(x, y, input) {
				basin := [][2]int{}
				fillBasin(&basin, x, y, input)
        l := len(basin)
        if l > a {
          c = b
          b = a
          a = l
        } else if l > b {
          c = b
          b = l
        } else if l > c {
          c = l
        }
			}
		}
	}

  fmt.Println("Product of top 3 basins:", a * b * c)
}

func has(basin *[][2]int, x, y int) bool {
	for _, v := range *basin {
		if v[0] == y && v[1] == x {
			return true
		}
	}
	return false
}

func fillBasin(basin *[][2]int, x, y int, input [][]int) *[][2]int {
	if input[y][x] == 9 || has(basin, x, y) {
		return basin
	}

	*basin = append(*basin, [2]int{y, x})
	adjacent := findAdjacent(x, y, input)
	for _, adj := range adjacent {
		fillBasin(basin, adj[1], adj[0], input)
	}

	return basin
}

func isLowpoint(x, y int, input [][]int) bool {
	adjacent := findAdjacent(x, y, input)
	for _, adj := range adjacent {
		if input[y][x] > input[adj[0]][adj[1]] {
			return false
		}
	}
	return true
}

func findAdjacent(x, y int, input [][]int) [][2]int {
	adj := [][2]int{}

	if y > 0 {
		adj = append(adj, [2]int{y - 1, x})
	}
	if y < ymax-1 {
		adj = append(adj, [2]int{y + 1, x})
	}
	if x > 0 {
		adj = append(adj, [2]int{y, x - 1})
	}
	if x < xmax-1 {
		adj = append(adj, [2]int{y, x + 1})
	}

	return adj
}
