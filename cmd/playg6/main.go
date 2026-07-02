package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	yob  uint32
}

func (p *Person) getAge() int {
	year := time.Now().Year()
	return year - int(p.yob)
}

func main() {

	fmt.Println("-- POINTERS --")
	var p *int32
	var i int32

	fmt.Printf("P : %v \nI: %v \n", p, i)
	fmt.Println()

	p = new(int32(10))
	fmt.Printf("P : %v \n*P: %v \n", p, *p)
	*p = 20
	fmt.Printf("P : %v \n*P: %v \n", p, *p)
	p = &i
	fmt.Printf("P : %v \n*P: %v \n", p, *p)
	fmt.Printf("i=%v and address: %p \n", i, &i)

	fmt.Println("\nSlice Pointers: ")
	// Copying Slices does not copy them fully but instead use pointers
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	fmt.Printf("Slice: %p \n", &slice[0])
	fmt.Printf("Slice Copy: %p \n", &sliceCopy[0])

	fmt.Println("\nFunction Pointers: ")
	var person = Person{"James", 2000}
	fmt.Printf("Person is %v years old\n", person.getAge())
}
