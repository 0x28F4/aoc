package point

import (
	"fmt"

	"github.com/0x28F4/aoc2024/utils"
)

type Point struct {
	X int
	Y int
}

func FromSlice(s []int) Point {
	utils.MustLen(s, 2)

	return Point{s[0], s[1]}
}

func FromStringSlice(s []string) Point {
	utils.MustLen(s, 2)
	return Point{utils.MustInt(s[0]), utils.MustInt(s[1])}
}

func (p Point) String() string {
	return fmt.Sprintf("X=%d|Y=%d", p.X, p.Y)
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Neg() Point {
	return Point{
		X: -p.X,
		Y: -p.Y,
	}
}

func (p Point) Sub(other Point) Point {
	return p.Add(other.Neg())
}

func (p Point) Mul(other Point) Point {
	return Point{
		p.X * other.X,
		p.Y * other.Y,
	}
}

func (p Point) Div(other Point) Point {
	return Point{
		p.X / other.X,
		p.Y / other.Y,
	}
}

func (p Point) Mod(mod Point) Point {
	x := p.X % mod.X
	y := p.Y % mod.Y
	if x < 0 {
		x += mod.X
	}
	if y < 0 {
		y += mod.Y
	}
	return Point{x, y}
}

func (p Point) MulScal(s int) Point {
	return Point{
		p.X * s,
		p.Y * s,
	}
}

type DirFn func(s Point) Point

var (
	UP        DirFn = func(s Point) Point { return Point{X: s.X, Y: s.Y - 1} }
	DOWN      DirFn = func(s Point) Point { return Point{X: s.X, Y: s.Y + 1} }
	LEFT      DirFn = func(s Point) Point { return Point{X: s.X - 1, Y: s.Y} }
	RIGHT     DirFn = func(s Point) Point { return Point{X: s.X + 1, Y: s.Y} }
	DOWNRIGHT DirFn = func(p Point) Point { return Point{X: p.X + 1, Y: p.Y + 1} }
	UPRIGHT   DirFn = func(p Point) Point { return Point{X: p.X + 1, Y: p.Y - 1} }
	DOWNLEFT  DirFn = func(p Point) Point { return Point{X: p.X - 1, Y: p.Y + 1} }
	UPLEFT    DirFn = func(p Point) Point { return Point{X: p.X - 1, Y: p.Y - 1} }
)
