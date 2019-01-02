package main

import "fmt"

func standardLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func nonStandardLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}
	endofprogram : fmt.Println("end of program")
}

func standardLoopForSlice(slice []string) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func rangeLoopForSlice(slice []string) {
	for i := range slice {
		fmt.Println(slice[i])
	}
}

func main() {
	standardLoop()
	nonStandardLoop()

	colors := []string{"Red", "Green", "Blue"}
	standardLoopForSlice(colors)
	rangeLoopForSlice(colors)
}
