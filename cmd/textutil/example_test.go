package textutil_test

import (
	"fmt"

	"github.com/jeff283/go-playground/cmd/textutil"
)

// ExampleReverse attaches to the Reverse function's doc page as a
// runnable example. The "Output:" comment is checked by `go test`,
// so this doubles as a real test.
func ExampleReverse() {
	fmt.Println(textutil.Reverse("Résumé"))
	// Output: émuséR
}

// ExamplePad attaches to Pad's doc page.
func ExamplePad() {
	fmt.Printf("%q\n", textutil.Pad("hi", 6))
	// Output: "  hi  "
}

// Example (no suffix) attaches to the *package's* overview page rather
// than to any single function, useful for a general usage snippet.
func Example() {
	c := &textutil.Counter{}
	c.Add("hello")
	c.Add("world")
	fmt.Println(c.Total())
	// Output: 10
}
