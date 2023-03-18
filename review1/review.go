package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Masukan nilai n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i, "FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
