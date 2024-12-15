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
	"github.com/0x28F4/aoc2024/utils/set"
)

var inputFile = flag.String("input", "example", "select input file")

func main() {
	flag.Parse()
	solve()
}

func handleInput() []string {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	return strings.Split(input, "\n")
}

type directionFn func(point.Point) point.Point

type direction string

var (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

var dirToFn = map[direction]directionFn{
	up:    func(s point.Point) point.Point { return point.Point{X: s.X, Y: s.Y - 1} },
	down:  func(s point.Point) point.Point { return point.Point{X: s.X, Y: s.Y + 1} },
	left:  func(s point.Point) point.Point { return point.Point{X: s.X - 1, Y: s.Y} },
	right: func(s point.Point) point.Point { return point.Point{X: s.X + 1, Y: s.Y} },
}

func transition(d direction) direction {
	if d == up {
		return right
	}
	if d == right {
		return down
	}
	if d == down {
		return left
	}
	if d == left {
		return up
	}
	panic(fmt.Sprintf("unknown dirction %s", d))
}

type pointWithDir struct {
	point.Point
	dir direction
}

type guard struct {
	pos        point.Point
	dir        direction
	path       set.Set[point.Point]
	beenBefore set.Set[pointWithDir]
}

func newGuard(pos point.Point, dir direction) guard {
	return guard{
		pos:        pos,
		dir:        dir,
		path:       set.New[point.Point](),
		beenBefore: set.New[pointWithDir](),
	}
}

func (g guard) traverse(con c.Container) (isLoop bool) {
	for {
		g.path.Add(g.pos)
		fn := dirToFn[g.dir]
		nxtPos := fn(g.pos)
		v, err := con.At(nxtPos)
		if err != nil {
			return false
		}
		if v == "#" {
			g.dir = transition(g.dir)
			g.beenBefore.Add(pointWithDir{g.pos, g.dir})
			continue
		}
		g.pos = nxtPos

		pd := pointWithDir{g.pos, g.dir}
		if g.beenBefore.Contains(pd) {
			return true
		}
		g.beenBefore.Add(pointWithDir{g.pos, g.dir})
	}
}

func mutateContainer(con c.Container, pos point.Point) (c.Container, error) {
	nc := con.Copy()
	if err := nc.Set(pos, "#"); err != nil {
		return c.Container{}, err
	}
	return nc, nil
}

func solve() {
	con := c.New(handleInput())
	p, err := con.FindFirst("^")
	utils.MustNil(err)
	g := newGuard(p, up)
	utils.MustFalse(g.traverse(con))
	fmt.Println("part1", g.path.Len())

	withLoops := set.New[point.Point]()
	for _, obstacle := range g.path.Items() {
		if obstacle == p {
			continue
		}

		newCon, err := mutateContainer(con, obstacle)
		if err != nil {
			fmt.Println("not able to mutate container with obstacle at", obstacle)
			continue
		}
		ng := newGuard(p, up)
		if ng.traverse(newCon) {
			withLoops.Add(obstacle)
		}
	}
	fmt.Println("part2", len(withLoops))
}
