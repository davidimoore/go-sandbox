package main

import "fmt"

func checkP(p *int){
	if p != nil {
		fmt.Println("value of p: ", *p)
	} else {
		fmt.Println("p is nil")
	}
}

func main() {
	//var p *int
	//checkP(p)

	v := 42
	p := &v
	checkP(p)

	v = 56
	checkP(p)

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value 1: ", *pointer1)

	*pointer1 = *pointer1 /31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}

