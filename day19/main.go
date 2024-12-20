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
	// fmt.Println(literals, designs)

	solution := 0
	for i, des := range designs {
		fmt.Printf("%d/%d\n", i, len(designs))
		// fmt.Println(des, Parse(des))
		if Parse(des) {
			solution++
		}
	}

	fmt.Println("part 1", solution)
}

func Parse(design string) bool {
	var _parse func(rem string) bool

	_parse = func(rem string) bool {
		if rem == "" {
			return true
		}

		for _, lit := range literals {
			if strings.HasPrefix(rem, lit) {
				if _parse(rem[len(lit):]) {
					return true
				}
			}
		}

		return false
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
