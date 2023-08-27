package main

type Scalar int
type Vector []int
type Matrix [][]int

func CopyVector(v Vector) Vector {
	copyV := make([]int, len(v))
	copy(copyV, v)
	return copyV
}

func CopyMatrix(m Matrix) Matrix {
	copyM := make([][]int, len(m))
	for i := 0; i < len(copyM); i++ {
		copyM[i] = CopyVector(m[i])
	}
	return copyM
}

func main() {
	// RunMultiplyingMatrices()
	RunTranspose()
}
