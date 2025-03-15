package main

import (
	"fmt"

	"github.com/jnsoft/goEuler/src/funcs"
	. "github.com/jnsoft/goEuler/src/funcs"
	"github.com/jnsoft/jngo/bag"
	"github.com/jnsoft/jngo/geohelper"
)

func main() {
	// Get a greeting message and print it.
	message := Hello("Johan")
	fmt.Println(message)

	b := bag.NewBag[int]()

	b.Add(4)
	b.Add(3)

	TimeFunction("Answer 1", func() (interface{}, error) {
		return Multiples_of_3_and_5(1000), nil
	})

	TimeFunction("Answer 2", func() (interface{}, error) {
		return Sum_even_valued_fibs(4000000), nil
	})

	TimeFunction("Answer 3", func() (interface{}, error) {
		return GetLargestFactor(600851475143), nil
	})

	TimeFunction("Answer 4", func() (interface{}, error) {
		return GetLargestPalindrome(), nil
	})

	TimeFunction("Answer 5", func() (interface{}, error) {
		return IsPerfectDivisible(20), nil
	})

	TimeFunction("Answer 6", func() (interface{}, error) {
		return SumSquareDiff(100), nil
	})

	triangles := ReadPoints("data/prob102.txt")
	TimeFunction("Answer 102", func() (interface{}, error) {
		c := 0
		for _, val := range triangles {
			if geohelper.ContainsOrigo(val) {
				c++
			}
		}
		return c, nil
	})

	//ns := []int{8128, 2882, 8281}

	funcs.IsMultiPolygonalNumers(8128, 2882, 8281, 0, 0, 0)
	// secondLastDigit := (n / 10) % 10 // != 0

	for i := 1010; i < 10000; i++ {
		if !IsMultiPolygonalNumer2(i) {
			continue
		}
		println(i)
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
								if IsMultiPolygonalNumers(i, j, k, l, m, n) {
									println(i, j, k, l, m, n)
								}
								//println(i, j, k, l, m, n)
							}

						}

					}

				}

			}
		}
	}

}
