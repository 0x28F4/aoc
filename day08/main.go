package main

import (
	"flag"
	"fmt"
	"io"
	"iter"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	c "github.com/0x28F4/aoc2024/utils/container/string"
	"github.com/0x28F4/aoc2024/utils/point"
	"github.com/0x28F4/aoc2024/utils/set"
)

var inputFile = flag.String("input", "example", "select input file")

var con c.Container
var antennas = make(map[string][]point.Point)

func solve() {
	handleInput()
	antinodes := set.New[point.Point]()
	for _, points := range antennas {
		pairs := makePairs(points)
		for _, pair := range pairs {
			for _, forward := range []bool{true, false} {
				next, stop := iter.Pull(pair.shoot(forward))
				p, _ := next()
				if _, err := con.At(p); err == nil {
					antinodes.Add(p)
				}
				stop()
			}
		}
	}
	fmt.Println("part1", len(antinodes))

	antinodes = set.New[point.Point]()
	for _, points := range antennas {
		pairs := makePairs(points)
		utils.MustTrue(len(pairs) > 0)
		for _, pa := range pairs {
			antinodes.Add(pa.a, pa.b)
			for _, forward := range []bool{true, false} {
				next, stop := iter.Pull(pa.shoot(forward))
				for {
					p, _ := next()
					if _, err := con.At(p); err != nil {
						stop()
						break
					}
					antinodes.Add(p)
				}

			}
		}
	}
	fmt.Println("part2", len(antinodes))
}

func main() {
	flag.Parse()
	solve()
}

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)
	lines := strings.Split(string(bytes), "\n")

	con = c.Container{
		Lines: lines,
	}

	for y, line := range lines {
		for x, r := range line {
			s := string(r)
			if s == "." {
				continue
			}
			if _, exists := antennas[s]; !exists {
				antennas[s] = make([]point.Point, 0)
			}

			sl := antennas[s]
			sl = append(sl, point.Point{X: x, Y: y})
			antennas[s] = sl
		}
	}
}

type pair struct {
	a point.Point
	b point.Point
}

func makePairs(points []point.Point) []pair {
	ret := make([]pair, 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			ret = append(ret, pair{points[i], points[j]})
		}
	}
	return ret
}

func (p pair) shoot(forward bool) iter.Seq[point.Point] {
	return func(yield func(point.Point) bool) {
		cur := p.a
		vec := p.a.Sub(p.b)
		if !forward {
			cur = p.b
			vec = p.b.Sub(p.a)
		}
		for {
			cur = cur.Add(vec)
			if !yield(cur) {
				return
			}
		}
	}
}
