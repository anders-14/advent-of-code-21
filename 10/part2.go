package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	scores := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if score, ok := calcScore(line); ok {
				scores = append(scores, score)
			}
		}
	}

	sort.Ints(scores)
	fmt.Println("Middle score:", scores[(len(scores)-1)/2])
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
			return score, false
		}

		open = open[:len(open)-1]
	}

	if len(open) == 0 {
		return score, false
	}

	for i := range open {
    closing := pairs[open[len(open)-1-i]]
		score *= 5
		switch closing {
		case ')':
			score += 1
			break
		case ']':
			score += 2
			break
		case '}':
			score += 3
			break
		case '>':
			score += 4
			break
		}
	}

	return score, true
}
