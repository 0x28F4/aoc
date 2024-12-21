package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
	"github.com/0x28F4/aoc2024/utils/point"
)

var inputFile = flag.String("input", "input", "select input file")

func main() {
	flag.Parse()
	handleInput()

	solution := 0
	alt := 0
	for _, ma := range machines {
		n, m := ma.solve()
		cost := m + 3*n
		if cost != ma.solveWithMath() {
			fmt.Println(cost, ma.solveWithMath())
		}

		alt += ma.solveWithMath()
		solution += cost
	}

	fmt.Println("part 1", solution, alt)
}

type machine struct {
	aButton point.Point
	bButton point.Point
	price   point.Point
}

func (ma machine) solve() (int, int) {
	divA := ma.price.Div(ma.aButton)
	divB := ma.price.Div(ma.bButton)

	for n := range max(divA.X, divA.Y) + 1 {
		for m := range max(divB.X, divB.Y) + 1 {
			if ma.price == ma.aButton.MulScal(n).Add(ma.bButton.MulScal(m)) {
				return n, m
			}
		}
	}
	return 0, 0
}

func (ma machine) solveWithMath() int {
	// num :=

	B := ma.aButton
	A := ma.bButton
	P := ma.price
	den := A.Y*B.X - A.X*B.Y
	if den == 0 {
		panic("div by zero")
	}

	num := P.Y*(B.X-3*A.X) - P.X*(B.Y-3*A.Y)

	if num%den == 0 {
		c := num / den

		mNum := (P.X - A.X*c)
		mDen := (B.X - 3*A.X)
		// divide 0 by 0 🤷
		if mNum == 0 && mDen == 0 {
			return c
		}
		if mDen == 0 {
			return 0
		}

		m := (P.X - A.X*c) / (B.X - 3*A.X)
		n := c - 3*m

		if m < 0 {
			return 0
		}
		if n < 0 {
			return 0
		}

		return c
	}

	return 0
}

var machines []machine

var buttonRe = regexp.MustCompile(`^Button .: X\+(\d+), Y\+(\d+)$`)
var priceRe = regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	parts := strings.Split(input, "\n\n")

	for _, part := range parts {
		lines := strings.Split(part, "\n")
		utils.MustLen(lines, 3)
		aMatch := buttonRe.FindStringSubmatch(lines[0])
		utils.MustLen(aMatch, 3)

		bMatch := buttonRe.FindStringSubmatch(lines[1])
		utils.MustLen(bMatch, 3)

		pMatch := priceRe.FindStringSubmatch(lines[2])
		utils.MustLen(pMatch, 3)

		machines = append(machines, machine{
			aButton: point.FromStringSlice(aMatch[1:]),
			bButton: point.FromStringSlice(bMatch[1:]),
			price:   point.FromStringSlice(pMatch[1:]),
		})
	}
}

// Ax * n + Bx * m = Px
// Ay * n + By * m = Py
// n + 3m = cost

// Ax * (cost - 3m) + Bx * m = Px
// Ay * (cost - 3m) + By * m = Py

// Ax * cost - 3Ax m + Bx m = Px
// Ay * cost - 3Ay m + By m = Py

// Px - Ax * cost = (Bx - 3Ax) m
// Py - Ay * cost = (By - 3Ay) m

// (Px - Ax * cost) / (Bx - 3Ax) = m
// (Py - Ay * cost) / (By - 3Ay) = m

// (Px - Ax * cost) / (Bx - 3Ax) = (Py - Ay * cost) / (By - 3Ay)
//
