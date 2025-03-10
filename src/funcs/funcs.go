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

func IsMultiPolygonalNumers(a, b, c, d, e, f int) bool {
	t1 := IsMultiPolygonalNumer(a)
	t2 := IsMultiPolygonalNumer(b)
	t3 := IsMultiPolygonalNumer(c)
	//t4 := IsMultiPolygonalNumer(d)
	//t5 := IsMultiPolygonalNumer(e)
	//t6 := IsMultiPolygonalNumer(f)

	rows := make([]string, 0)
	rows = append(rows, boolsToString(t1))
	rows = append(rows, boolsToString(t2))
	rows = append(rows, boolsToString(t3))
	//rows = append(rows, boolsToString(t4))
	//rows = append(rows, boolsToString(t5))
	//rows = append(rows, boolsToString(t6))

	path := findPath(rows)
	fmt.Println(rows)
	fmt.Println(path)
	return false
}

func IsMultiPolygonalNumer(n int) []bool {
	ns := make([]bool, 0)
	ns = append(ns, FindPolygonalNumberIx(n, TriangularNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, SquareNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, PentagonalNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, HexagonalNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, HeptagonalNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, OctagonalNumber) != 0)
	return ns
}

func IsMultiPolygonalNumer2(n int) bool {
	ns := make([]bool, 0)
	ns = append(ns, FindPolygonalNumberIx(n, TriangularNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, SquareNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, PentagonalNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, HexagonalNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, HeptagonalNumber) != 0)
	ns = append(ns, FindPolygonalNumberIx(n, OctagonalNumber) != 0)
	return misc.Reduce(ns, func(a int, b bool) int {
		if b {
			return a + 1
		}
		return a
	}, 0) > 0
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
		return ix - 1
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

func boolsToString(bools []bool) string {
	var builder strings.Builder

	for _, value := range bools {
		if value {
			builder.WriteString("1")
		} else {
			builder.WriteString("0")
		}
	}

	return builder.String()
}

func FindPath(matrix []string) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	path := make([]int, rows)
	visited := make([]bool, cols)

	for i := 0; i < rows; i++ {
		found := false
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' && !visited[j] {
				path[i] = j + 1 // Store the column index (1-based)
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return nil // No valid path found
		}
	}

	return path
}

// check four digit numbers if the last two digits of num1 matches the first two digits of num2
func CompareDigits(num1, num2 int) bool {
	// Extract the last two digits of num1
	lastTwoDigitsNum1 := num1 % 100

	// Extract the first two digits of num2
	firstTwoDigitsNum2 := num2 / 100

	return lastTwoDigitsNum1 == firstTwoDigitsNum2
}

func GetLargestAndSmalles(n int) (int, int) {

	return (n%100)*100 + 10, (n%100)*100 + 99
}
