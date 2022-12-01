package Day1all

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/darimm/AOCFunctions"
)

type Day1 struct {
	data     []string
	calories []int
}

func (d *Day1) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day1\\a\\day1a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func (d *Day1) CalculateInts() {
	subtotal := 0
	for i := 0; i < len(d.data); i++ {
		num, err := strconv.Atoi(d.data[i])
		if err != nil {
			if subtotal != 0 {
				d.calories = append(d.calories, subtotal)
			}
			subtotal = 0
			continue
		}
		subtotal += num
	}
}

func (d *Day1) ReturnMostCalories() int {
	if len(d.calories) > 0 {
		sort.Ints(d.calories)
		return d.calories[len(d.calories)-1:][0]
	}
	return -1
}

func (d *Day1) ReturnTopXCalories(count int) int {
	sum := 0
	if len(d.calories) > count {
		for _, v := range d.calories[len(d.calories)-count:] {
			sum += v
		}
	}
	return sum
}

func Day1all() {
	var day Day1
	day.New()
	day.CalculateInts()
	fmt.Println(day.ReturnMostCalories())
	fmt.Println(day.ReturnTopXCalories(3))
}
