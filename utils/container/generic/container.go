package container

import (
	"errors"
	"fmt"
	"slices"

	"github.com/0x28F4/aoc2024/utils/point"
)

type Container[T comparable] struct {
	Rows [][]T
}

func New[T comparable](rows [][]T) Container[T] {
	return Container[T]{rows}
}

func NewPadded[T comparable](rows [][]T, padd T) Container[T] {
	tb := slices.Repeat([]T{padd}, len(rows[0])+2)
	ret := [][]T{tb}
	for _, row := range rows {
		row = slices.Insert(row, 0, padd)
		row = append(row, padd)
		ret = append(ret, row)
	}
	tb = slices.Repeat([]T{padd}, len(rows[0])+2)
	ret = append(ret, tb)
	return Container[T]{ret}
}

func (c Container[T]) Print() {
	for _, line := range c.Rows {
		fmt.Println(line)
	}
}

var ErrOutOfBounds = errors.New("out of bounds")

func (c Container[T]) At(p point.Point) (T, error) {
	x := p.X
	y := p.Y
	if x < 0 {
		return *new(T), ErrOutOfBounds
	}
	if y < 0 {
		return *new(T), ErrOutOfBounds
	}
	if x >= len(c.Rows[0]) {
		return *new(T), ErrOutOfBounds
	}
	if y >= len(c.Rows) {
		return *new(T), ErrOutOfBounds
	}
	return c.Rows[y][x], nil
}

func (c Container[T]) Set(p point.Point, v T) error {
	if _, err := c.At(p); err != nil {
		return err
	}

	c.Rows[p.Y] = slices.Insert(c.Rows[p.Y], p.X, v)
	return nil
}

func (c Container[T]) FindFirst(s T) (point.Point, error) {
	for y := 0; y < len(c.Rows); y++ {
		for x := 0; x < len(c.Rows[0]); x++ {
			if c.Rows[y][x] == s {
				return point.Point{X: x, Y: y}, nil
			}
		}
	}
	return point.Point{}, fmt.Errorf("not found")
}

func (c Container[T]) FindAll(s T) []point.Point {
	var ret []point.Point
	for y := 0; y < len(c.Rows); y++ {
		for x := 0; x < len(c.Rows[0]); x++ {
			if c.Rows[y][x] == s {
				ret = append(ret, point.Point{X: x, Y: y})
			}
		}
	}
	return ret
}

func (c Container[T]) Copy() Container[T] {
	ret := make([][]T, 0)
	for _, row := range c.Rows {
		newRow := make([]T, len(row))
		copy(newRow, row)
		ret = append(ret, newRow)
	}
	return Container[T]{
		Rows: ret,
	}
}
