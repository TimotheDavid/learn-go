package main

import (
	"fmt"
	"io/ioutil"
)


func main(){
	// read the file from the io 
	file , err := ioutil.ReadFile("text.txt")
	// do you have error ? 
	if err != nil {
		fmt.Print(err)
	}
	// convert to string and get in a var 
	str := string(file)
	// print the file 
	fmt.Println(str)


}