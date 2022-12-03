package Day3all

import (
	"fmt"

	"github.com/darimm/AOCFunctions"
)

type Day3 struct {
	data []string
}

func (d *Day3) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day3\\a\\day3a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func getItemPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r - 38)
	}
	return -1
}

func (d *Day3) Day3a() {
	sum := 0
SACK:
	for _, v := range d.data {
		left := make(map[rune]struct{})
		l := []rune(v)[0 : len([]rune(v))/2]
		r := []rune(v)[len([]rune(v))/2 : len([]rune(v))]

		for _, i := range l {
			left[i] = struct{}{}
		}
		for _, j := range r {
			_, ok := left[j]
			if ok {
				sum += getItemPriority(j)
				continue SACK
			}
		}
	}
	fmt.Println(sum)
}

func (d *Day3) Day3b() {
	sum := 0
GROUPOFELVES:
	for i := 0; i < len(d.data)-1; i += 3 {
		key := make(map[rune]struct{})
		key2 := make(map[rune]struct{})
		for _, v := range d.data[i] {
			key[v] = struct{}{}
		}
		for _, v := range d.data[i+1] {
			_, ok := key[v]
			if ok {
				key2[v] = struct{}{}
			}
		}
		for _, v := range d.data[i+2] {
			_, ok := key2[v]
			if ok {
				sum += getItemPriority(v)
				continue GROUPOFELVES
			}
		}
	}
	fmt.Println(sum)
}

func Day3all() {
	var day Day3
	day.New()
	day.Day3a()
	day.Day3b()
}
