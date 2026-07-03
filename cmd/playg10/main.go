package main

import "fmt"

func main() {
	var intNumbers = []int{1, 2, 3}
	fmt.Println("Sum of Ints: ", genericsSums(intNumbers))
	fmt.Println("Ints Slice Length: ", genericsAnyLen(intNumbers))
	fmt.Println()

	var floatNums = []float32{4.5, 6.7, 8.9}
	fmt.Println("Sum of Floats: ", genericsSums(floatNums))
	fmt.Println("Floats Slice Length: ", genericsAnyLen(floatNums))
	fmt.Println()

	var sliceStrings = []string{"a", "b", "c"}
	fmt.Println("String Slice Length: ", genericsAnyLen(sliceStrings))
	fmt.Println()

}

func genericsSums[T int | float32](numbers []T) T {
	var sum T

	for _, v := range numbers {
		sum += v
	}

	return sum

}

func genericsAnyLen[T any](slices []T) int {
	return len(slices)
}
