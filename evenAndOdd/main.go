package main

import "fmt"

func main() {
	ints := []int{}
	for i := 0; i < 11; i++ {
		ints = append(ints, i)
	}

	for _, num := range ints {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
