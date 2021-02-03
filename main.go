package main

import (
	"fmt"
	"go-general/random"
)

func main() {
	fmt.Println(random.String(10))
	fmt.Println(random.Int(10))
}