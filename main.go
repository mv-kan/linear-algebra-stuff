package main

type Scalar int

/*
VECTOR
in linear algebra we have vectors not like that
[1,2,3]
but like that
(vertically)

	1
	2
	3
*/
type Vector [][]int
type Matrix [][]int

func CopySlice(s []int) []int {
	copyS := make([]int, len(s))
	copy(copyS, s)
	return copyS
}

func CopyVector(v Vector) Vector {
	copyV := make([][]int, len(v))
	for i := 0; i < len(copyV); i++ {
		copyV[i] = CopySlice(v[i])
	}
	return copyV
}

func CopyMatrix(m Matrix) Matrix {
	copyM := make([][]int, len(m))
	for i := 0; i < len(copyM); i++ {
		copyM[i] = CopySlice(m[i])
	}
	return copyM
}

func main() {
	RunMultiplyingMatrices()
	RunTranspose()
}
