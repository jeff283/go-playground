package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("-- STRINGS and RUNES --")
	exploreStrings()
	fmt.Println()

	fmt.Println("-- RUNES --")
	exploreRunes()
	fmt.Println()

	fmt.Println("-- String Building --")
	stringBuilding()
	fmt.Println()
}

func exploreStrings() {
	var aString = "Résumé"

	fmt.Println(aString)

	indexed := aString[0]
	fmt.Println(indexed)
	fmt.Printf("Value: %v and Type: %T \n", indexed, indexed)

	fmt.Printf("Weird half index of é is: %v , rather than 233 \n", aString[1])

	fmt.Println("\nIterations:")

	for i, v := range aString {
		fmt.Println(i, v)
	}
	fmt.Println()

}

func exploreRunes() {
	var aString = []rune("Résumé")

	fmt.Printf("Index 0: %v and Index 1 (é): %v \n", aString[0], aString[1])

	fmt.Println("\nContinous Index: ")
	for i, v := range aString {
		fmt.Println(i, v)
	}
	fmt.Println()

	var singleRune = 'r'
	fmt.Printf("A Rune: %v\n", singleRune)

}

func stringBuilding() {

	var stringSlice = []string{"'", "A", " ", "S", "t", "r", "i", "n", "g", "'"}
	var concatString = ""

	// Strings are immutable
	// Concat Strings needs to use a new string
	fmt.Println("Loop Builder: ")
	for _, v := range stringSlice {
		concatString += v
	}
	fmt.Println(concatString)

	// String builders
	var stringBuilder strings.Builder
	fmt.Println("\nStrings Builder: ")
	for _, v := range stringSlice {
		stringBuilder.WriteString(v)
	}
	fmt.Println(stringBuilder.String())

}
