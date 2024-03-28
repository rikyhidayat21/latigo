package main

import "fmt"

// Pointer Part 1

func main() {
	// declare variable of type int with a value of 10
	count := 10

	// display the "value of" and "address of" count variable
	fmt.Printf("count:\tValue of %v\t Addr of %v\n", count, &count)

	// pass the "value of" the count to a function
	fmt.Println("triggered increment function")
	// this is a value semantics, or pass by value (everything in Go is pass by value)
	// pass by value means, the params count is the copied value of count variable
	// so there's no side effect to the main count variable, even in the function
	// increment already make changes to the count params, which count++
	increment(count)

	// display after the function called to check count variable
	fmt.Printf("count:\tValue of %v\t Addr of %v\n", count, &count)
}

func increment(count int) int {
	count++
	fmt.Printf("count after increment:\tValue of %v\t Addr of %v\n", count, &count)
	return count
}
