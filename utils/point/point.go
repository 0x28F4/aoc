package point

import "fmt"

type Point struct {
	X int
	Y int
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

func (p Point) MulScal(s int) Point {
	return Point{
		p.X * s,
		p.Y * s,
	}
}
