package Day2all

import (
	"fmt"
	"strings"

	"github.com/darimm/AOCFunctions"
)

type Day2 struct {
	data []string
}

func (d *Day2) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day2\\a\\day2a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func (d *Day2) TransformAndTest() int {
	score := 0
	for _, v := range d.data {
		game := strings.Fields(v)
		if len(game) != 2 {
			continue
		}
		score += TestAndScore(game[0], game[1])
	}
	return score
}

func (d *Day2) TransformAndTest2() int {
	score := 0
	for _, v := range d.data {
		game := strings.Fields(v)
		if len(game) != 2 {
			continue
		}
		score += TestAndScore2(game[0], game[1])
	}
	return score
}

type RPS int

const rock RPS = 1
const paper RPS = rock << 1
const scissors RPS = paper << 1

func TestAndScore(opp, me string) int {
	var rpstest = map[string]RPS{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	var points = map[RPS]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}

	mythrow := rpstest[me]
	oppthrow := rpstest[opp]

	if mythrow&oppthrow != 0x0 {
		return points[mythrow] + 3 //draw
	}
	if mythrow>>1 == oppthrow || (mythrow == rock && oppthrow == scissors) {
		return points[mythrow] + 6 //win
	}
	return points[mythrow] //lose
}

func TestAndScore2(opp, me string) int {

	var rpstest = map[string]RPS{
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	var points = map[RPS]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}

	oppthrow := rpstest[opp]

	switch me {
	case "X": //lose
		if oppthrow == rock {
			return points[scissors]
		}
		return points[oppthrow>>1]
	case "Y": //draw
		return points[oppthrow] + 3
	case "Z": //win
		if oppthrow == scissors {
			return points[rock] + 6
		}
		return points[oppthrow<<1] + 6
	default:
		return 0
	}
}

func Day2all() {
	var day Day2
	day.New()
	fmt.Println(day.TransformAndTest())
	fmt.Println(day.TransformAndTest2())
}
