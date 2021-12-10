package main

import (
	"bufio"
	"fmt"
	"os"
)

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

  total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
      if score, ok := calcScore(line); !ok {
        total += score
			}
		}
	}

  fmt.Println("Total score:", total)
}

func calcScore(line string) (int, bool) {
	open := []rune{}
  score := 0

	for _, c := range line {
		if _, ok := pairs[c]; ok {
			open = append(open, c)
			continue
		}

		if c != pairs[open[len(open)-1]] {
      switch c {
        case ')':
          score = 3
          break
        case ']':
          score = 57
          break
        case '}':
          score = 1197
          break
        case '>':
          score = 25137
          break
      }
			return score, false
		}

    open = open[:len(open)-1]
	}

	return score, true
}
