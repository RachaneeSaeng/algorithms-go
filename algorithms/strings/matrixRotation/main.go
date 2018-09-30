package main

// Rotate square matrix
// degree is product of 90 e.g. 90, 180, 270, 360
func rotateMatrix(matrix [][]int32, degree int32) [][]int32 {
	rotateTime := (degree / 90) % 4

	if len(matrix) == 0 || len(matrix) != len(matrix[0]) || rotateTime == 0 {
		return matrix // cannot rotate matrix that is not a square metrix
	}

	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		firstIdx := layer
		lastIdx := n - 1 - layer
		for i, j := firstIdx, 0; i < lastIdx; i, j = i+1, j+1 { // move n-1 times for each round (corner items will be moved as the first item of next round)
			/* 0: top, 1: left, 2: bottom, 3: right
			top = (rotateTime) % 4
			left = (1 + rotateTime) % 4
			bottom = (2 + rotateTime) %4
			right = (3 + rotateTime) %4
			*/

			// store temp data
			var allSideValue [4]int32
			allSideValue[0] = matrix[firstIdx][i]         //top
			allSideValue[1] = matrix[lastIdx-j][firstIdx] //left
			allSideValue[2] = matrix[lastIdx][lastIdx-j]  //bottom
			allSideValue[3] = matrix[i][lastIdx]          //right

			// top = (rotateTime) % 4
			matrix[firstIdx][i] = allSideValue[rotateTime%4]
			// left = (1 + rotateTime) % 4
			matrix[lastIdx-j][firstIdx] = allSideValue[(1+rotateTime)%4]
			// bottom = (2 + rotateTime) %4
			matrix[lastIdx][lastIdx-j] = allSideValue[(2+rotateTime)%4]
			// right = (3 + rotateTime) %4
			matrix[i][lastIdx] = allSideValue[(3+rotateTime)%4]
		}
	}
	return matrix
}

func main() {

}
