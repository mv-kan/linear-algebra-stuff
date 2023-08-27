package main

import "fmt"

func TransposeMatrix(m Matrix) Matrix {
	tm := make([][]int, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		tm[i] = make([]int, len(m))
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			tm[j][i] = m[i][j]
		}
	}
	return tm
}

func RunTranspose() {
	fmt.Println("transpose:")
	var m1 Matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("%v -> %v\n", m1, TransposeMatrix(m1))

	fmt.Println("transpose2:")
	m1 = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Printf("%v -> %v\n", m1, TransposeMatrix(m1))
}
