package main

import (
	"fmt"
)

func main(){

	arr1 := new([5]int)
	// get a normal array
	for i := 0; i< len(arr1); i++ {
		arr1[i] = i
		fmt.Println("char ", arr1[i])
	}
	ranges()
}

func ranges(){

	nums := []int{1,2,3,5,6,7,8,9,10}

	for _, num := range nums {
			fmt.Println(num)
	}

	for index , num := range nums {
		fmt.Println(index, num)

	}
}