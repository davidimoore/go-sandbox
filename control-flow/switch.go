package main

import (
	"fmt"
	"math/rand"
	"time"	
)

func switchStandard(num int){
	result := ""
	// initialize switch value
	switch num {
	case 1:
		result = "It's Sunday!"
	case 7:
		result = "It's Sunday!"
	default:
		result = "It's a weekday!"
	}
	fmt.Println(result)
}

func swithWithInitializedValue(){
	result := ""
	// initialize switch value
	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "It's Sunday!"
	case 7:
		result = "It's Sunday!"
	default:
		result = "It's a weekday!"
	}
	fmt.Println(result)
}

func switchWithoutInitalValue(num int){
	result := ""
	switch {
	case num < 0:
		result = "Less than zero"
		//fallthrough
	case num == 0:
		result = "Equal to zero"
	default:
		result = "Greater than zero"
	}
	fmt.Println(result)
}

func generateRandomNumb(){
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)
}

func main() {
	switchWithoutInitalValue(-42)
}
