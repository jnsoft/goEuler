package geohelper

import (
    "fmt"
    "math"
)

// Define the Point struct
type Point struct {
    X, Y int
}

type Point struct { X, Y float64 }

// The winding number algorithm. 
// This algorithm counts how many times the polygon winds around the origin. 
// If the winding number is non-zero, the origin is inside the polygon.
func ContainsOrigo(points []Point) bool {
    windingNumber := 0

    for i := 0; i < len(points); i++ {
        next := (i + 1) % len(points)
        if points[i].Y <= 0 {
            if points[next].Y > 0 && isLeft(points[i], points[next], Point{0, 0}) > 0 {
                windingNumber++
            }
        } else {
            if points[next].Y <= 0 && isLeft(points[i], points[next], Point{0, 0}) < 0 {
                windingNumber--
            }
        }
    }

    return windingNumber != 0
}

func isLeft(p1, p2, p Point) float64 {
    return (p2.X-p1.X)*(p.Y-p1.Y) - (p.X-p1.X)*(p2.Y-p1.Y)
}


func findGridBoundaries(points []Point) (minX, maxX, minY, maxY int) {
    minX, maxX, minY, maxY = math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
    for _, p := range points {
        if p.X < minX {
            minX = p.X
        }
        if p.X > maxX {
            maxX = p.X
        }
        if p.Y < minY {
            minY = p.Y
        }
        if p.Y > maxY {
            maxY = p.Y
        }
    }
    return
}

func createGrid(minX, maxX, minY, maxY int) [][]rune {
    gridWidth := maxX - minX + 1
    gridHeight := maxY - minY + 1

    grid := make([][]rune, gridHeight)
    for i := range grid {
        grid[i] = make([]rune, gridWidth)
        for j := range grid[i] {
            grid[i][j] = '.'
        }
    }
    return grid
}

func placePointsOnGrid(grid [][]rune, points []Point, minX, maxX, minY, maxY int) {
    for _, p := range points {
        x := p.X - minX
        y := maxY - p.Y
        if x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
            grid[y][x] = 'P'
        }
    }
}

func printGrid(grid [][]rune) {
    for _, row := range grid {
        for _, cell := range row {
            fmt.Printf("%c ", cell)
        }
        fmt.Println()
    }