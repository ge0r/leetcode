package main

/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	reshape := make([][]int, r)
	reshape[0] = make([]int, c)

	k, l := 0, 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			reshape[k][l] = mat[i][j]
			l++
			if l == c {
				l = 0
				k++
				if k == r {
					return reshape
				}
				reshape[k] = make([]int, c)
			}
		}

	}
	return reshape
}

// @lc code=end
