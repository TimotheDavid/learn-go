package main


import (
	"fmt"
	"os"
)


func main(){

	if _ , err := os.Stat("text.tt"); os.IsNotExist(err) {
		fmt.Println("File does not exist\n")
		return
	}

	fmt.Println("the file exist :D")



}