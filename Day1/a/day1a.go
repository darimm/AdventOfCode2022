package Day1a

import (
	"fmt"

	"github.com/darimm/AOCFunctions"
)

func Day1a() {
	lines, err := AOCFunctions.ReadFile(".\\Day1\\a\\day1a.txt")
	if err != nil {
		panic("Could not read input")
	}
	fmt.Println(lines)
}
