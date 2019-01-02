package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file")
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	myFunc()
	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("undeferred statement")
}

func myFunc() {
	defer fmt.Println("myFunc DEFERRED!")
	fmt.Println("myFunc NOT DEFERRED!")
}
