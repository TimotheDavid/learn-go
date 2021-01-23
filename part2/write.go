package main


import (
	"os"
	"fmt"
	"io/ioutil"
)
func main(){

	file, err := os.Create("file.txt")

	if err != nil {
		return 
	}

	defer file.Close()

	file.WriteString("Hello everyone")

	read , err:= ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Print(err)
	}


	str := string(read)

	fmt.Println(str)

}