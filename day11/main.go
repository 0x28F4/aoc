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
var steps = flag.Int("steps", 75, "select number of steps")

func main() {
	flag.Parse()
	solve()
}

func solve() {
	a := handleInput()

	a.solve()
	fmt.Println(a.count())

}

type stone struct {
	int
	steps int
}

func (s stone) String() string {
	return fmt.Sprintf("%d", s.int)
}

func (s stone) step() []stone {
	if s.int == 0 {
		return []stone{{1, s.steps - 1}}
	}
	digits := s.String()
	if len(digits)%2 == 0 {
		lhs := utils.MustInt(digits[0 : len(digits)/2])
		rhs := utils.MustInt(digits[len(digits)/2:])

		return []stone{{lhs, s.steps - 1}, {rhs, s.steps - 1}}
	}

	return []stone{{s.int * 2024, s.steps - 1}}
}

type arrangement struct {
	stoneMap map[stone]int
}

func (a *arrangement) count() (c int) {
	for s := range a.stoneMap {
		c += a.stoneMap[s]
	}
	return
}

func (a *arrangement) addStone(s stone, times int) {
	if _, ok := a.stoneMap[s]; !ok {
		a.stoneMap[s] = 0
	}
	a.stoneMap[s] = a.stoneMap[s] + times
}

func (a *arrangement) solve() {
	for {
		allDone := true
		for s := range a.stoneMap {
			if s.steps > 0 {
				allDone = false
				break
			}
		}
		if allDone {
			break
		}

		for s := range a.stoneMap {
			if s.steps == 0 {
				continue
			}

			for _, newStone := range s.step() {
				a.addStone(newStone, a.stoneMap[s])
			}

			delete(a.stoneMap, s)
		}
	}
}

func handleInput() *arrangement {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	a := &arrangement{stoneMap: make(map[stone]int)}
	for _, r := range strings.Fields(string(bytes)) {
		s := stone{int: utils.MustInt(r), steps: *steps}
		a.addStone(s, 1)
	}

	return a
}

// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
