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

var seen = set.New[point.Point]()
var gardens = make([]*garden, 0)
var cross = map[string]point.DirFn{
	"UP":    point.UP,
	"DOWN":  point.DOWN,
	"LEFT":  point.LEFT,
	"RIGHT": point.RIGHT,
}

func main() {
	flag.Parse()
	handleInput()
	solve()
}

var con c.Container

func solve() {
	for _, p := range con.Points() {
		if seen.Contains(p) {
			continue
		}

		v, err := con.At(p)
		utils.MustNil(err)

		g := newGarden(v, p)
		g.grow()
		gardens = append(gardens, g)
	}

	price := 0
	for _, g := range gardens {
		if g.kind == "#" {
			continue
		}
		area := len(g.plots)
		perimeter := g.countNbs()
		pr := perimeter * area
		// fmt.Printf("A region of %s plants with price %d * %d = %d.\n", g.kind, area, perimeter, pr)
		price += pr
	}

	fmt.Println("part 1", price)

	price = 0
	for _, g := range gardens {
		if g.kind == "#" {
			continue
		}

		area := len(g.plots)
		sides := g.countSides()
		pr := area * sides

		// fmt.Printf("A region of %s plants with price area=%d * sides=%d = %d.\n", g.kind, area, sides, pr)
		price += pr
	}
	fmt.Println("part 2", price)
}

type edge struct {
	point point.Point
	dir   string
}

type garden struct {
	start point.Point
	plots set.Set[point.Point]
	edges set.Set[edge]
	kind  string
}

func newGarden(kind string, first point.Point) *garden {
	return &garden{
		start: first,
		plots: set.New[point.Point](),
		edges: set.New[edge](),
		kind:  kind,
	}
}

func (g *garden) grow() {
	startV, err := con.At(g.start)
	utils.MustNil(err)

	var _grow func(point.Point)
	_grow = func(pos point.Point) {
		if seen.Contains(pos) {
			return
		}
		if g.plots.Contains(pos) { // does this happen?
			return
		}
		nxtV, err := con.At(pos)
		if err != nil {
			return
		}
		if startV != nxtV {
			return
		}

		g.plots.Add(pos)
		seen.Add(pos)
		for dir, dirFn := range cross {
			nxt := dirFn(pos)
			_grow(nxt)
			if g.isOnEdge(nxt) {
				g.edges.Add(edge{point: pos, dir: dir})
			}
		}
	}
	_grow(g.start)
}

func (g *garden) countNbs() int {
	var nbs []point.Point
	for _, p := range g.plots.Items() {
		for _, dirFn := range cross {
			nxt := dirFn(p)
			v, err := con.At(nxt)
			if err != nil {
				continue
			}
			if v != g.kind {
				nbs = append(nbs, nxt)
			}
		}
	}
	return len(nbs)
}

func (g *garden) isOnEdge(p point.Point) bool {
	v, err := con.At(p)
	if err != nil {
		return false
	}
	if v == g.kind {
		// point is inside the garden
		return false
	}

	for _, dirFn := range cross {
		nxt := dirFn(p)
		v, err := con.At(nxt)
		if err != nil {
			continue
		}
		if v == g.kind {
			return true
		}
	}

	return false
}

var tangents = map[string][]string{
	"RIGHT": {"DOWN", "UP"},
	"LEFT":  {"UP", "DOWN"},
	"DOWN":  {"LEFT", "RIGHT"},
	"UP":    {"RIGHT", "LEFT"},
}

func (g *garden) countSides() int {
	// only count edges which are non neighbors
	keep := set.New[edge]()
	for {
		cur := g.edges.First()
		keep.Add(cur)

		// shoot tangentially until no more edges to be removed
		for _, dir := range tangents[cur.dir] {
			pos := cur.point
			for {
				nxt := edge{cross[dir](pos), cur.dir}
				if g.edges.Contains(nxt) && !keep.Contains(nxt) {
					g.edges.Rem(nxt)
					pos = nxt.point
				} else {
					break
				}
			}
		}
		g.edges.Rem(cur)

		if g.edges.Len() == 0 {
			break
		}
	}

	return keep.Len()
}

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	con = c.NewPadded(lines, "#")
}
