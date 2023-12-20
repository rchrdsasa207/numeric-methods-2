package main

import "fmt"

func SampleMatrix() [][]float64 {
	return [][]float64{{2, 1, -1, 8}, {-3, -1, 2, -11}, {-2, 1, 2, -3}}
}

func printMatrix(m [][]float64) {
	for i := range m {
		fmt.Print("|")
		for j := range m[i] {
			fmt.Printf("%10v|", m[i][j])
		}
		fmt.Println()
	}
}

func Eliminate(matrix [][]float64) {
	for i := 0; i < 2; i++ {
		for j := i + 1; j < len(matrix); j++ {
			a := matrix[j][i]
			for k := 0; k < len(matrix)+1; k++ {
				matrix[j][k] -= a * matrix[i][k] / matrix[i][i]
			}
		}
	}
}

func solve(matrix [][]float64) []float64 {
	solution := make([]float64, len(matrix))
	for i := len(matrix) - 1; i >= 0; i-- {
		toSubtract := 0.0
		for j := len(matrix) - 1; j > i; j-- {
			toSubtract += matrix[i][j] * solution[j]
		}
		solution[i] = (matrix[i][len(matrix)] - toSubtract) / matrix[i][i]
	}
	return solution
}

func main() {
	matrix := SampleMatrix()
	Eliminate(matrix)
	printMatrix(matrix)
	a := solve(matrix)
	fmt.Println(a)
}
