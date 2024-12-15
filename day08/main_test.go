package main

import (
	"iter"
	"testing"

	"github.com/0x28F4/aoc2024/utils"
	"github.com/0x28F4/aoc2024/utils/point"
)

func TestAntinodes(t *testing.T) {
	pair := pair{a: point.Point{X: 1, Y: 1}, b: point.Point{X: 2, Y: 2}}

	_next, stop := iter.Pull(pair.shoot(true))
	next := func() point.Point {
		v, _ := _next()
		return v
	}

	utils.MustEq(next(), point.Point{X: 0, Y: 0})
	utils.MustEq(next(), point.Point{X: -1, Y: -1})
	utils.MustEq(next(), point.Point{X: -2, Y: -2})
	utils.MustEq(next(), point.Point{X: -3, Y: -3})
	stop()
}
