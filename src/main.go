package main

import (
	"fmt"

	"github.com/jnsoft/goEuler/src/funcs"
	"github.com/jnsoft/jngo/bag"
)

func main() {
	// Get a greeting message and print it.
	message := funcs.Hello("Gladys")
	fmt.Println(message)

	b := bag.NewBag[int]()

	b.Add(4)
	b.Add(3)
	b.Add(3)

	fmt.Println(b.Count(3))
	fmt.Println(b.Count(2))

}
