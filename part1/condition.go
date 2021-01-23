package main

import (
	"fmt"
)


func main(){
	
	var number int
	fmt.Println("Give us a number ;) ")
	fmt.Scanf("%d", &number)

	if(number > 0){
		res := number / 2
		fmt.Printf("the division per 2 of %d is %d", number, res) 


	}



}