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
	var input [][]string

	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			input = append(input, strings.Split(scanner.Text(), ""))
		}
	}

	sum := 0
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			if comparePoint(x, y, input) {
				fmt.Println(input[y][x], x, y)
				sum += num(input[y][x]) + 1
			}
		}
	}

	fmt.Printf("Sum of risk levels: %d\n", sum)
}

func num(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

func comparePoint(x, y int, input [][]string) bool {
	point := num(input[y][x])
	if y > 0 {
		if point >= num(input[y-1][x]) {
			return false
		}
	}
	if y < ymax-1 {
		if point >= num(input[y+1][x]) {
			return false
		}
	}
	if x > 0 {
		if point >= num(input[y][x-1]) {
			return false
		}
	}
	if x < xmax-1 {
		if point >= num(input[y][x+1]) {
			return false
		}
	}

	return true
}
