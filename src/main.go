package main

import (
	"fmt"

	"github.com/jnsoft/goEuler/src/funcs"
)

func main() {
	// Get a greeting message and print it.
	message := funcs.Hello("Gladys")
	fmt.Println(message)
}
