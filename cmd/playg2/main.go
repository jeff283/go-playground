package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	Afunction()

	fmt.Println("Let us do some maths")

	var num1 = rand.Intn(100000) + 1000
	num2 := rand.Intn(100) + 10
	num3 := rand.Intn(10)

	var results, err = doMaths(num1, num2, num3)

	if err != nil {
		fmt.Println("Oops, maths failed")
		fmt.Println(err.Error())

	} else {
		fmt.Printf("We did some maths and here is what we got %.2f\n", results)
		fmt.Println()
	}

	switchStatements(results, err)
	fmt.Println()

}

func Afunction() {
	fmt.Println("I am a function")
}

func doMaths(num1 int, num2 int, num3 int) (float64, error) {

	if num1 == 0 || num2 == 0 || num3 == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	var result float64 = float64(num1) / float64(num2) / float64(num3)

	return result, nil
}

func switchStatements(results float64, err error) {
	// Option 1
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case results > 0.0:
		fmt.Printf("We did some maths and here is what we got %.2f\n", results)
	default:
		fmt.Println("The Maths gods are not on your side today")

	}

	// Option we know what we are doing
	switch results {
	case 0.0:
		fmt.Println("You got the magic number, you are the best")
	default:
		fmt.Printf("We did some maths and here is what we got %.2f\n", results)

	}
}
