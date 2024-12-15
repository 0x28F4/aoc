package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

type level struct {
	numbers []int
	isAsc   bool
}

func newLevel(nums []int) level {
	isAsc := nums[0] < nums[1]
	return level{
		numbers: nums,
		isAsc:   isAsc,
	}
}

func (l level) isSave() bool {
	// check for dists
	for i := 1; i < len(l.numbers); i++ {
		current := l.numbers[i]
		previous := l.numbers[i-1]
		dist := utils.Abs(previous - current)
		if (dist != 1) && (dist != 2) && (dist != 3) {
			return false
		}
	}

	// check for asc / desc
	for i := 1; i < len(l.numbers); i++ {
		current := l.numbers[i]
		previous := l.numbers[i-1]
		if l.isAsc {
			if previous > current {
				return false
			}
		} else {
			if previous < current {
				return false
			}
		}
	}
	return true
}

func (l level) findAlts() []level {
	alts := make([]level, 0)
	nums := l.numbers
	for i := range l.numbers {
		altNums := make([]int, 0)
		for j, num := range nums {
			if i == j {
				continue
			}
			altNums = append(altNums, num)
		}
		alts = append(alts, newLevel(altNums))
	}
	return alts
}

func solve() {
	file, err := os.Open("./input_d2")
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	levels := make([]level, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		nums := make([]int, 0)
		for _, item := range items {
			nums = append(nums, utils.MustInt(item))
		}
		levels = append(levels, newLevel(nums))
	}

	score := 0
	for _, level := range levels {
		if !level.isSave() {
			continue
		}
		score += 1
	}

	fmt.Println(score)
}

func solveB() {
	file, err := os.Open("./input_d2")
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	levels := make([]level, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		nums := make([]int, 0)
		for _, item := range items {
			nums = append(nums, utils.MustInt(item))
		}
		levels = append(levels, newLevel(nums))
	}

	score := 0
	for _, level := range levels {
		if level.isSave() {
			score += 1
			continue
		}

		altSave := false
		for _, alt := range level.findAlts() {
			if alt.isSave() {
				altSave = true
				break
			}
		}

		if altSave {
			score += 1
		}
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
