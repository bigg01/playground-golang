package main

import (
	"fmt"
	"os"
	"log"
)

var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)


func ArgServer() {
    fmt.Printf("%s \n", os.Args)
}

func main() {

	name := "oli"
	//var name2 char = "oli"

	fmt.Printf("hello %s world = %s", name, user)
	ArgServer()

	fd, err := os.Open("test.go")
    	if err != nil {
        	log.Fatal(err)
	}
	_ = fd

	// test function from test.go


}
