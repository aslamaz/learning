package main

import "fmt"

func main() {
	var rows int
	var columns int
	var matrix [100][100]int
	fmt.Println("enter the numbers of columns")
	fmt.Scanf("%d", &columns)
	fmt.Println("enter the number of rows")
	fmt.Scanf("%d", &rows)
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			fmt.Scanf("%d", &matrix[r][c])
		}
	}
	for r := 1; r < rows-1; r++ {
		for c := 1; c < columns-1; c++ {
			if matrix[r-1][c] == 1 && matrix[r+1][c] == 1 && matrix[r+1][c] == 1 && matrix[r+1][c] == 1 {
				matrix[r][c] = 1
			}
		}
	}
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			fmt.Printf("%02d ", matrix[row][column])
		}
		fmt.Println()
	}
}

func matrix() {
	var rows int
	var columns int
	var matrix [100][100]int
	fmt.Println("enetr number of columns")
	fmt.Scanf("%d", &columns)
	fmt.Println("enter number of rows")
	fmt.Scanf("%d", &rows)
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			fmt.Scanf("%d", &matrix[row][column])
		}
	}
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			fmt.Printf("%02d ", matrix[row][column])
		}
		fmt.Println()
	}
}

func sumOfNaturalNumbersDivisibleBy3Or5(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum

}
func squareDiff(num int) int {
	var sum int
	var sumOfSquares int
	for i := 1; i <= num; i++ {
		sum = sum + i
		sumOfSquares += i * i
	}
	squareOfSum := sum * sum

	return sumOfSquares - squareOfSum
}
