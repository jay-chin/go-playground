package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	max_i := 6
	max_j := 6
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error")
	}

	matrix := make([][]int, max_j)
	lines := strings.Split(string(content), "\n")
	for i := 0; i < max_j; i++ {
		nums := strings.Split(lines[i], " ")
		matrix[i] = make([]int, max_i)
		for j := range nums {
			//fmt.Printf("i = %d, j = %d\n", i, j)
			matrix[i][j], err = strconv.Atoi(nums[j])
		}
	}
	fmt.Printf("%d\n", matrix)
	largest := math.MinInt64
	for j := 0; j < max_j-2; j++ {
		for i := 0; i < max_i-2; i++ {
			s := sumHG(matrix, i, j)
			if s > largest {
				largest = s
			}
		}
	}
	fmt.Printf("largest = %d\n", largest)
}

func sumHG(matrix [][]int, x int, y int) int {
	s := 0
	for i := x; i < x+3; i++ {
		//fmt.Printf("matrix[%d][%d] = %d\n", i, y, matrix[y][i])
		s += matrix[y][i]
		//fmt.Printf("matrix = %d\n", matrix[i][y+2])
		s += matrix[y+2][i]
	}
	s += matrix[y+1][x+1]
	//fmt.Printf("Sum = %d", s)
	return s
}
