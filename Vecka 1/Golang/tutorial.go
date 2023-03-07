package main

import (
	"Go-Tutorial/tasks"
	"fmt"
)

func main() {
	tasks.Submarine1()
	tasks.Submarine2()

}




func slice(){
	s := [6]int {2, 5, 10, 33, 58}
	var a []int = s[1: 4] 
	fmt.Println(a)
}

func mapping(){
	m := make(map[string]int)
	m["nitton"] = 19
	m["tvåhundra"] = 200
	fmt.Println(m["nitton"])
	fmt.Println(m["tvåhundra"])
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

