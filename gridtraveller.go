package main

import (
	"fmt"
	"strconv"
)

func gridTraveller(n int, m int) int {
	if n == 1 && m == 1 {
		return 1
	}
	if n == 0 || m == 0 {
		return 0
	}

	return gridTraveller(n-1, m) + gridTraveller(n, m-1)
}

func getGridNumber() (int, int) {
	var n int
	var m int
	fmt.Println("Enter the row number:")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the column number:")
	fmt.Scanf("%d", &m)
	return n, m
}

func gridTravellerUsingMemo(n int, m int, memo map[string]int) int {
	key := strconv.Itoa(n) + ":" + strconv.Itoa(m)
	if memo[key] != 0 {
		return memo[key]
	}
	if n == 1 && m == 1 {
		return 1
	}
	if n == 0 || m == 0 {
		return 0
	}

	memo[key] = gridTravellerUsingMemo((n-1), m, memo) + gridTravellerUsingMemo(n, (m-1), memo)
	return memo[key]
}
