package main

import (
	"fmt"
)

func main() {

	fmt.Println("arrays e slices")

	var listUser [5]string
	fmt.Println(listUser)

	listUser[0] = "andre"

	fmt.Println(listUser)

	arrayFunc := [2]string{"anhdre", "mendes"}
	fmt.Println(arrayFunc)

	arrayEmp := [...]string{"anhdre", "mendes", "", ""}
	fmt.Println(arrayEmp)

	//slice

	sliceFunc := []int{1, 1, 2, 3, 4, 5, 6, 76}
	fmt.Println(sliceFunc)

	var sliceFunc2 []int

	sliceFunc2 = append(sliceFunc2, 1, 2, 4, 5, 7, 8)
	fmt.Println(sliceFunc2)

	//array interno

	slice3 := make([]string, 13, 20)
	fmt.Println(slice3)
	fmt.Println(len(slice3))

}
