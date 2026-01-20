package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=======start=========")
	fmt.Println(os.Getenv("DB_PASS"))
	fmt.Println(os.Getenv("PASS_FROM_SECRETS"))
	fmt.Println("=======end=========")
}
