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

	strs := []string{"flower", "flow", "flight"}
	TimeFunction("Answer", func() (interface{}, error) {
		return LongestCommonPrefix(strs), nil
	})

	return

	best := 0
	for q_root := 10; q_root < 200; q_root++ {
		qube := q_root * q_root * q_root
		ns := GetUniquePermutations_cp(qube)
		no_of_qubes := No_of_cubes(ns)
		if no_of_qubes > best {
			best = no_of_qubes
			//fmt.Printf("Number: %d: %d cubes\n", q_root, no_of_qubes)
		}
		if q_root%10 == 0 {
			//println(q_root)
		}
	}

	res := 1000000

	// Step 1: Generate lots of cube numbers
	maxCubeRoot := 50000 // Adjust this for more or fewer cubes
	cubes := GenerateCubes(maxCubeRoot)

	// Step 2: Group cube numbers by their length
	groups := GroupByLength(cubes)

	// Step 3: Process each group to find permutations
	for _, numbers := range groups {
		permutations := FindPermutations(numbers)
		for key, nums := range permutations {
			if len(nums) == 4 && key[0] != '0' { // Only print groups with actual permutations
				fmt.Printf("  Permutations of %s: %v\n", key, nums)
				s := FindSmallest(nums)
				test, s_qr := IsPerfectPower_copied(s, 3)
				if test {
					if s_qr < res {
						res = s_qr
					}
					fmt.Printf("%d\n", s_qr)
				} else {
					println("wtf")
				}
			}
		}
	}
	println(res)
	println(res * res * res)

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
