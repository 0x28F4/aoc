package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

func solve() {
	file, err := os.Open("./input_d1")
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		utils.MustLen(items, 2)
		l1 = append(l1, utils.MustInt(items[0]))
		l2 = append(l2, utils.MustInt(items[1]))
	}

	sort.Ints(l1)
	sort.Ints(l2)
	totalDist := 0
	for i, left := range l1 {
		right := l2[i]
		totalDist += utils.Abs(right - left)
	}
	fmt.Println("result", totalDist)
}

func solveB() {
	file, err := os.Open("./input_d1")
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		utils.MustLen(items, 2)
		l1 = append(l1, utils.MustInt(items[0]))
		l2 = append(l2, utils.MustInt(items[1]))
	}

	freq := make(map[int]int)
	for _, v := range l2 {
		before, ok := freq[v]
		if !ok {
			freq[v] = 0
		}

		freq[v] = before + 1
	}

	score := 0
	for _, v := range l1 {
		f := freq[v]
		score += v * f
	}
	fmt.Println(score)
}

var isPartTwo = flag.Bool("b", false, "select if part two")

func main() {
	flag.Parse()
	if *isPartTwo {
		solveB()
	} else {
		solve()
	}
}
