package container

import (
	"errors"
	"fmt"
	"strings"

	"github.com/0x28F4/aoc2024/utils/point"
)

type Container struct {
	Lines []string
}

func New(lines []string) Container {
	return Container{lines}
}

func NewPadded(lines []string) Container {
	tb := strings.Repeat("#", len(lines[0])+2)

	ret := []string{tb}
	for _, line := range lines {
		ret = append(ret, fmt.Sprintf("#%s#", line))
	}

	ret = append(ret, tb)
	return Container{ret}
}

func (c Container) Print() {
	for _, line := range c.Lines {
		fmt.Println(line)
	}
}

var ErrOutOfBounds = errors.New("out of bounds")

func (c Container) At(p point.Point) (string, error) {
	x := p.X
	y := p.Y
	if x < 0 {
		return "", ErrOutOfBounds
	}
	if y < 0 {
		return "", ErrOutOfBounds
	}
	if x >= len(c.Lines[0]) {
		return "", ErrOutOfBounds
	}
	if y >= len(c.Lines) {
		return "", ErrOutOfBounds
	}
	return c.Lines[y][x : x+1], nil
}

func (c Container) Set(p point.Point, r string) error {
	if _, err := c.At(p); err != nil {
		return err
	}

	l := c.Lines[p.Y]
	c.Lines[p.Y] = l[:p.X] + r + l[p.X+1:]
	return nil
}

func (c Container) FindFirst(s string) (point.Point, error) {
	for y := 0; y < len(c.Lines); y++ {
		for x := 0; x < len(c.Lines[0]); x++ {
			if c.Lines[y][x:x+1] == s {
				return point.Point{X: x, Y: y}, nil
			}
		}
	}
	return point.Point{}, fmt.Errorf("not found")
}

func (c Container) Copy() Container {
	nl := make([]string, len(c.Lines))
	copy(nl, c.Lines)
	return Container{
		Lines: nl,
	}
}
