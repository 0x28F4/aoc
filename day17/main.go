package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

var inputFile = flag.String("input", "input", "select input file")

func main() {
	flag.Parse()
	solve()
}

func solve() {
	a, progSeq := handleInput()
	fmt.Println("part 1", program(a))
	targets := [][]int{
		{3, 0},
		{5, 5, 3, 0},
		{1, 6, 5, 5, 3, 0},
		{4, 1, 1, 6, 5, 5, 3, 0},
		{0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
		{7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
		{1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
		{2, 4, 1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
	}

	solutions := []int{}
	for k, target := range targets {
		// first run
		if k == 0 {
			for i := range utils.Pow(2, 6) {
				output := program(i)
				if utils.IsSliceEq(target, output) {
					solutions = append(solutions, i)
				}
			}
		} else {
			var solutionsNxt []int
			for _, sol := range solutions {
				for i := range utils.Pow(2, 6) {
					output := program(sol + i)
					if utils.IsSliceEq(target, output) {
						solutionsNxt = append(solutionsNxt, sol+i)
					}
				}
			}
			solutions = solutionsNxt
		}
		fmt.Printf("finished step %d with %d solutions -- got %d\n", k, len(solutions), target)
		if k == 7 {
			break
		}
		for i, sol := range solutions {
			solutions[i] = sol << 6
		}
	}

	solution := utils.Min(solutions)

	// validate
	utils.MustSliceEq(progSeq, program(solution))
	fmt.Println("part 2", solution)
}

func handleInput() (a int, prog []int) {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)

	parts := strings.Split(input, "\n\n")

	utils.MustLen(parts, 2)

	registerLines := strings.Split(parts[0], "\n")
	utils.MustLen(registerLines, 3)
	a = utils.MustInt(strings.Split(registerLines[0], ": ")[1])

	pParts := strings.Split(parts[1], ": ")
	utils.MustLen(pParts, 2)

	for _, inst := range strings.Split(pParts[1], ",") {
		prog = append(prog, utils.MustInt(inst))
	}

	return
}

// looking hard enough at the input yields this program
func program(input int) (output []int) {
	A := input
	B := 0
	C := 0

	for A > 0 {
		// [2,4]
		B = A % 8
		// [1,5]
		B = B ^ 5 // 0x00000101
		// [7,5]
		C = A / (utils.Pow(2, B))
		// [0,3]
		A = A / (utils.Pow(2, 3))
		// [4,1]
		B = B ^ C
		// [1,6]
		B = B ^ 6 // 0x00001010
		// [5,5]
		output = append(output, B%8)
	}

	return output
}
