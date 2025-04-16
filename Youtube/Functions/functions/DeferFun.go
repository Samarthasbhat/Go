package functions

import (
	"fmt"
)
//  Defer statement captures a function(exit) call to run later
//  If we have more than a defer it works on the function scope

func DeferFun (){

	num := 10

	defer fmt.Println(num*100)

	num =11
	fmt.Println(num)
}

func DoIt() (a int){
	defer func() {
		a = 2
	}()
	 
	a= 1
	return 
}