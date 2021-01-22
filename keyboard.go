package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//call the os function 
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What's your name ?")
	
	//permit to read line from console
	city, _ := reader.ReadString('\n')
	fmt.Printf("You live in " + city)
	 /*
	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("give me a number")
	*/

	var out bool
	out = true
	// while loop ( for true)
	for out {
	var number int

	// read the io int 
	fmt.Println("give me a number between 1 and 10")
	fmt.Scanf("%d",&number)
			
	if  ( number <= 10 && number >= 1) {
		fmt.Println("check my boys !")
		// go out of the loop 
		out = false	
	}

}
}