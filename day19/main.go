package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

var inputFile = flag.String("input", "example", "select input file")

func main() {
	flag.Parse()
	handleInput()
	solution := 0
	for _, des := range designs {
		if Parse(des) > 0 {
			solution++
		}
	}
	fmt.Println("part 1", solution)

	solution = 0
	for _, des := range designs {
		solution += Parse(des)
	}
	fmt.Println("part 2", solution)
}

var cache map[string]int = make(map[string]int)

func Parse(design string) int {
	cache[""] = 1

	var _parse func(rem string) int
	_parse = func(rem string) (ret int) {
		if num, hit := cache[rem]; hit {
			return num
		}

		for _, lit := range literals {
			if strings.HasPrefix(rem, lit) {
				ret += _parse(rem[len(lit):])
			}
		}

		cache[rem] = ret
		return ret
	}

	return _parse(design)
}

var literals []string
var designs []string

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	parts := strings.Split(input, "\n\n")
	utils.MustLen(parts, 2)

	for _, s := range strings.Split(parts[0], ", ") {
		literals = append(literals, s)
	}

	for _, des := range strings.Split(parts[1], "\n") {
		designs = append(designs, des)
	}
}
