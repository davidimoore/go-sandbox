package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("my error string")
	fmt.Println(myError)
	 attendance := map[string]bool{
	 	"Ann": true,
	 	"Mike": true}

	 attended, ok := attendance["Mike"]

	 if ok {
	 	fmt.Println("Mike attended?", attended)
	 } else {
	 	fmt.Println("No info for Mike")
	 }
}
