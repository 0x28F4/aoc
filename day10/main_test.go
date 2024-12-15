package main

import (
	"slices"
	"testing"

	"github.com/0x28F4/aoc2024/utils"
)

func TestTrails(t *testing.T) {
	maps := []string{
		`0123
1234
8765
9876`,
		`...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`, `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`, `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`, `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
	}
	trailScores := [][]int{{1}, {2}, {4}, {1, 2}, {5, 6, 5, 3, 1, 3, 5, 3, 5}}
	trailSums := []int{1, 2, 4, 3, 36}

	utils.MustEq(len(trailScores), len(maps))
	utils.MustEq(len(trailScores), len(trailSums))
	for i := 0; i < len(maps); i++ {
		m := handleInput(maps[i])
		scores := trailScores[i]
		actualScores := solve(m)

		utils.MustEq(len(scores), len(actualScores))
		sum := 0
		for _, score := range actualScores {
			utils.MustTrue(slices.Contains(actualScores, score))
			sum += score
		}
		utils.MustEq(trailSums[i], sum)
	}
}
