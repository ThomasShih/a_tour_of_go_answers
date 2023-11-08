package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	next_value, next_next_value := 0, 1
	return func() int {
		current_value := next_value
		next_value, next_next_value = next_next_value, current_value+next_next_value

		return current_value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
