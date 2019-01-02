package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "implicitly typed string"
	fmt.Printf("str1: %v - type:%T \n", str1, str1)
	str2 := "explicitly typed string"
	fmt.Printf("str2: %v - type:%T \n", str2, str2)
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"

	fmt.Println("Equal?", lValue, "and", uValue, (lValue == uValue))
	fmt.Println("Equal non case sensitive?", lValue, "and", uValue, strings.EqualFold(lValue, uValue))

	fmt.Println(str1,"contains exp?", strings.Contains(str1, "exp"))
	fmt.Println(str2, "contains exp?", strings.Contains(str2, "exp"))
}
