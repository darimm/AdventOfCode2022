package Day4all

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/darimm/AOCFunctions"
)

type Day4 struct {
	data []string
}

func (d *Day4) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day4\\a\\day4a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func Day4all() {
	var day Day4
	day.New()

	totaloverlap := 0
	partialoverlap := 0
	for _, v := range day.data {
		ranges := strings.Split(v, ",")
		if len(ranges) == 2 {
			first := strings.Split(ranges[0], "-")
			fLow, err := strconv.Atoi(first[0])
			if err != nil {
				fmt.Println("First Low NaN")
			}
			fHigh, err := strconv.Atoi(first[1])
			if err != nil {
				fmt.Println("First High NaN")
			}
			second := strings.Split(ranges[1], "-")
			sLow, err := strconv.Atoi(second[0])
			if err != nil {
				fmt.Println("Second Low NaN")
			}
			sHigh, err := strconv.Atoi(strings.Trim(second[1], "\r"))
			if err != nil {
				fmt.Println("Second High NaN")
			}
			if (fLow >= sLow && fHigh <= sHigh) || (fLow <= sLow && fHigh >= sHigh) {
				totaloverlap++
			}
			if AOCFunctions.Maxint(fLow, sLow) <= AOCFunctions.MinInt(fHigh, sHigh) {
				partialoverlap++
			}
		}
	}
	fmt.Println(totaloverlap)
	fmt.Println(partialoverlap)
}
