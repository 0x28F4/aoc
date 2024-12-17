package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	container "github.com/0x28F4/aoc2024/utils/container/string"
	"github.com/0x28F4/aoc2024/utils/point"
)

var inputFile = flag.String("input", "example", "select input file")
var bathroom = point.Point{X: 101, Y: 103}

func main() {
	flag.Parse()

	if !strings.Contains(*inputFile, "input") {
		bathroom = point.Point{X: 11, Y: 7}
		fmt.Println("bathroom size changed!", bathroom)
	}

	solve()
}

type quad struct {
	origin    point.Point
	dimension point.Point
}

func (q quad) inside(p point.Point) bool {
	if p.X < q.origin.X {
		return false
	}
	if p.Y < q.origin.Y {
		return false
	}

	one := point.Point{X: 1, Y: 1}
	br := q.origin.Add(q.dimension).Sub(one)
	if p.X > br.X {
		return false
	}
	if p.Y > br.Y {
		return false
	}

	return true
}

func solve() {
	robots := handleInput()
	for range 100 {
		for _, r := range robots {
			r.update()
		}
	}

	half := bathroom.Div(point.Point{X: 2, Y: 2})
	quads := []quad{
		{
			origin:    point.Point{X: 0, Y: 0},
			dimension: half,
		},
		{
			origin:    point.Point{X: half.X + 1, Y: 0},
			dimension: half,
		},
		{
			origin:    point.Point{X: 0, Y: half.Y + 1},
			dimension: half,
		},
		{
			origin:    point.Point{X: half.X + 1, Y: half.Y + 1},
			dimension: half,
		},
	}

	var quadSums []int
	for _, q := range quads {
		i := 0
		for _, r := range robots {
			isInside := q.inside(r.pos)
			if isInside {
				i++
			}
		}
		quadSums = append(quadSums, i)
	}
	score := 1
	for _, s := range quadSums {
		score *= s
	}
	fmt.Println("part 1", score)

	buildMap := func() container.Container {
		lines := make([]string, bathroom.Y)
		for y := range bathroom.Y {
			lines[y] = strings.Repeat(" ", bathroom.X)
		}
		return container.New(lines)
	}

	robots = handleInput()
	rx := regexp.MustCompile(`########`)
	for t := range 100000 {
		m := buildMap()
		for _, r := range robots {
			r.update()
			m.Set(r.pos, "#")
		}
		if m.Re(rx) {
			m.Print()
			fmt.Println("part 2", t+1)
			break
		}
	}
}

type robot struct {
	pos point.Point
	vel point.Point
}

func (r *robot) String() string {
	return fmt.Sprintf("p=%d,%d v=%d,%d", r.pos.X, r.pos.Y, r.vel.X, r.vel.Y)
}

func (r *robot) update() {
	r.pos = r.pos.Add(r.vel).Mod(bathroom)
}

func handleInput() (ret []*robot) {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	parseVec := func(raw string) point.Point {
		parts := strings.Split(raw, "=")
		utils.MustLen(parts, 2)

		partsXY := strings.Split(parts[1], ",")
		utils.MustLen(partsXY, 2)
		return point.Point{X: utils.MustInt(partsXY[0]), Y: utils.MustInt(partsXY[1])}
	}
	for _, line := range lines {
		fields := strings.Fields(line)
		utils.MustLen(fields, 2)
		r := &robot{}
		r.pos = parseVec(fields[0])
		r.vel = parseVec(fields[1])
		ret = append(ret, r)
	}

	return
}
