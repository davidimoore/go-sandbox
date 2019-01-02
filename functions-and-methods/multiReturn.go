package main

import "fmt"

func main() {
  n1, l1 := FullName("David", "Moore")
  n2, l2 := FullNameNakedReturn("David", "Moore")
  fmt.Printf("Fullname: %v, number of chars: %v\n\n", n1, l1)
  fmt.Printf("FullNameNakedReturn: %v, number of chars: %v\n\n", n2, l2)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
