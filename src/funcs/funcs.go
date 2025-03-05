package funcs

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Multiples_of_3_and_5(max int) int {

	return sumOfMultiples(max, 3) + sumOfMultiples(max, 5) - sumOfMultiples(max, 15)
}

func Fib_numbers(n int) []int {
	res := []int{1, 2}
	n1 := 1
	n2 := 2
	sum := 0
	for n1+n2 < n {
		sum = n1 + n2
		res = append(res, sum)
		n1 = n2
		n2 = sum
	}
	return res
}

func sumOfMultiples(limit, n int) int {
	k := (limit - 1) / n
	return n * k * (k + 1) / 2
}
