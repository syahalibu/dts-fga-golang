package main

import "fmt"

func main() {

	i := 0
	for i < 5 {
		fmt.Printf("Nilai i = %d\n", i)
		i++
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			var characters = []rune{}
			characters = append(characters, 'С', 'А', 'Ш', 'А', 'Р', 'В', 'О')
			nn := 0
			for _, value := range characters {
				fmt.Printf("character %U '%c' starts at by byte position %d\n", value, value, nn)
				nn += 2
			}
		} else {
			fmt.Println("Nilai j =", j)
		}
	}
}
