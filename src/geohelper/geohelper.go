package geohelper

import (
	"fmt"
	"math"
	"strings"

	"github.com/jnsoft/jngo/geohelper"
)

func PrintPoints(points []geohelper.Point, width, height int) string {
	scaledPoints, minX, maxX, minY, maxY := scalePoints(points, width, height)
	gridStr := generateGrid(scaledPoints, minX, maxX, minY, maxY, width, height)
	return gridStr
}

func scalePoints(points []geohelper.Point, width, height int) ([]geohelper.Point, float64, float64, float64, float64) {
	if len(points) == 0 {
		return points, 0, 0, 0, 0
	}

	// Find min and max values for X and Y
	minX, maxX := points[0].X, points[0].X
	minY, maxY := points[0].Y, points[0].Y

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

	// Avoid division by zero if all points are the same
	rangeX := maxX - minX
	rangeY := maxY - minY
	if rangeX == 0 {
		rangeX = 1
	}
	if rangeY == 0 {
		rangeY = 1
	}

	// Scale points
	scaled := make([]geohelper.Point, len(points))
	for i, p := range points {
		scaledX := (p.X - minX) / rangeX * float64(width-1)
		scaledY := (p.Y - minY) / rangeY * float64(height-1)
		scaled[i] = geohelper.Point{X: scaledX, Y: scaledY}
	}

	return scaled, minX, maxX, minY, maxY
}

// generateGrid creates a scaled grid with visible X and Y axes
func generateGrid(points []geohelper.Point, minX, maxX, minY, maxY float64, width, height int) string {
	grid := make([][]rune, height+2) // Extra space for axes labels
	for i := range grid {
		grid[i] = make([]rune, width+4) // Extra space for axis labels
		for j := range grid[i] {
			grid[i][j] = ' ' // Default empty space
		}
	}

	// Determine axis positions
	yAxis := int((0 - minX) / (maxX - minX) * float64(width-1))
	xAxis := int((maxY - 0) / (maxY - minY) * float64(height-1))

	// Draw Y-axis (vertical)
	if yAxis >= 0 && yAxis < width {
		for i := 0; i < height; i++ {
			grid[i][yAxis+2] = '|'
		}
	}

	// Draw X-axis (horizontal)
	if xAxis >= 0 && xAxis < height {
		for j := 0; j < width; j++ {
			grid[xAxis][j+2] = '-'
		}
	}

	// Mark the origin (0,0)
	if xAxis >= 0 && xAxis < height && yAxis >= 0 && yAxis < width {
		grid[xAxis][yAxis+2] = '+'
	}

	// Plot points
	for _, p := range points {
		x := int(p.X)
		y := int(height - int(p.Y) - 1) // Flip Y-axis

		if x >= 0 && x < width && y >= 0 && y < height {
			grid[y][x+2] = '*' // Shift right to accommodate Y-axis numbers
		}
	}

	// Convert grid to string with axis labels
	var sb strings.Builder

	// Add Y-axis labels
	for i, row := range grid {
		if i < height {
			sb.WriteString(fmt.Sprintf("%2d ", int(maxY-float64(i)*(maxY-minY)/float64(height-1))))
		} else {
			sb.WriteString("   ") // Empty space for X-axis labels
		}
		sb.WriteString(string(row) + "\n")
	}

	// Add X-axis labels
	sb.WriteString("   ") // Space for Y-axis numbers
	sb.WriteString("+" + strings.Repeat("-", width) + "\n   ")
	for i := 0; i < width; i++ {
		sb.WriteString(fmt.Sprintf("%d", int(minX+float64(i)*(maxX-minX)/float64(width-1))))
		if i < width-1 {
			sb.WriteString(" ") // Space between numbers
		}
	}

	return sb.String()
}

func PrintPoints_old(points []geohelper.Point) string {
	minX, maxX, minY, maxY := findGridBoundaries(points)
	grid := createGrid(minX, maxX, minY, maxY)
	placePointsOnGrid(grid, points, minX, maxX, minY, maxY)
	return formatGrid(grid)
}

func findGridBoundaries(points []geohelper.Point) (minX, maxX, minY, maxY float64) {
	minX, maxX, minY, maxY = math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
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

func createGrid(minX, maxX, minY, maxY float64) [][]rune {
	gridWidth := int(maxX-minX) + 1
	gridHeight := int(maxY-minY) + 1

	grid := make([][]rune, gridHeight)
	for i := range grid {
		grid[i] = make([]rune, gridWidth)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func placePointsOnGrid(grid [][]rune, points []geohelper.Point, minX, maxX, minY, maxY float64) {
	for _, p := range points {
		x := int(p.X - minX)
		y := int(maxY - p.Y)
		if x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
			grid[y][x] = 'P'
		}
	}
}

func formatGrid(grid [][]rune) string {
	var result string
	for _, row := range grid {
		for _, cell := range row {
			result += fmt.Sprintf("%c ", cell)
		}
		result += "\n"
	}
	return result
}
