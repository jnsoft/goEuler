package funcs

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jnsoft/jngo/geohelper"
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
	for n1 := 999; n1 > 99; n1-- {
		for n2 := n1; n2 > 99; n2-- {
			prd := n1 * n2
			if prd == reverse(prd) {
				return prd
			}
		}
	}
	return -1
}

func reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = 10*rev + n%10
		n = n / 10
	}
	return rev
}

func IsPerfectDivisible(no_of_divisors int) int {
	limit := int(math.Sqrt(float64(no_of_divisors)))
	primes := inthelper.PrimesSieve(no_of_divisors)
	a := make([]int, len(primes)) // no of factors if static: var a [10]int
	N := 1
	i := 0
	check := true

	for len(primes) > i {
		a[i] = 1
		if check {
			if primes[i] <= limit {
				a[i] = int(math.Floor(math.Log(float64(no_of_divisors)) / math.Log(float64(primes[i]))))
			} else {
				check = false
			}
		}
		N *= inthelper.Pow(primes[i], a[i])
		i++
	}

	return N
}

func SumSquareDiff(n int) int {
	ns := rangeSlice(1, n)
	sum := misc.Fold(ns, func(acc, num int) int {
		return acc + num
	}, 0)
	sumSq := inthelper.Pow(sum, 2)
	sqSum := misc.Fold(ns, func(acc, num int) int {
		return acc + inthelper.Pow(num, 2)
	}, 0)
	return sumSq - sqSum
}

func rangeSlice(a, b int) []int {
	size := b - a + 1
	slice := make([]int, size)
	for i := range slice {
		slice[i] = a + i
	}
	return slice
}

func ReadPoints(fileName string) [][]geohelper.Point {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var triangles [][]geohelper.Point

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		var points []geohelper.Point
		for i := 0; i < len(nums); i += 2 {
			x, _ := strconv.ParseFloat(nums[i], 64)
			y, _ := strconv.ParseFloat(nums[i+1], 64)
			points = append(points, geohelper.Point{x, y})
		}
		triangles = append(triangles, points)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return triangles
}

func IsMultiPolygonalNumer(a, b, c, d, e, f int) bool {
	ns := make([]int, 4)
	ns = append(ns, FindPolygonalNumberIx(a, TriangularNumber))
	ns = append(ns, FindPolygonalNumberIx(b, SquareNumber))
	ns = append(ns, FindPolygonalNumberIx(c, PentagonalNumber))
	//ns = append(ns, FindPolygonalNumberIx(d, HexagonalNumber))
	//ns = append(ns, FindPolygonalNumberIx(e, HeptagonalNumber))
	//ns = append(ns, FindPolygonalNumberIx(f, OctagonalNumber))

	if misc.Fold(ns, func(a, b int) int { return a * b }, 1) != 0 {
		if !misc.HasDuplicates(ns) {
			return true
		}
	}

	return false
}

type PolygonalNumberFunc func(int) int

func FindPolygonalNumberIx(n int, fn PolygonalNumberFunc) int {
	ix := 1
	val := 0
	for val < n {
		val = fn(ix)
		ix++
	}
	if val == n {
		return ix
	} else {
		return 0
	}
}

func TriangularNumber(n int) int {
	return n * (n + 1) / 2
}

func SquareNumber(n int) int {
	return n * n
}

func PentagonalNumber(n int) int {
	return n * (3*n - 1) / 2
}

func HexagonalNumber(n int) int {
	return n * (2*n - 1)
}

func HeptagonalNumber(n int) int {
	return n * (5*n - 3) / 2
}

func OctagonalNumber(n int) int {
	return n * (3*n - 2)
}
