package main

import (
	"fmt"
	"os"
)

func main() {
	value := os.Getenv("INFURA_API_KEY")
	fmt.Println("The value of INFURA_API_KEY is:", value)  
}