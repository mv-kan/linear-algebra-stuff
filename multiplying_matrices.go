package main

import "fmt"

// source
// https://www.mathsisfun.com/algebra/matrix-multiplying.html
/*
Notes Identity of Matrix

I =
[
   1 0 0
   0 1 0
   0 0 1
]
I - is identity matrix (it is like 1 for scalar)
	1. Identity matrix is "square"
	2. Can be as large as we want it to be (1x1, 10x10, 100x100...)
	3. has 1's on main diagonal
	4. It's symbol is I (capital I)

it is special matrix because...
A * I = A (dot product)
I * A = A

DOT PRODUCT
A(BC)=(AB)C
AB != BA
*/

func ScalarMultiplication(scalar Scalar, matrix Matrix) Matrix {
	copyMatrix := CopyMatrix(matrix)
	for i := 0; i < len(copyMatrix); i++ {
		for j := 0; j < len(copyMatrix); j++ {
			copyMatrix[i][j] = int(scalar) * matrix[i][j]
		}
	}
	return copyMatrix
}

func DotProduct(m1 Matrix, m2 Matrix) Matrix {
	// pre
	if len(m1[0]) != len(m2) {
		panic("len(m1[0]) != len(m2)")
	}
	// init
	// To multiply an m×n matrix by an n×p matrix, the ns must be the same,
	// and the result is an m×p matrix.
	product := make([][]int, len(m1))
	for i := 0; i < len(m1); i++ {
		product[i] = make([]int, len(m2[0]))
	}
	// calculate
	for i1 := 0; i1 < len(m1); i1++ {
		for i2 := 0; i2 < len(m2[0]); i2++ {
			for j := 0; j < len(m2); j++ {
				product[i1][i2] += m1[i1][j] * m2[j][i2]
			}
		}
	}
	return product
}

func RunMultiplyingMatrices() {
	var (
		s Scalar = 2
		m Matrix = [][]int{
			{4, 0},
			{1, -9},
		}
	)
	fmt.Println("Scalar multiplication:")
	fmt.Printf("%d * %v = %v\n", s, m, ScalarMultiplication(s, m))

	var (
		m1 Matrix = [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}
		m2 Matrix = [][]int{
			{7, 8},
			{9, 10},
			{11, 12},
		}
	)
	fmt.Println("Dot product:")
	fmt.Printf("%v * %v = %v\n", m1, m2, DotProduct(m1, m2))
	m1 = [][]int{
		{3, 4, 2},
	}
	m2 = [][]int{
		{13, 9, 7, 15},
		{8, 7, 4, 6},
		{6, 4, 0, 3},
	}
	fmt.Println("Dot product 2:")
	fmt.Printf("%v * %v = %v\n", m1, m2, DotProduct(m1, m2))
	m1 = [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("Dot product (identity matrix): ")
	fmt.Printf("%v * %v = %v\n", m1, m2, DotProduct(m1, m2))
}
