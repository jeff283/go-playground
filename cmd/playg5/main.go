package main

import (
	"fmt"
	"strings"
)

type Owner struct {
	name string
}

type Meds struct {
	name string
}

type Cat struct {
	name      string
	breed     string
	age       int
	ownerInfo Owner
	Meds
}

type Dog struct {
	name  string
	breed string
	color string
	age   int
}

type Animal interface {
	breedName() string
	getAge() int
}

func (c Cat) breedName() string {
	var bName strings.Builder
	bName.WriteString(c.name)
	bName.WriteString(c.breed)
	return bName.String()
}

func (c Cat) getAge() int {
	return c.age
}

func (d Dog) breedName() string {
	var bName strings.Builder
	bName.WriteString(d.name)
	bName.WriteString(d.breed)
	return bName.String()
}

func (d Dog) getAge() int {
	return d.age
}

func isOlderThan(a Animal, age int) bool {
	return a.getAge() > age
}

func main() {
	var myCat Cat = Cat{name: "Kat", breed: "Mia", ownerInfo: Owner{"Kamau"}, Meds: Meds{"Amoxil"}, age: 2}

	myCat.name = "Katu"

	fmt.Println("CAT")
	fmt.Printf("Type: %T \n", myCat)
	fmt.Println(myCat)
	fmt.Println(myCat.breedName())
	fmt.Printf("Cat is older than 2: %v\n", isOlderThan(myCat, 2))
	fmt.Println()

	fmt.Println("DOG")
	var myDog = Dog{name: "Jamis", breed: "Labrador", color: "Brown", age: 4}
	fmt.Printf("Type: %T \n", myDog)
	fmt.Println(myDog)
	fmt.Println(myDog.breedName())
	fmt.Printf("Dog is older than 2: %v\n", isOlderThan(myDog, 2))
	fmt.Println()

	fmt.Println()
	var anonStruct = struct {
		name  string
		class int8
	}{name: "anon", class: 8}
	fmt.Printf("Anonymous Stuct: %v \n", anonStruct)

}
