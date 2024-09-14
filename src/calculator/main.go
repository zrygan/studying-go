package main

import "fmt"

// returns (x - y)
func add(x int, y int) int {
	return x + y
}

// returns (x - y) if x is greater than y. Otherwise returns (-999999999).
func sub(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return -999999999
	}
}

// returns the choice of the user
func menu() int {
	var choice int
	fmt.Println("[1]\tadd\n[2]\tsub")
	fmt.Print("Select your choice: ")
	fmt.Scan(&choice)

	fmt.Printf("You chose %v\n\n", choice)

	return choice
}

func inn() (int, int) {
	var x, y int

	fmt.Print("Value for x: ")
	fmt.Scan(&x)

	fmt.Print("Value for y: ")
	fmt.Scan(&y)

	return x, y
}

func main() {
	var (
		operation = menu()
		res       int
	)
	x, y := inn()

	switch operation {
	case 1:
		res = add(x, y)
		break
	case 2:
		res = sub(x, y)
		break
	}

	fmt.Println("The result is", res)
}
