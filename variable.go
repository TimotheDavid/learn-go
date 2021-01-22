package main

import (
	"fmt"
	"time"
)


func main(){
	
	// declare variable for the age and the date of birth
	var age int
	var date int

	fmt.Printf("give me your age ( in year ;) ")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("give me your year of birth ")
	fmt.Scanf("%d", &date)

	fmt.Println(age,date)
	
	// calculate the actual year 
	actual := date + age 
	time := time.Now().Year()
	fmt.Printf("%d \n", actual)
	// calculate if time is diff since the year
	if actual == time {
		fmt.Printf("good jobs ;)")
	}else{
		if(time > actual){
			res := time - actual
			fmt.Printf("error of %d year :X", res)
		}else if(time < actual) {
			res := actual - time
			fmt.Printf("error of %d year :X", res)
		}
		
	}

	

}