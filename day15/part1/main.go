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

func (b *box) push(dir string, isRobot bool) point.Point {
	if len(dir) != 1 {
		panic("length of dir is not equal to 1")
	}
	dirfn, exists := cross[dir]
	utils.MustTrue(exists)
	b.tx = tx(func(b *box) { b.pos = dirfn(b.pos) })
	nxt := dirfn(b.pos)
	return nxt
}

func (b *box) commit() {
	b.tx(b)
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
	for y := range s.dimension.X {
		lines[y] = strings.Repeat(".", s.dimension.X)
	}
	con := container.New(lines)

	for _, b := range s.boxes {
		con.Set(b.pos, "O")
	}

	for _, w := range s.walls.Items() {
		con.Set(w, "#")
	}

	con.Set(s.robot.pos, "@")
	con.Print()
}

func (s *simulation) score() (score int) {
	for _, box := range s.boxes {
		score += box.score()
	}
	return
}

func (s *simulation) update() (done bool) {
	var dir string
	dir, s.inst = utils.StringPopLeft(s.inst)
	if dir == "" {
		return true
	}

	rollback := false
	cur := s.robot.push(dir, true)
	for {
		if s.walls.Contains(cur) {
			rollback = true
			break
		}

		foundBox := false
		for _, box := range s.boxes {
			if box.pos == cur {
				cur = box.push(dir, false)
				foundBox = true

			}
		}

		if foundBox {
			continue
		}

		break
	}

	if !rollback {
		utils.MustNotNil(s.robot.tx)
		s.robot.commit()

		for _, box := range s.boxes {
			if box.tx != nil {
				box.commit()
			}
		}
	}

	s.robot.clear()
	for _, box := range s.boxes {
		box.clear()
	}

	return
}

func solve() {
	sim := parseMap()
	for !sim.update() {
	}
	fmt.Println("part 1", sim.score())
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
	ret.dimension = point.Point{X: len(lines[0]), Y: len(lines)}
	for y, line := range lines {
		for x, r := range line {
			p := point.Point{X: x, Y: y}
			v := string(r)
			if v == "@" {
				ret.robot = &box{pos: p}
				continue
			}

			if v == "#" {
				ret.walls.Add(p)
				continue
			}

			if v == "O" {
				ret.boxes = append(ret.boxes, &box{pos: p})
				continue
			}

			utils.MustEq(v, ".")
		}
	}

	return
}
