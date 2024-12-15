package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

var isPartTwo = flag.Bool("b", false, "select if part two")
var inputFile = flag.String("input", "example", "select input file")

func main() {
	flag.Parse()
	solve()
}

type rule struct {
	lhs int
	rhs int
}

var rules []rule

func (r rule) apply(line []int) error {
	for i := 0; i < len(line); i++ {
		before := line[0:i]
		// after := line[i+1:]
		v := line[i]

		if r.lhs == v && slices.Contains(before, r.rhs) {
			return fmt.Errorf("violates rule %d|%d", r.lhs, r.rhs)
		}
	}

	return nil
}

var sequences [][]int

func handleInput() {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	var rawRules []string
	var rawSequences []string
	for i, l := range lines {
		if l == "" {
			rawRules = lines[0:i]
			rawSequences = lines[i+1:]
			break
		}
	}

	for _, raw := range rawRules {
		split := strings.Split(raw, "|")
		utils.MustLen(split, 2)

		lhs := utils.MustInt(split[0])
		rhs := utils.MustInt(split[1])

		rules = append(rules, rule{lhs, rhs})
	}

	for _, seq := range rawSequences {
		split := strings.Split(seq, ",")

		ret := make([]int, 0)
		for _, s := range split {
			ret = append(ret, utils.MustInt(s))
		}
		sequences = append(sequences, ret)
	}
}

func solve() {
	handleInput()

	if !*isPartTwo {
		score := 0
		for _, seq := range sequences {
			valid := true
			for _, rule := range rules {
				if err := rule.apply(seq); err != nil {
					fmt.Println(seq, err)
					valid = false
					break
				}

			}
			if valid {
				mid := seq[len(seq)/2]
				score += mid
			}
		}
		fmt.Println(score)
	}

	score := 0
	for _, seq := range sequences {
		invalid := false
		for _, rule := range rules {
			if err := rule.apply(seq); err != nil {
				invalid = true
				break
			}
		}

		if invalid {
			isFixed := false
			fixed := seq
			for !isFixed {
				isFixed = true
				for _, r := range rules {
					if err := r.apply(fixed); err != nil {
						isFixed = false
						i := slices.IndexFunc(fixed, func(s int) bool { return s == r.lhs })
						j := slices.IndexFunc(fixed, func(s int) bool { return s == r.rhs })
						fixed[i], fixed[j] = fixed[j], fixed[i]
						break
					}
				}
			}
			mid := seq[len(seq)/2]
			score += mid
		}
	}
	fmt.Println(score)
}
