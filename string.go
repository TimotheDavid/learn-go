package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "This is a string"
	fmt.Println(str1)
	
	var str2 = "This is a second line \nhello third line"
	fmt.Println(str2)

	var str3 = "another string"
	var str4 = " fourth string"
	
	// join block of string 
	s := strings.Join([]string{str3, str4},",")
	fmt.Println(s)

	// print string as data 
	fmt.Printf("%d:%s", 2021, "year")




}