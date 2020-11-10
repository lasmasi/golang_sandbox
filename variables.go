package main

import "fmt"

func main() {

	// infers type
	var a = "initial"
	fmt.Println(a)

	// can init multiple vars at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// will initialise as zero valued if no value provided
	var e int
	fmt.Println(e)

	// the same as:
	// var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
