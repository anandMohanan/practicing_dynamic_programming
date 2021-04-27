package main

import (
	"fmt"
)

func getNumber() int {
	var count int
	fmt.Println("Enter the count:")
	fmt.Scanf("%d", &count)
	return count
}

func fibonacci(count int) {
	var a []int
	n1 := 0
	n2 := 1
	a = append(a, n1)
	a = append(a, n2)
	for i := 0; i < count; i++ {
		n3 := n1 + n2
		a = append(a, n3)
		n1 = n2
		n2 = n3
	}
	fmt.Println(a)

}

func fibo(count int) {
	a, b := 0, 1
	for i := 0; i <= count-1; i++ {

		a, b = b, a+b

	}
	fmt.Printf("%d", a)
}

func fiboMemo(count int, memo map[int]int) int {
	if memo[count] != 0 {
		return memo[count]
	}
	if count <= 2 {
		return 1
	}

	memo[count] = fiboMemo(count-1, memo) + fiboMemo(count-2, memo)
	return memo[count]
}

// var m map[int]int

// func fibNum(n int, m map[int]int) int {

// 	m = make(map[int]int)
// 	if n == m[n] {
// 		return m[n]
// 	}
// 	if n <= 2 {
// 		return 1
// 	}
// 	m[n] = fibNum(n-1, m) + fibNum(n-2, m)
// 	return m[n]
// }
