package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (d Cat) Speak() string {
	return "Meow!"
}

type Cow struct {
}

func (d Cow) Speak() string {
	return "Moo!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals{
		fmt.Println(animal.Speak())
	}
}
