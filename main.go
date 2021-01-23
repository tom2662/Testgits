package main

import (
	"fmt"

	L "./applib"
)

func main() {
	fmt.Println(" ")
	// soal 1
	input := []int{1, 1, 0, 1, 1, 1}

	res := L.GetMaxLength(input)

	fmt.Println("Test 1 :")
	fmt.Println(res)
	fmt.Println(" ")

	// soal 2

	input2 := []string{"h", "e", "l", "l", "o"}

	res2 := L.Reverse(input2)

	fmt.Println("Test 2 :")
	fmt.Println(res2)
	fmt.Println(" ")

	// soal 3

	input3 := "{ [ ( ) ] }))"

	res3 := L.FindBalance(input3)

	fmt.Println("Test 3 :")
	fmt.Println(res3)
	fmt.Println(" ")

}
