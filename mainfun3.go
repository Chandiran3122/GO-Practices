package main

import "fmt"

func main() {
	
	var a int64 = 4
	var b int = int(a) 

	var c float64 = 6.5

	
	var result = float64(b) + c  

	fmt.Println(result)

	var myInt int = 65
	var myUint uint = uint(myInt)
	var myFloat float64 = float64(myInt)

	fmt.Println(myInt, myUint, myFloat)
}
