package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	container "github.com/0x28F4/aoc2024/utils/container/string"
	"github.com/0x28F4/aoc2024/utils/point"
	"github.com/0x28F4/aoc2024/utils/set"
)

var inputFile = flag.String("input", "example2", "select input file")

var cross = map[string]point.DirFn{
	"^": point.UP,
	"v": point.DOWN,
	"<": point.LEFT,
	">": point.RIGHT,
}

func main() {
	flag.Parse()
	solve()
}

type tx func(b *box)

type box struct {
	pos point.Point
	tx  tx
}

func (b *box) push(dir string) (point.Point, point.Point) {
	if len(dir) != 1 {
		panic("length of dir is not equal to 1")
	}
	dirfn, exists := cross[dir]
	utils.MustTrue(exists)
	b.tx = tx(func(b *box) { b.pos = dirfn(b.pos) })
	nxt := dirfn(b.pos)
	return nxt, nxt.Add(R)
}

func (b *box) collide(other point.Point) bool {
	return b.pos == other || b.pos.Add(R) == other
}

func (b *box) commit() {
	if b.tx != nil {
		b.tx(b)
	}
}

func (b *box) clear() {
	b.tx = nil
}

func (b *box) score() int {
	return 100*b.pos.Y + b.pos.X
}

type simulation struct {
	walls set.Set[point.Point]
	robot *box
	boxes []*box
	inst  string

	dimension point.Point
}

func (s *simulation) print() {
	lines := make([]string, s.dimension.Y)
	for y := range s.dimension.Y {
		lines[y] = strings.Repeat(".", s.dimension.X)
	}
	con := container.New(lines)

	for _, b := range s.boxes {
		con.Set(b.pos, "[")
		con.Set(b.pos.Add(R), "]")
	}

	for _, w := range s.walls.Items() {
		con.Set(w, "#")
	}

	con.Set(s.robot.pos, "@")
	con.Print()
}

func (s *simulation) score() (score int) {
	for _, box := range s.boxes {
		sc := box.score()
		score += sc
	}
	return
}

var R = point.Point{X: 1, Y: 0}

type force struct {
	at  point.Point
	dir string
}

func (f force) apply(s *simulation) bool {
	var _apply func(p point.Point) bool
	_apply = func(p point.Point) (wallCollide bool) {
		if s.walls.Contains(p) {
			return true
		}
		for _, b := range s.boxes {
			if b.collide(p) && b.tx == nil {
				lhs, rhs := b.push(f.dir)
				// found box, should push the left and right side of it
				return _apply(lhs) || _apply(rhs)
			}
		}
		// default case, found .
		return false
	}

	return _apply(f.at)
}

func (s *simulation) update() (done bool) {
	var dir string
	dir, s.inst = utils.StringPopLeft(s.inst)
	if dir == "" {
		return true
	}

	cur, _ := s.robot.push(dir)
	f := force{
		at:  cur,
		dir: dir,
	}
	if !f.apply(s) {
		s.robot.commit()
		for _, b := range s.boxes {
			b.commit()
		}

	}
	s.robot.clear()
	for _, b := range s.boxes {
		b.clear()
	}
	return
}

func solve() {
	sim := parseMap()
	for !sim.update() {
	}
	sim.print()
	fmt.Println("part 2", sim.score())
}

func handleInput() (string, string) {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)

	parts := strings.Split(input, "\n\n")
	utils.MustLen(parts, 2)

	return parts[0], parts[1]
}

func parseMap() (ret *simulation) {
	rawMap, instr := handleInput()
	ret = &simulation{
		walls: set.New[point.Point](),
	}
	ret.inst = strings.ReplaceAll(instr, "\n", "")

	lines := strings.Split(rawMap, "\n")
	ret.dimension = point.Point{X: len(lines[0]) * 2, Y: len(lines)}
	for y, line := range lines {
		for x, r := range line {
			v := string(r)
			a := point.Point{X: x * 2, Y: y}
			b := point.Point{X: x*2 + 1, Y: y}

			if v == "@" {
				ret.robot = &box{pos: a}
				continue
			}

			if v == "#" {
				ret.walls.Add(a)
				ret.walls.Add(b)
				continue
			}

			if v == "O" {
				ret.boxes = append(ret.boxes, &box{pos: a})
				continue
			}

			utils.MustEq(v, ".")
		}
	}

	return
}
