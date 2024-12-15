package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	"github.com/0x28F4/aoc2024/utils"
)

var inputFile = flag.String("input", "template/input", "select input file")

func main() {
	flag.Parse()
	solve()
}

func solve() {
	

}

var con 

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")
	fmt.Println(lines[0])
}