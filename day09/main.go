package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

var inputFile = flag.String("input", "example4", "select input file")

func main() {
	flag.Parse()
	solve()
}

type memory struct {
	data   []*int
	cursor int
}

func newMemory() *memory {
	return &memory{
		data:   make([]*int, 0),
		cursor: 0,
	}
}

func (m memory) String() string {
	b := strings.Builder{}
	for _, d := range m.data {
		if d == nil {
			b.WriteString(".")
		} else {
			b.WriteString(fmt.Sprintf("%d", *d))
		}
	}
	return b.String()
}

func (m *memory) Seek(i int) {
	m.cursor = i
}

func (m *memory) write(v *int) {
	if m.cursor == len(m.data) {
		m.data = append(m.data, v)
	} else {
		m.data[m.cursor] = v
	}
	m.cursor = m.cursor + 1
}

func (m *memory) findGap(minSize int, until int) int {
	window := 0
	for i := 0; i <= until; i++ {
		if m.data[i] == nil {
			window++
			if window >= minSize {
				return i - window + 1
			}
		} else {
			window = 0
		}
	}

	return -1
}

func (m *memory) score() int {
	score := 0
	for pos, b := range m.data {
		if b == nil {
			continue
		}
		score += pos * *b
	}
	return score
}

func handleInput() *memory {
	mem := newMemory()
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)

	isBlock := true
	id := 0
	for _, r := range input {
		i := utils.MustInt(string(r))
		if isBlock {
			for range i {
				v := id
				mem.write(&v)
			}
		} else {
			for range i {
				mem.write(nil)
			}
			id++
		}
		isBlock = !isBlock
	}

	return mem
}

func solve() {
	mem := handleInput()
	for i := len(mem.data) - 1; i >= 0; i-- {
		d := mem.data[i]
		if d == nil {
			continue
		}

		for j := 0; j < i; j++ {
			if mem.data[j] != nil {
				continue
			}
			mem.Seek(j)
			mem.write(d)
			mem.Seek(i)
			mem.write(nil)
			break
		}
	}

	fmt.Println("part 1", mem.score())

	mem = handleInput()
	cursor := len(mem.data) - 1
	for {
		if cursor <= 0 {
			break
		}
		if value := mem.data[cursor]; value != nil {
			fileCursor := cursor
			for {
				if fileCursor == 0 {
					break
				}
				if mem.data[fileCursor] == nil {
					fileCursor++
					break
				}
				if *value != *mem.data[fileCursor] {
					fileCursor++
					break
				}
				fileCursor--
			}
			fileSize := cursor - fileCursor + 1
			if gap := mem.findGap(fileSize, fileCursor); gap != -1 {
				mem.Seek(gap)
				for range fileSize {
					mem.write(value)
				}

				mem.Seek(fileCursor)
				for range fileSize {
					mem.write(nil)
				}

			}
			cursor = fileCursor - 1
		} else {
			cursor--
		}
	}

	fmt.Println("part 2", mem.score())
}
