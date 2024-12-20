package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/0x28F4/aoc2024/utils"
)

var inputFile = flag.String("input", "input", "select input file")

func main() {
	flag.Parse()
	solve()
}

type op struct {
	opcode  int
	operand int
}

type instruction interface {
	Run(reg *program) (jumpTo *int)
	String() string
	GetOp() op
}

func newInstruction(opcode, operand int) instruction {
	o := op{
		opcode:  opcode,
		operand: operand,
	}

	if opcode == 0 {
		return adv{o}
	}

	if opcode == 1 {
		return bxl{o}
	}

	if opcode == 2 {
		return bst{o}
	}

	if opcode == 3 {
		return jnz{o}
	}

	if opcode == 4 {
		return bxc{o}
	}

	if opcode == 5 {
		return out{o}
	}

	if opcode == 6 {
		return bdv{o}
	}

	if opcode == 7 {
		return cdv{o}
	}

	panic(fmt.Sprintf("unreadhable code, if opcode is in bounds 0 to 7, got %d", opcode))
}

func (ins op) String() string {
	return fmt.Sprintf("[%d,%d]", ins.opcode, ins.operand)
}

func (o op) GetOp() op {
	return o
}

func (ins op) combo(reg [3]register) int {
	utils.MustSmaller(ins.opcode, 8)
	utils.MustGreaterEq(ins.opcode, 0)
	utils.MustNotEq(ins.operand, 7)
	if ins.operand <= 3 {
		return ins.operand
	}
	return reg[ins.operand-4].value
}

type adv struct {
	op
}

func (o adv) Run(p *program) (jumpTo *int) {
	// [0,3]
	// A = A / 2^3
	numerator := p.reg[0].value
	denominator := utils.Pow(2, o.combo(p.reg))
	p.reg[0].value = numerator / denominator
	return
}

type bxl struct {
	op
}

func (o bxl) Run(p *program) (jumpTo *int) {
	// [1,5] takes B and bitwise XOR with b0101
	// B = B ^ b0101

	// [1,6] takes B and bitwise XOR with b0110
	// B = B ^ b0110

	p.reg[1].value = p.reg[1].value ^ o.operand
	return
}

type bst struct {
	op
}

func (o bst) Run(p *program) (jumpTo *int) {
	// [2,4] takes the a register and writes mod8 of it to B register
	// B = mod8(A)
	p.reg[1].value = o.combo(p.reg) % 8
	return
}

type jnz struct {
	op
}

func (o jnz) Run(p *program) (jumpTo *int) {
	// [3,0] jump to start if A != 0
	if p.reg[0].value == 0 {
		return
	}
	jumpTo = &o.operand
	return
}

type bxc struct {
	op
}

func (o bxc) Run(p *program) (jumpTo *int) {
	// [4,1]
	// B = B ^ C

	// For legacy reasons, this instruction reads an operand but ignores it. ???
	p.reg[1].value = p.reg[1].value ^ p.reg[2].value
	return
}

type out struct {
	op
}

func (o out) Run(p *program) (jumpTo *int) {
	// [5,5]
	// fmt.Println(B)
	p.output = append(p.output, o.combo(p.reg)%8)
	return
}

type bdv struct {
	op
}

func (o bdv) Run(p *program) (jumpTo *int) {
	numerator := p.reg[0].value
	denominator := utils.Pow(2, o.combo(p.reg))
	p.reg[1].value = numerator / denominator
	return
}

type cdv struct {
	op
}

func (o cdv) Run(p *program) (jumpTo *int) {
	// [7,5] takes A register and divides it by 2^B:
	// C = A / 2^B
	numerator := p.reg[0].value
	denominator := utils.Pow(2, o.combo(p.reg))
	p.reg[2].value = numerator / denominator
	return
}

type register struct {
	value int
}

func (r register) String() string {
	return fmt.Sprintf("%d", r.value)
}

type program struct {
	// the instruction pointer
	ip int

	instructions []instruction
	reg          [3]register
	output       []int
}

func (p *program) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Register A: %s\n", p.reg[0]))
	sb.WriteString(fmt.Sprintf("Register B: %s\n", p.reg[1]))
	sb.WriteString(fmt.Sprintf("Register C: %s\n", p.reg[2]))

	var insRaw []string
	for _, ins := range p.instructions {
		insRaw = append(insRaw, ins.String())
	}

	sb.WriteString(strings.Repeat(" ", 10+6*p.ip))
	sb.WriteString("v\n")
	sb.WriteString("Program: ")
	sb.WriteString(strings.Join(insRaw, " "))
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat(" ", 10+6*p.ip))
	sb.WriteString("^\n")
	sb.WriteString(fmt.Sprintf("Output: %s", getResult(p.output)))

	return sb.String()
}

func (p *program) getProgram() (ret []int) {
	for _, ins := range p.instructions {
		ret = append(ret, ins.GetOp().opcode, ins.GetOp().operand)
	}

	return ret
}

func getResult(output []int) string {
	var st []string
	for _, o := range output {
		st = append(st, fmt.Sprintf("%d", o))
	}
	return strings.Join(st, ",")
}

func (p *program) run() {
	// run until halt
	for p.ip < len(p.instructions) {
		jumpTo := p.instructions[p.ip].Run(p)
		// fmt.Println(p)
		if jumpTo != nil {
			p.ip = *jumpTo
		} else {
			p.ip++
		}
	}
}

