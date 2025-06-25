package main

import "fmt"

func main() {
	var myBoolean bool = true
	var anotherBoolean = false 

	var truth = 3 <= 5
	var falsehood = 10 != 10
	var res1 = 10 > 20 && 5 == 5    
	var res2 = 2*2 == 4 || 10%3 == 0 
	fmt.Println(myBoolean, anotherBoolean, truth, falsehood, res1, res2)
}
