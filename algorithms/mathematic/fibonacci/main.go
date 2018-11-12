package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	f_2 := 0
	f_1 := 1
	for index := 2; index < n; index++ {
		f := f_1 + f_2
		f_2 = f_1
		f_1 = f
	}
	return f_1 + f_2
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
