package main

import (
	"fmt"

	"github.com/michaelmdresser/custom-action-demo-code/pkg/mylib"
)

func main() {
	fmt.Println("Hello! This is the variant of the project that will be run in my custom GitHub Action.")
	mylib.PrintSomething()
}
