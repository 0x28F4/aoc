package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/0x28F4/aoc2024/utils"
)

type mult struct {
	enabled bool
	a       int
	b       int
}

type parser struct {
	state string
}

func (p *parser) shrink(i int) {
	if len(p.state) < i {
		p.state = ""
		return
	}

	p.state = p.state[i:]
}

func (p *parser) takeLit(s string) bool {
	l := len(s)
	if len(p.state) < l {
		return false
	}
	if p.state[0:l] == s {
		p.shrink(l)
		return true
	}
	return false
}

func (p *parser) takeInt() (int, bool) {
	ret := ""
	for {
		r := p.state[0:1]
		if _, err := strconv.Atoi(r); err == nil {
			p.shrink(1)
			ret += r
		} else {
			break
		}
	}
	if ret == "" {
		return 0, false
	}
	return utils.MustInt(ret), true
}

func parse(raw string) []mult {
	p := &parser{raw}

	insts := make([]mult, 0)
	enabled := true
	for {
		if len(p.state) <= 0 {
			break
		}

		if p.takeLit("do()") {
			enabled = true
		}

		if p.takeLit("don't()") {
			enabled = false
		}

		if !p.takeLit("mul") {
			p.shrink(1)
			continue
		}
		if !p.takeLit("(") {
			p.shrink(1)
			continue
		}

		m := mult{enabled: enabled}
		if num, foundInt := p.takeInt(); !foundInt {
			p.shrink(1)
			continue
		} else {
			m.a = num
		}

		if !p.takeLit(",") {
			p.shrink(1)
			continue
		}
		if num, foundInt := p.takeInt(); !foundInt {
			p.shrink(1)
			continue
		} else {
			m.b = num
		}
		if !p.takeLit(")") {
			p.shrink(1)
			continue
		}

		insts = append(insts, m)
	}

	return insts
}

var isPartTwo = flag.Bool("b", false, "select if part two")
var inputFile = flag.String("input", "day3/input", "select input file")

func main() {
	flag.Parse()
	solve()
}

func handleInput() string {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)
	bytes, err := io.ReadAll(file)
	utils.HandleError(err)
	return string(bytes)
}

func solve() {
	mults := parse(handleInput())
	score := 0
	for _, mult := range mults {
		add := mult.a * mult.b
		if *isPartTwo && !mult.enabled {
			add = 0
		}
		score += add
	}
	fmt.Println(score)
}
