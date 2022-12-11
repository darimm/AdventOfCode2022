package Day6all

import (
	"fmt"

	"github.com/darimm/AOCFunctions"
)

type Day6 struct {
	data []string
}

func (d *Day6) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day6\\a\\day6a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func Day6all() {
	var day Day6
	day.New()
	day.Day6a()
	day.Day6b()
}

func (d *Day6) Day6a() {
	mySignal := d.data[0]
	for i := 3; i < len(mySignal); i++ {
		if !testDuplicateCharacters(mySignal[i-3 : i+1]) {
			continue
		} else {
			fmt.Printf("%s%s%s%s\r\n", string(mySignal[i-3]), string(mySignal[i-2]), string(mySignal[i-1]), string(mySignal[i]))
			fmt.Println(i + 1)
			break
		}
	}
}

func (d *Day6) Day6b() {
	mySignal := d.data[0]
	for i := 13; i < len(mySignal); i++ {
		if !testDuplicateCharacters(mySignal[i-13 : i+1]) {
			continue
		} else {
			fmt.Println(i + 1)
			break
		}
	}
}

func testDuplicateCharacters(s string) bool {
	keys := make(map[rune]struct{})
	for _, v := range s {
		keys[v] = struct{}{}
	}
	return len(keys) == len(s)
}
