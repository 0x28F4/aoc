package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unsafe"

	"github.com/0x28F4/aoc2024/utils"
	"github.com/0x28F4/aoc2024/utils/perm"
)

var inputFile = flag.String("input", "input", "select input file")

func main() {
	if s := unsafe.Sizeof(0); s != 8 {
		panic(fmt.Sprintf("only works on 64bit systems, sizeof int: %d", s))
	}
	flag.Parse()
	solve()
}

func solve() {
	calibrations := handleInput()
	score := 0
	for _, cal := range calibrations {
		if cal.isPossible([]string{"+", "*"}) {
			score += cal.testValue
		}
	}
	fmt.Println("part 1", score)

	score = 0
	for _, cal := range calibrations {
		if cal.isPossible([]string{"+", "*", "||"}) {
			score += cal.testValue
		}
	}
	fmt.Println("part 2", score)
}

type calibration struct {
	items     []int
	testValue int
}

func (c calibration) eval(ops []string) bool {
	utils.MustEq(len(ops), len(c.items)-1)
	res := c.items[0]
	for i := 1; i < len(c.items); i++ {
		if res > c.testValue {
			return false
		}
		op := ops[i-1]
		if op == "+" {
			res = res + c.items[i]
			continue
		}
		if op == "*" {
			res = res * c.items[i]
			continue
		}
		if op == "||" {
			res = utils.MustInt(fmt.Sprintf("%d%d", res, c.items[i]))
			continue
		}
		panic(fmt.Sprintf("unknown operator %s", op))
	}
	return res == c.testValue
}

func (c calibration) isPossible(ops []string) bool {
	for _, ops := range perm.Equal(len(c.items)-1, ops) {
		if c.eval(ops) {
			return true
		}
	}
	return false
}

func handleInput() []calibration {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	var ret []calibration
	for _, line := range lines {
		c := calibration{}
		parts := strings.Split(line, ":")
		utils.MustLen(parts, 2)
		c.testValue = utils.MustInt(parts[0])
		for _, v := range strings.Fields(parts[1]) {
			c.items = append(c.items, utils.MustInt(v))
		}
		ret = append(ret, c)
	}

	return ret
}
