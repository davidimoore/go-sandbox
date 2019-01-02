package main

import (
	"fmt"
	"sort"
)

func printStates(states map[string]string){
	fmt.Println(states)
}

func alphabetizeStates(states map[string]string){
	// create a slice for keys
	keys := make([]string, len(states))

	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	//sort keys
	sort.Strings(keys)
	fmt.Println("\nSorted")

	// iterate through keys and print value of states in order
	for i:= range keys {
		fmt.Println(states[keys[i]])
	}
}

func main() {
	states := make(map[string]string)
	printStates(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	printStates(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	printStates(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	alphabetizeStates(states)
}