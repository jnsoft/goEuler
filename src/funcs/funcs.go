package funcs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jnsoft/jngo/inthelper"
	"github.com/jnsoft/jngo/misc"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func TimeFunction(label string, f func() (interface{}, error)) {
	start := time.Now()
	result, err := f()
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s: %v (%s)\n", label, result, elapsed)
}

func Multiples_of_3_and_5(max int) int {

	return sumOfMultiples(max, 3) + sumOfMultiples(max, 5) - sumOfMultiples(max, 15)
}

func Sum_even_valued_fibs(limit int) int {
	ns := fib_numbers(limit)
	return misc.Reduce(ns, func(sum, number int) int {
		if number%2 == 0 {
			return sum + number
		}
		return sum
	}, 0)

}

func fib_numbers(n int) []int {
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

func GetLargestFactor(n int) int {
	fs := inthelper.Factor(n)
	return fs[len(fs)-1]
}

func GetLargestPalindrome() int {
	res := 0
	for n1 := 1; n1 < 1000; n1++ {
		for n2 := 1; n2 < 1000; n2++ {
			prd := n1 * n2
			if IsPalindrome(strconv.Itoa(prd)) {
				if prd > res {
					res = prd
				}
			}
		}
	}
	return res
}

func IsPalindrome(s string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(normalized)/2; i++ {
		if normalized[i] != normalized[len(normalized)-1-i] {
			return false
		}
	}
	return true
}
