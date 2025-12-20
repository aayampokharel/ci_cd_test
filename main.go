package main

import "fmt"

func addTwoNumbers(a, b int) int {
	return a + b
}
func main() {
	fmt.Print("====main() started ====")
	addTwoNumbers(1, 2)
	fmt.Print("====main() end ====")
}
