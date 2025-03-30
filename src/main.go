package main

import (
	"fmt"

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

	TimeFunction("Answer 61", func() (interface{}, error) {
		return BruteForce61(), nil
	})

	TimeFunction("Answer 62", func() (interface{}, error) {
		return CP2(), nil
	})

	//////////////////////////////////////////////////////////////////////

	//arr := []int{2, 3, 1, 2, 4, 3}
	//arr2 := []int{2, 5, 6}

	TimeFunction("Answer", func() (interface{}, error) {

		//fmt.Printf("%v\n", r)
		return Test("aa", "aa"), nil
	})

	return

	//////////////////////////////////////////////////////////////////////

}

/*

	// n := 41063625
	//best_ress := make([]int, 0)
	best_c := 0
	best_q := 0
	for i := 2; i < 10000; i++ {
		test, _ := IsPerfectPower_copied(i, 3)
		if !test {
			continue
		}
		ns := GetUniquePermutations_cp(i)
		c := 0
		q := 0
		ress := make([]int, 0)
		for _, p := range ns {
			res, n := IsPerfectPower_copied(p, 3)
			if res {
				q = n
				c++
				ress = append(ress, p)
			}
		}
		if c > best_c {
			//best_ress = ress
			best_c = c
			best_q = q
			fmt.Printf("new best: %d, %d->%v\n", c,q, ress)
		}

	}
	fmt.Printf("ans: %d\n", best_q)

}

*/
