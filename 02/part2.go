package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	var ret []string

	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func main() {
	input := readInput()

	horizontal_pos := 0
	depth := 0
  aim := 0

	for _, line := range input {
		cmd := strings.Split(line, " ")
		dist, _ := strconv.Atoi(cmd[1])

		switch cmd[0] {
		case "forward":
			horizontal_pos += dist
      depth += aim * dist
			break
		case "up":
			aim -= dist
			break
		case "down":
			aim += dist
			break
		}
	}

	fmt.Printf("Horizontal position: %d\n", horizontal_pos)
	fmt.Printf("Depth:               %d\n", depth)
  fmt.Printf("Product:             %d\n", horizontal_pos * depth)
}
