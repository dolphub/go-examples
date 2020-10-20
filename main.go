package main

import (
	"fmt"

	"github.com/dolphub/go-examples/service"
)

func main() {
	x := service.Product(5, 3)
	fmt.Println(x)
	service.PrintName("Go")
}
