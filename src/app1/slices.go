package main

import (
	"fmt"
	"os"
)

func isDumbTerm() bool {
	return os.Getenv("TERM") == "dumb"
}


func main() {

	name := "oli"

	fmt.Sprint(name)
	
	var colors = []string{"Green", "Read", "Yellow"}

	fmt.Println(colors)

	fmt.Println("\x1b[31;1mHello, World!\x1b[0m")
	
	fmt.Println(isDumbTerm())
}
