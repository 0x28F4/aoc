package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	"github.com/0x28F4/aoc2024/utils/point"
)

var inputFile = flag.String("input", "example", "select input file")

func main() {
	flag.Parse()
	handleInput()
	fmt.Println(mapSize)
	for _, coord := range coordinates {
		fmt.Println(coord)

	}
}

var mapSize int
var coordinates []point.Point

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	parts := strings.Split(input, "\n\n")

	mapSize = utils.MustInt(parts[0])
	for _, c := range strings.Split(parts[1], "\n") {
		cxy := strings.Split(c, ",")
		utils.MustLen(cxy, 2)
		coordinates = append(coordinates, point.Point{X: utils.MustInt(cxy[0]), Y: utils.MustInt(cxy[1])})
	}
}
