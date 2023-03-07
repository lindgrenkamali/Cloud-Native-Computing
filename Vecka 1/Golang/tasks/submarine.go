package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PrintPosition func(int, int)



type Moveinstruction struct {
	DEPTH int
	STEP int
   
	   position PrintPosition
   }

func fileAndScanner(path string) *bufio.Scanner {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	return scanner

}

func Submarine2() {

	scanner := fileAndScanner("files/day2input.txt")
	m := Moveinstruction{
		DEPTH: 0,
		STEP:  0,
		position: func(Depth int, Step int) {
			fmt.Println("DEPTH:", Depth, "HORIZONTAL:", Step)
		},
	}

	for scanner.Scan() {
		line := scanner.Text()
		var kommando []string = strings.Split(line, " ")
		var currentValue int64
		var err error
		currentValue, err = strconv.ParseInt(kommando[1], 10, 0)

		if err == nil {

			switch kommando[0] {
			case "forward":
				m.STEP += int(currentValue)

			case "down":
				m.DEPTH += int(currentValue)

			case "up":
				m.DEPTH -= int(currentValue)
			}

		}
		m.position(m.DEPTH, m.STEP)
	}
}

func Submarine1() {

	scanner := fileAndScanner("files/day1input.txt")
	var currentValue int64
	var lastValue int64 = -1

	var err error

	var decreaseCount int = 0

	var increaseCount int = 0

	for scanner.Scan() {
		line := scanner.Text()

		currentValue, err = strconv.ParseInt(line, 10, 0)
		if err == nil {
			if lastValue == -1 {

			} else {
				if currentValue > lastValue {
					increaseCount++
					fmt.Println(currentValue, "- större än", lastValue, "- ANTAL=", increaseCount)
				} else if currentValue < lastValue {
					decreaseCount++
					fmt.Println(currentValue, "- mindre än", lastValue, "- ANTAL=", decreaseCount)
				}
			}

		}
		lastValue = currentValue
	}

}
