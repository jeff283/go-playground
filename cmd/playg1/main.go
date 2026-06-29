package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Types
	fmt.Println("-- TYPES --")
	var number int = -254
	fmt.Println(number)
	var unumber uint = 255
	fmt.Println(unumber)
	fmt.Println()

	var floatNum float64 = 0.552
	fmt.Println(floatNum)
	fmt.Println()

	// Type casting
	fmt.Println("-- TYPE CASTING --")

	// Weird things happen when (check below), because
	/*
	* Casting a negative int to uint doesn't change the underlying bits,
	* it just reinterprets them. A negative int -n becomes 2^64 - n.
	* e.g. uint(-254) == 18446744073709551362 (2^64 - 254)
	 */
	fmt.Println(uint(number))
	fmt.Println()

	var floatNum32 float32 = 6.7
	var intNum32 int32 = 67
	var results float32 = floatNum32 + float32(intNum32)
	fmt.Println(results)
	fmt.Println()

	// Ops
	fmt.Println("-- OPERATIONS --")
	var num1 int = 3
	var num2 int = 2
	// Rounded Down
	fmt.Print("3/2: ")
	fmt.Println(num1 / num2)

	// Remainder
	fmt.Print("Remainder: ")
	fmt.Println(num1 % num2)

	// Type Casting
	fmt.Print("Type Casted to Float: ")
	fmt.Println(float32(num1) / float32(num2))

	// STRINGS
	fmt.Println("-- STRINGS --")

	// Strings are normal mostly
	var basicString string = "Guess What, I am String?\n"
	fmt.Println(basicString)
	var wysiwygString string = `
	Line 1
	Line 2
	Line x
	`
	fmt.Println(wysiwygString)

	fmt.Println("-- Strings Counts --")
	// Count bytes not chars in a string
	var oneByteString string = "1"
	var twoBytesString string = "β"

	fmt.Println("Bytes Count: ")
	fmt.Println(len(oneByteString))
	fmt.Println(len(twoBytesString))

	// Calculating actual chars
	fmt.Println("Chars Count: ")
	fmt.Println(utf8.RuneCountInString(oneByteString))
	fmt.Println(utf8.RuneCountInString(twoBytesString))
	fmt.Println()

	fmt.Println("-- Bools --")

	var isBully bool = false
	fmt.Println(isBully)
	fmt.Println()

	weirdSyntax, goWeird := "I am weird in go", "Go weird assignment"
	fmt.Println(weirdSyntax)
	fmt.Println(goWeird)
	fmt.Println()

	const Constantinople string = "Even Photoshop can't change me"
	const Acoast string = " I am having so much fun with this"

	fmt.Println(Constantinople)
	fmt.Println(Acoast)
	fmt.Println()
}
