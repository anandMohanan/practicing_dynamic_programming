package main

import "fmt"

func main() {
	memo := make(map[int]int)
	gridMemo := make(map[string]int)
	fmt.Printf(" Please select a function by entering the number \n 1:fibonacci series \n 2:fibonacci number for a given number\n 3:fibonacci number for a given number using memoization \n 4: grid traveller \n 5: grid traveller using memoization \n")
	var s int
	fmt.Scanf("%d", &s)

	switch s {
	case 1:
		fibonacci(getNumber())
	case 2:
		fibo(getNumber())
	case 3:
		fmt.Println(fiboMemo(getNumber(), memo))
	case 4:
		fmt.Println(gridTraveller(getGridNumber()))
	case 5:
		fmt.Println(gridTravellerUsingMemo(30, 30, gridMemo))
	default:
		fmt.Println("please choose a correct number")
	}

}
