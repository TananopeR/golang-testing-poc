package b

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintResults(t *testing.T) {
	fmt.Printf("%v\n", &doSomething)
	origDoSomething := doSomething
	fmt.Printf("%v %v\n", &doSomething, &origDoSomething)
	doSomething = func() string {
		// Insert fake implementation here
		return "complete"
	}

	b := NewB()
	defer func() {
		fmt.Printf("%v %v\n", &doSomething, &origDoSomething)
		fmt.Println("defer")
		fmt.Println(b.Results())
	}()

	result := b.Results()
	fmt.Printf("%v  %v %T\n", result, b, b)

	assert.Equal(t, "complete", result)
}

func TestPrintResultsDone(t *testing.T) {

	b := NewB()
	result := b.Results()
	fmt.Printf("%v  %v %T\n", result, b, b)
	assert.Equal(t, "done", result)
}
