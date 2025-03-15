package funcs

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
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
			if prd == inthelper.Reverse(prd) {
				return prd
			}
		}
	}
	return -1
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
	t4 := IsMultiPolygonalNumer(d)
	t5 := IsMultiPolygonalNumer(e)
	t6 := IsMultiPolygonalNumer(f)

	if reflect.DeepEqual(t1, t2) &&
		len(misc.Filter(t1, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t1, t3) &&
		len(misc.Filter(t1, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t1, t4) &&
		len(misc.Filter(t1, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t1, t5) &&
		len(misc.Filter(t1, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t1, t6) &&
		len(misc.Filter(t1, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t2, t3) &&
		len(misc.Filter(t2, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t2, t4) &&
		len(misc.Filter(t2, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t2, t5) &&
		len(misc.Filter(t2, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t2, t6) &&
		len(misc.Filter(t2, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t3, t4) &&
		len(misc.Filter(t3, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t3, t5) &&
		len(misc.Filter(t3, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t3, t6) &&
		len(misc.Filter(t3, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t4, t5) &&
		len(misc.Filter(t4, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t4, t6) &&
		len(misc.Filter(t4, func(b bool) bool { return b })) == 1 {
		return false
	}
	if reflect.DeepEqual(t5, t6) &&
		len(misc.Filter(t5, func(b bool) bool { return b })) == 1 {
		return false
	}

	rows := make([]string, 0)
	rows = append(rows, boolsToString(t1))
	rows = append(rows, boolsToString(t2))
	rows = append(rows, boolsToString(t3))
	rows = append(rows, boolsToString(t4))
	rows = append(rows, boolsToString(t5))
	rows = append(rows, boolsToString(t6))

	//fmt.Println(rows)
	path := findPath(rows)

	//fmt.Println(path)
	return len(path) > 0
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

func findPath(matrix []string) []int {
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

	if n < 1000 {
		return 0, 0
	}
	//println((n%100)*100 + 10)
	//println((n%100)*100 + 99)

	return (n%100)*100 + 10, (n%100)*100 + 99
}

func isTri(n int) bool {
	// Coefficients for the quadratic equation n^2 + n - 2i = 0
	a := 1
	b := 1
	c := -2 * n
	res, _ := hasPositiveIntegerSolution(a, b, c)
	return res
}

func isSquare(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func isPent(n int) bool {
	// Coefficients for the quadratic equation 3n^2 - n - 2i = 0
	a := 3
	b := -1
	c := -2 * n
	res, _ := hasPositiveIntegerSolution(a, b, c)
	return res
}

func isHexa(n int) bool {
	// Coefficients for the quadratic equation 2n^2 - n - i = 0
	a := 2
	b := -1
	c := -n
	res, _ := hasPositiveIntegerSolution(a, b, c)
	return res
}

func isHept(n int) bool {
	// Coefficients for the quadratic equation 5n^2 - 3n - 2i = 0
	a := 5
	b := -3
	c := -2 * n
	res, _ := hasPositiveIntegerSolution(a, b, c)
	return res
}

func isOcta(n int) bool {
	// Coefficients for the quadratic equation 3n^2 - 2n - i = 0
	a := 3
	b := -2
	c := -n
	res, _ := hasPositiveIntegerSolution(a, b, c)
	return res
}

// checks if a*x^2 + b*x + c = 0 has a positive integer solution
func hasPositiveIntegerSolution(a, b, c int) (bool, int) {
	if a == 0 {
		return false, 0 // Not a quadratic equation if a == 0
	}

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return false, 0 // No real roots
	}

	// Compute the roots
	sqrtDiscriminant := math.Sqrt(float64(discriminant))
	if sqrtDiscriminant != float64(int(sqrtDiscriminant)) {
		return false, 0 // Discriminant is not a perfect square, no integer roots
	}

	// Calculate both roots
	root1 := (-b + int(sqrtDiscriminant)) / (2 * a)
	root2 := (-b - int(sqrtDiscriminant)) / (2 * a)

	// Check for positive integer solutions
	if root1 > 0 {
		return true, root1
	}
	if root2 > 0 {
		return true, root2
	}

	return false, 0 // No positive integer solutions
}

func BruteForce61() int {
	for i := 1010; i < 10000; i++ {
		if !IsMultiPolygonalNumer2(i) {
			continue
		}
		a, b := GetLargestAndSmalles(i)
		for j := a; j <= b; j++ {
			if i == j {
				continue
			}
			if !IsMultiPolygonalNumer2(j) {
				continue
			}
			c, d := GetLargestAndSmalles(j)
			for k := c; k <= d; k++ {
				if j == k || i == k {
					continue
				}
				if !IsMultiPolygonalNumer2(k) {
					continue
				}
				e, f := GetLargestAndSmalles(k)
				for l := e; l <= f; l++ {
					if l == k || l == j || l == i {
						continue
					}
					if !IsMultiPolygonalNumer2(l) {
						continue
					}
					g, h := GetLargestAndSmalles(l)
					for m := g; m <= h; m++ {
						if m == l || m == k || m == j || m == i {
							continue
						}
						if !IsMultiPolygonalNumer2(m) {
							continue
						}
						x, y := GetLargestAndSmalles(m)
						for n := x; n <= y; n++ {
							if m == n || n == k || n == j || n == i || n == l {
								continue
							}
							if !CompareDigits(n, i) {
								continue
							}
							if IsMultiPolygonalNumer2(n) {
								//								fmt.Printf("%d, %d, %d, %d, %d, %d\n", i, j, k, l, m, n)
								if IsMultiPolygonalNumers(i, j, k, l, m, n) {
									//println(i, j, k, l, m, n)
									//println(i+j+k+l+m+n)
									return i + j + k + l + m + n
								}
							}
						}

					}

				}

			}
		}
	}
	panic("not found")
}

func IsPerfectPower_copied(n int, power int) (bool, int) {
	if n < 0 && power%2 == 0 {
		return false, 0 // Negative numbers cannot have real even roots
	}
	root := int(math.Round(math.Pow(float64(n), 1.0/float64(power))))
	// Verify that root^power equals n
	if int(math.Pow(float64(root), float64(power))) == n {
		return true, root
	}
	return false, 0
}

// get all permutations of a number, ignore repeated digits in input number
func GetUniquePermutations_cp(n int) []int {
	str := strconv.Itoa(n)
	digits := []rune(str)

	// Use a map to prevent duplicates
	permutationsSet := make(map[string]struct{})
	permute_unique_cp(digits, 0, permutationsSet)

	// Convert the map keys to integers
	result := []int{}
	for perm := range permutationsSet {
		num, _ := strconv.Atoi(perm)
		if len(str) == len(strconv.Itoa(num)) {
			result = append(result, num)
		}

	}
	return result
}

func GetQubePermutations(n int) []int {
	str := strconv.Itoa(n)
	digits := []rune(str)

	permutationsSet := make(map[string]struct{})
	permute_unique_cp(digits, 0, permutationsSet)

	result := []int{}
	for perm := range permutationsSet {
		num, _ := strconv.Atoi(perm)
		if len(str) == len(strconv.Itoa(num)) {
			result = append(result, num)
		}

	}
	return result
}

func permute_unique_cp(digits []rune, start int, permutationsSet map[string]struct{}) {
	if start == len(digits)-1 {
		permutationsSet[string(digits)] = struct{}{} // Add to map to avoid duplicates
		return
	}

	for i := start; i < len(digits); i++ {
		// Swap current element with the starting element
		digits[start], digits[i] = digits[i], digits[start]

		// Recursively generate permutations for the remaining digits
		permute_unique_cp(digits, start+1, permutationsSet)

		// Swap back to restore the original state
		digits[start], digits[i] = digits[i], digits[start]
	}
}

func GetPermutations_copied(n int) []int {
	str := strconv.Itoa(n)
	digits := []rune(str)

	permutations := []string{}
	permute_copied(digits, 0, &permutations)

	result := []int{}
	for _, perm := range permutations {
		num, _ := strconv.Atoi(perm)
		result = append(result, num)
	}
	return result
}

func permute_copied(digits []rune, start int, permutations *[]string) {
	if start == len(digits)-1 {
		*permutations = append(*permutations, string(digits))
		return
	}

	for i := start; i < len(digits); i++ {
		// Swap current element with the starting element
		digits[start], digits[i] = digits[i], digits[start]

		// Recursively generate permutations for the remaining digits
		permute_copied(digits, start+1, permutations)

		// Swap back to restore the original state
		digits[start], digits[i] = digits[i], digits[start]
	}
}

func Permutations_copied(arr []any) [][]any {
	var result [][]any
	generatePermutations_copied(arr, 0, &result)
	return result
}

func generatePermutations_copied(arr []any, start int, result *[][]any) {
	if start == len(arr)-1 {
		// Append a copy of the current permutation to the result
		temp := make([]any, len(arr))
		copy(temp, arr)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(arr); i++ {
		// Swap current element with the starting element
		arr[start], arr[i] = arr[i], arr[start]

		// Recursively generate permutations for the remaining elements
		generatePermutations_copied(arr, start+1, result)

		// Swap back to restore the original state
		arr[start], arr[i] = arr[i], arr[start]
	}
}

func No_of_cubes(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if ok, _ := IsPerfectPower_copied(arr[i], 3); ok {
			res++
		}

	}
	return res
}

func GenerateCubes(max int) []int {
	cubes := []int{}
	for i := 1; i <= max; i++ {
		cube := i * i * i
		cubes = append(cubes, cube)
	}
	return cubes
}

// Function to group numbers by their length
func GroupByLength(cubes []int) map[int][]int {
	groups := make(map[int][]int)
	for _, cube := range cubes {
		length := len(strconv.Itoa(cube)) // Get the number of digits
		groups[length] = append(groups[length], cube)
	}
	return groups
}

// Helper function to sort digits of a number
func sortDigits(n int) string {
	str := strconv.Itoa(n)
	digits := []rune(str)
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
	return string(digits)
}

// Function to find permutations within a group
func FindPermutations(numbers []int) map[string][]int {
	permutations := make(map[string][]int)
	for _, num := range numbers {
		sorted := sortDigits(num) // Sort the digits to identify permutations
		if len(sorted) == len(strconv.Itoa(num)) {
			permutations[sorted] = append(permutations[sorted], num)
		}
	}
	return permutations
}

func FindSmallest(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty") // Handle empty slice case
	}

	smallest := nums[0]
	for _, num := range nums[1:] {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

// AACCGGTT

//Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
//Output: 2

func MinMutation(startGene string, endGene string, bank []string) int {
	//steps := findSteps(startGene, endGene, bank, 0)
	steps := bfsShortestPath(startGene, endGene, bank)
	return steps

}

func getCandidates(input string, bank []string) []string {
	candidates := []string{}

	for _, word := range bank {
		if len(word) != len(input) {
			continue
		}

		diffCount := 0
		for i := range input {
			if input[i] != word[i] {
				diffCount++
			}
			if diffCount > 1 {
				break
			}
		}

		if diffCount == 1 {
			candidates = append(candidates, word)
		}
	}

	return candidates
}

func bfsShortestPath(startGene string, endGene string, bank []string) int {
	// Use a queue for BFS
	queue := &Queue{}
	queue.Enqueue(startGene)

	// Track visited nodes to avoid revisiting
	visited := map[string]bool{}
	visited[startGene] = true

	// Count steps (levels in BFS)
	steps := 0

	for queue.Len() > 0 {
		size := queue.Len()

		// Process all nodes at the current level
		for i := 0; i < size; i++ {
			current, _ := queue.Dequeue()

			// If we reach the endGene, return the number of steps
			if current == endGene {
				return steps
			}

			// Get all neighbors (candidates differing by 1 character)
			candidates := getCandidates(current, bank)
			for _, candidate := range candidates {
				if !visited[candidate] {
					visited[candidate] = true
					queue.Enqueue(candidate)
				}
			}
		}

		// Increment the step count after processing the current level
		steps++
	}

	// If we exhaust the queue without finding the endGene, return -1
	return -1
}

type Queue struct {
	elements []string
}

func (q *Queue) Enqueue(value string) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.elements) == 0 {
		return "", false
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Len() int {
	return len(q.elements)
}

func MaxArea(height []int) int {
	best := 0
	for i := 0; i < len(height); i++ {
		if getArea(height[i], height[i], len(height)-i) > best {
			for j := i + 1; j < len(height); j++ {
				fmt.Printf("i: %d, j: %d\n", height[i], height[j])
				area := getArea(height[i], height[j], j-i)
				if area > best {
					best = area
				}
			}
		}
	}
	return best
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getArea(h1, h2, dist int) int {
	h := getMin(h1, h2)
	return h * dist
}

func LongestCommonPrefix(strs []string) string {
	res := ""

	return res
}