func newAttempt(parts [6]uint8, k int, change uint8) int {
	ret := 0
	for shift, part := range parts {
		v := part
		if k == shift {
			v = change
		}
		ret += int(v) << (shift * 8)
	}
	return ret
}

func solve() {
	// p := handleInput()
	// p.reg[0].value = 0b1100011_00100111_00110011_10111000_10001111_00000001
	// p.reg[0].value = 0b1100011_00100111_00110011_10101001_10001111_00000001
	// p.reg[0].value = 0b1100001_00000000_00000000_10101001_10001111_00000001
	// p.run()
	// fmt.Println(getResult(p.output))
	// fmt.Println(getResult(p.getProgram()))
	// fmt.Println("dist", utils.DistanceBinary(p.output, p.getProgram()))

	// areg := 109020022607617 //0b1100011_00100111_00110011_10111000_10001111_00000001
	// for i := range 48 {
	// 	(1 << 48)

	// for change := range 2 ^ 3 {
	// 	attempt := replacePart(areg, change, i)
	// 	if dist := utils.DistanceBinary(myProgram(attempt), p.getProgram()); dist == 0 {
	// 		if attempt < areg {
	// 			fmt.Printf("found better: %d\n", attempt)
	// 			areg = attempt
	// 		}
	// 	}
	// }

	// }

	// fmt.Printf("%b\n", 109020022607617)
	// fmt.Printf("%b\n", 109020021624577)
	// target := []int{5, 3, 0}

	// fmt.Println(myProgram(0b1100011_00100111_00110011_10111000_10001111_00000001))
	// fmt.Println(myProgram(0b11000110))

	targets := [][]int{
		{3, 0},
		{5, 5, 3, 0},
		{1, 6, 5, 5, 3, 0},
		{4, 1, 1, 6, 5, 5, 3, 0},
		{0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
		{7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
		{1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
		{2, 4, 1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
	}

	solutions := []int{}
	for k, target := range targets {
		// first run
		if k == 0 {
			for i := range utils.Pow(2, 6) {
				output := myProgram(i)
				if utils.IsSliceEq(target, output) {
					solutions = append(solutions, i)
				}
			}
		} else {
			var solutionsNxt []int
			for _, sol := range solutions {
				for i := range utils.Pow(2, 6) {
					output := myProgram(sol + i)
					if utils.IsSliceEq(target, output) {
						solutionsNxt = append(solutionsNxt, sol+i)
					}
				}
			}
			solutions = solutionsNxt
		}

		fmt.Printf("finished step %d with %d solutions -- got %d\n", k, len(solutions), target)
		// fmt.Println("solutions\n---")
		// for _, sol := range solutions {
		// 	fmt.Printf("%b\n", sol)
		// }

		// at the end shift solution 6 bits to the left to make space for the next batch

		if k == 7 {
			break
		}
		for i, sol := range solutions {
			solutions[i] = sol << 6
		}
	}

	solution := utils.Min(solutions)

	// validate
	utils.MustSliceEq([]int{2, 4, 1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0}, myProgram(solution))
	fmt.Println("part 2", solution)
}

func handleInput() *program {
	file, err := os.Open(*inputFile)
	utils.HandleError(err)

	bytes, err := io.ReadAll(file)
	utils.HandleError(err)

	input := string(bytes)

	parts := strings.Split(input, "\n\n")

	utils.MustLen(parts, 2)

	registerLines := strings.Split(parts[0], "\n")
	utils.MustLen(registerLines, 3)

	registers := [3]register{}
	for i, line := range registerLines {
		rParts := strings.Split(line, ": ")
		utils.MustLen(rParts, 2)
		num := utils.MustInt(rParts[1])
		registers[i] = register{
			num,
		}
	}

	pParts := strings.Split(parts[1], ": ")
	utils.MustLen(pParts, 2)

	instRaw := strings.Split(pParts[1], ",")

	var instructions []instruction
	for i := 0; i < len(instRaw); i += 2 {
		inst := newInstruction(utils.MustInt(instRaw[i]), utils.MustInt(instRaw[i+1]))
		instructions = append(instructions, inst)
	}

	return &program{
		instructions: instructions,
		reg:          registers,
	}
}

func myProgram(input int) (output []int) {
	A := input
	B := 0
	C := 0

	for A > 0 {
		// [2,4]
		B = A % 8
		// [1,5]
		B = B ^ 5 // 0x00000101
		// [7,5]
		C = A / (utils.Pow(2, B))
		// [0,3]
		A = A / (utils.Pow(2, 3))
		// [4,1]
		B = B ^ C
		// [1,6]
		B = B ^ 6 // 0x00001010
		// [5,5]
		output = append(output, B%8)
	}

	return output
}

// B = O ^ 6

// B = (O ^ 6) ^ C

// B^5 = (O ^ 6) ^ C

// (A%8)^5 = (O ^ 6) ^ C

// (A%8)^5 = (O ^ 6) ^ (A / (A%8)^5)

// (A / (A%8)^5) ^ (A%8)^5 = O^6

// ((K1 _ K2 _ K3 _ K4 _ K5 _ K6) / Ki) ^ Ki = Oi^6

// ((K1 _ K2 _ K3 _ K4 _ K5 _ K6) / K1) ^ K1 = Oi^6
