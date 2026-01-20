package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world ")
	githubVar := os.Getenv("API_SECRET")
	dbPassword := os.Getenv("DB_PASS")
	fmt.Println(githubVar)
	fmt.Println(dbPassword)
	fmt.Println("end ====== ")

}
