package main

import "fmt"

func compareToZero(x int) {
	var result string
	if x < 0 {
		result = "Less than zero"
	} else if (x == 0) {
		result = "Equal to zero"
	} else {
		result = "Greater than or equal to zero"
	}
	fmt.Println(result)
}

func main() {
	compareToZero(-42)

}
