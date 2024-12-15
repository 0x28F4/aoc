package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	c "github.com/0x28F4/aoc2024/utils/container/string"
	"github.com/0x28F4/aoc2024/utils/point"
)

var isPartTwo = flag.Bool("b", true, "select if part two")
var inputFile = flag.String("input", "example4", "select input file")

func main() {
	flag.Parse()
	if !*isPartTwo {
		solve()
		return
	}
	solveB()
}

func handleInput() []string {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	return strings.Split(input, "\n")
}

func solve() {
	con := c.NewPadded(handleInput())

	points := []point.Point{{X: 0, Y: 0}}
	cur := point.Point{X: 0, Y: 0}
	s := 0
	for {
		scan := boundScanner[s]
		prev := cur
		cur = scan(cur)
		if (cur.X == 0) && (cur.Y == 0) {
			break
		}
		if v, err := con.At(cur); err == nil {
			points = append(points, cur)
			utils.MustEq(v, "#")
		} else {
			cur = prev
			s += 1
		}

	}
	utils.MustLen(points, len(con.Lines[0])*2+len(con.Lines)*2-4)

	scannedLines := make([]string, 0)
	for _, scan := range scanners {
		for _, p := range points {
			var err error
			var v string
			cur := p
			var s string
			for err == nil {
				v, err = con.At(cur)
				s += v
				cur = scan(cur)
			}
			scannedLines = append(scannedLines, s)
		}
	}

	score := 0
	for _, line := range scannedLines {
		lScore := strings.Count(line, "XMAS")
		score += lScore
		if lScore > 0 {
			fmt.Println("----")
			fmt.Println(line)
			fmt.Println("count ", lScore)
		}
	}
	fmt.Println(score)
}

type scanner func(p point.Point) point.Point

var scanners = []scanner{
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X + 1, Y: p.Y} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X - 1, Y: p.Y} }),

	scanner(func(p point.Point) point.Point { return point.Point{X: p.X, Y: p.Y + 1} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X, Y: p.Y - 1} }),

	scanner(func(p point.Point) point.Point { return point.Point{X: p.X + 1, Y: p.Y + 1} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X + 1, Y: p.Y - 1} }),

	scanner(func(p point.Point) point.Point { return point.Point{X: p.X - 1, Y: p.Y + 1} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X - 1, Y: p.Y - 1} }),
}

// down right up left
var boundScanner = []scanner{
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X, Y: p.Y + 1} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X + 1, Y: p.Y} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X, Y: p.Y - 1} }),
	scanner(func(p point.Point) point.Point { return point.Point{X: p.X - 1, Y: p.Y} }),
}

func solveB() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")
	con := c.NewPadded(lines)
	con.Print()

	score := 0
	for y := 0; y < len(con.Lines); y++ {
		for x := 0; x < len(con.Lines[0]); x++ {
			for _, sten := range stencils {
				sum := sten.conv(con, x, y)
				if sum == 5 {
					fmt.Println("found match at: ", x, y)
					score += 1
				}

			}
		}
	}
	fmt.Println(score)
}

type stencil []string

func (s stencil) conv(con c.Container, x, y int) (sum int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := 0
			sP := s[j][i : i+1]
			if cP, err := con.At(point.Point{X: x + i, Y: y + j}); err == nil {
				if sP == cP {
					v = 1
				}
			}
			sum += v
		}
	}
	return
}

var stencils = []stencil{
	{
		"M*S",
		"*A*",
		"M*S",
	},
	{
		"S*M",
		"*A*",
		"S*M",
	},
	{
		"M*M",
		"*A*",
		"S*S",
	},
	{
		"S*S",
		"*A*",
		"M*M",
	},
}
