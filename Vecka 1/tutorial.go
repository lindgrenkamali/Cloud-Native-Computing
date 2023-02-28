package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	submarine2()
}

type Moveinstruction struct {
 DEPTH int
 STEP int
}

func submarine2(){
	file, _ := os.Open("day2input.txt")

	scanner := bufio.NewScanner(file)
	m := Moveinstruction {
		DEPTH: 0,
		STEP: 0,
	};

	for scanner.Scan() {
		line := scanner.Text()
		var kommando[] string = strings.Split(line, " ");
		var currentValue int64;
		var err error;
		currentValue, err = strconv.ParseInt(kommando[1], 10, 0);

		if(err == nil){

		switch kommando[0] {
		case "forward":
			m.STEP += int(currentValue)

		case "down":
			m.DEPTH += int(currentValue)

		case "up":
			m.DEPTH -= int(currentValue)
		}
			
	}
	fmt.Println("DEPTH:", m.DEPTH, "HORIZONTAL:", m.STEP)
	}
}

func submarine1(){
	file, _ := os.Open("day1input.txt")

	scanner := bufio.NewScanner(file)

	var currentValue int64;
	var lastValue int64 = -1;

	var err error;

	var decreaseCount int = 0;

	var increaseCount int = 0;

	for scanner.Scan() {
		line := scanner.Text()
		
		currentValue, err = strconv.ParseInt(line, 10, 0);
		if(err == nil){
			if(lastValue == -1){

			}else {
				if(currentValue > lastValue){
					increaseCount++;
					fmt.Println(currentValue, "- större än", lastValue, "- ANTAL=", increaseCount)
				} else if(currentValue < lastValue) {
					decreaseCount++
					fmt.Println(currentValue, "- mindre än", lastValue, "- ANTAL=", decreaseCount)
				}
			}

		}
		lastValue = currentValue
		}

	}


func hundred(){
	fmt.Print("Skriv in nummer: ");
	var nummer int;
	fmt.Scanln(&nummer)

	if (nummer >= 100) {
		fmt.Print(true);
	} else {
		fmt.Print(false);
	}
}

func printText() {
	print("2023")
}

func printName() {
	var name string = "David"
	var age int = 22

	fmt.Println("Mitt namn är", name, "och jag är", age, "år gammal");

}

func printInput() {
	fmt.Print("Skriv in ditt förnamn: ");

	var förnamn string;

	var efternamn string;

	fmt.Scanln(&förnamn)

	fmt.Print("Skriv in ditt efternamn: ");

	fmt.Scanln(&efternamn)

	fmt.Println("Du heter", efternamn, förnamn)

}

func printOperators(){
	var tal1 int;
	var tal2 int;
	fmt.Scanln(&tal1)
	fmt.Scanln(&tal2)

	fmt.Println(tal1 * tal2)
	fmt.Println(tal1 ^ 2)

}

func printFloat(){
	var tal1 float32;
	var tal2 float32;
	fmt.Scanln(&tal1)
	fmt.Scanln(&tal2)

	fmt.Println(tal1 * tal2)

}

func printNameRepeat(){
	fmt.Print("Skriv in ditt förnamn: ");
	

	var namn string;
	var namnUpprepat string;

	fmt.Scanln(&namn)
	for i := 0; i < 5; i++ {
		namnUpprepat += namn;
	}
	fmt.Println(namnUpprepat)
}

