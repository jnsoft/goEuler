package main

import (
	"fmt"

	. "github.com/jnsoft/goEuler/src/funcs"
	"github.com/jnsoft/goEuler/src/geohelper"
	"github.com/jnsoft/jngo/bag"
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

	p1 := geohelper.Point{X: -340, Y: 495}
	p2 := geohelper.Point{-153, -910}
	p3 := geohelper.Point{X: 835, Y: -947}

	// Print the Points
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	points := []geohelper.Point{
		{X: -340, Y: 495},
		{X: -153, Y: -910},
		{X: 835, Y: -947},
	}

	minX, maxX, minY, maxY := findGridBoundaries(points)
	grid := createGrid(minX, maxX, minY, maxY)
	placePointsOnGrid(grid, points, minX, maxX, minY, maxY)
	printGrid(grid)

	//res := IsPerfectDivisible(20)
	//fmt.Printf("Res %v", res)
}
