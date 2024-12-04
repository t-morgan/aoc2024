package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		return
	}

	arg := os.Args[1]
	switch arg {
	case "1":
		dayOne()
	case "2":
		dayTwo()
	case "3":
		dayThree()
	case "4":
		dayFour()
	default:
		fmt.Println("Invalid argument:", arg)
	}
}
