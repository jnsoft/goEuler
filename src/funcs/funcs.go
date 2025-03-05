package funcs

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/jnsoft/jngo/inthelper"
	"github.com/jnsoft/jngo/misc"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
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

func Factor(n int) (int, int, error) {
	if inthelper.MillerRabin(big.NewInt(int64(n)), 13) {
		return n, 1, nil
	}
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return i, n / i, nil
		}
	}
	return -1, -1, errors.New("no factor found")
}

func FullFactorization(n int) ([]int, error) {
	factors := []int{}
	var err error
	var factor int

	for n > 1 {
		factor, n, err = Factor(n)
		if err != nil {
			return nil, err
		}
		factors = append(factors, factor)
	}

	return factors, nil
}
