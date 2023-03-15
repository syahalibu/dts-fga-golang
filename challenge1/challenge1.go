package main

import "fmt"

func main() {

	var i int = 21
	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n%%\n", i)

	var j bool = true
	fmt.Printf("%t\n", j)

	//------------------------
	fmt.Println()

	russia := 'Я'
	fmt.Printf("%c\n", russia)

	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)

	i -= 6
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%U\n", russia)

	//------------------------
	fmt.Println()

	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%.6E", k)

}
