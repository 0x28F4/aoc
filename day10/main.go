package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	c "github.com/0x28F4/aoc2024/utils/container/generic"
	"github.com/0x28F4/aoc2024/utils/point"
	"github.com/0x28F4/aoc2024/utils/set"
)

var inputFile = flag.String("input", "example", "select input file")

func main() {
	flag.Parse()
	score := 0
	input := handleInput("")
	for _, s := range solve(input) {
		score += s
	}

	fmt.Println("part1", score)
	fmt.Println("part2", solveB(input))

}

func solve(m Map) []int {
	startingPoints := m.FindAll(value{height: 0})
	var scores []int
	for _, start := range startingPoints {
		score := len(find9s(m, start))
		scores = append(scores, score)
	}
	return scores
}

func solveB(m Map) (score int) {
	startingPoints := m.FindAll(value{height: 0})
	for _, start := range startingPoints {
		score += findDistinctPaths(m, start)
	}
	return
}

func find9s(m Map, start point.Point) set.Set[point.Point] {
	var traverse func(pos point.Point) set.Set[point.Point]
	traverse = func(pos point.Point) set.Set[point.Point] {
		targets := set.New[point.Point]()
		v, err := m.At(pos)
		if err != nil {
			log.Panicf("[traverse] error when access %s | err = %s", pos, err)
		}
		if v.height == 9 {
			targets.Add(pos)
			return targets
		}

		for _, dir := range []dirfn{up, down, left, right} {
			nb := dir(pos)
			if nbv, err := m.At(nb); err == nil {
				if nbv.height-1 == v.height {
					targets.Add(traverse(nb).Items()...)
				}
			}
		}
		return targets
	}

	return traverse(start)
}

func findDistinctPaths(m Map, start point.Point) int {
	var traverse func(pos point.Point) int
	traverse = func(pos point.Point) int {
		v, err := m.At(pos)
		if err != nil {
			log.Panicf("[traverse] error when access %s | err = %s", pos, err)
		}
		if v.height == 9 {
			return 1
		}

		var paths int
		for _, dir := range []dirfn{up, down, left, right} {
			nb := dir(pos)
			if nbv, err := m.At(nb); err == nil {
				if nbv.height-1 == v.height {
					paths += traverse(nb)
				}
			}
		}
		return paths
	}

	return traverse(start)
}

type value struct {
	height      int
	notRelevant bool
}

func (v value) String() string {
	if v.notRelevant {
		return "."
	}
	return fmt.Sprintf("%d", v.height)
}

type Map struct {
	c.Container[value]
}

type dirfn func(s point.Point) point.Point

var (
	up    dirfn = func(s point.Point) point.Point { return point.Point{X: s.X, Y: s.Y - 1} }
	down  dirfn = func(s point.Point) point.Point { return point.Point{X: s.X, Y: s.Y + 1} }
	left  dirfn = func(s point.Point) point.Point { return point.Point{X: s.X - 1, Y: s.Y} }
	right dirfn = func(s point.Point) point.Point { return point.Point{X: s.X + 1, Y: s.Y} }
)

func handleInput(raw string) Map {
	if raw == "" {
		file, err := os.Open(*inputFile)
		utils.HandleError(err)

		bytes, err := io.ReadAll(file)
		utils.HandleError(err)
		raw = string(bytes)
	}

	var values [][]value
	for _, line := range strings.Split(raw, "\n") {
		var vals []value
		for _, r := range line {
			if s := string(r); s == "." {
				vals = append(vals, value{notRelevant: true})
			} else {
				vals = append(vals, value{height: utils.MustInt(string(r))})
			}
		}
		values = append(values, vals)
	}
	return Map{c.New(values)}
}
