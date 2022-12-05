package Day5all

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/darimm/AOCFunctions"
)

type Day5 struct {
	data []string
}

func (d *Day5) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day5\\a\\day5a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func Day5all() {
	var day Day5
	day.New()
	stack := day.GetStacks()
	day.ProcessStackPartA(stack)
	stack = day.GetStacks()
	fmt.Println()
	day.ProcessStackPartB(stack)
}

func (d *Day5) GetStacks() map[int][]string {
	stack := make(map[int][]string)
	for i := 0; i < len(d.data); i++ {
		_, err := strconv.Atoi(string(d.data[i][1]))
		if err == nil {
			break
		}
		for linepos := 1; linepos < len(d.data[i]); linepos += 4 {
			if string(d.data[i][linepos]) == " " {
				continue
			}
			stack[(linepos/4)+1] = append(stack[(linepos/4)+1], string(d.data[i][linepos]))
		}
	}
	//reverse the stacks
	for _, v := range stack {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
	}
	return stack
}

func (d *Day5) ProcessStackPartA(stack map[int][]string) {
	for i := 0; i < len(d.data); i++ {
		instructions := strings.Fields(d.data[i])
		if len(instructions) == 0 || instructions[0] != "move" {
			continue
		}
		toMove, err := strconv.Atoi(instructions[1])
		if err != nil {
			panic(err.Error())
		}
		fromstack, err := strconv.Atoi(instructions[3])
		if err != nil {
			panic(err.Error())
		}
		tostack, err := strconv.Atoi(instructions[5])
		if err != nil {
			panic(err.Error())
		}
		for mov := 0; mov < toMove; mov++ {
			stacklength := len(stack[fromstack]) - 1
			removed := stack[fromstack][stacklength]
			stack[fromstack] = stack[fromstack][:stacklength]
			stack[tostack] = append(stack[tostack], removed)
		}
	}
	for i := 1; i <= len(stack); i++ {
		fmt.Printf("%v", stack[i][len(stack[i])-1])
	}
}

func (d *Day5) ProcessStackPartB(stack map[int][]string) {
	for i := 0; i < len(d.data); i++ {
		instructions := strings.Fields(d.data[i])
		if len(instructions) == 0 || instructions[0] != "move" {
			continue
		}
		toMove, err := strconv.Atoi(instructions[1])
		if err != nil {
			panic(err.Error())
		}
		fromstack, err := strconv.Atoi(instructions[3])
		if err != nil {
			panic(err.Error())
		}
		tostack, err := strconv.Atoi(instructions[5])
		if err != nil {
			panic(err.Error())
		}
		stacklength := len(stack[fromstack]) - toMove
		removed := stack[fromstack][stacklength:]
		stack[fromstack] = stack[fromstack][:stacklength]
		stack[tostack] = append(stack[tostack], removed...)
	}
	for i := 1; i <= len(stack); i++ {
		fmt.Printf("%v", stack[i][len(stack[i])-1])
	}
}
