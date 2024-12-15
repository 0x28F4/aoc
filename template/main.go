package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

var isPartTwo = flag.Bool("b", false, "select if part two")
var inputFile = flag.String("input", "template/input", "select input file")

func main() {
	flag.Parse()
	solve()
}

func solve() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")
	fmt.Println(lines[0])

	if *isPartTwo {
		panic("not implemented")
	}

	panic("not implemented")
}
