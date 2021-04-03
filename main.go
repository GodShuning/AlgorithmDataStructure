package main

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	var res = make([]int, len(triangle[n-1]))
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == n-1 {
				res[j] = triangle[i][j]
			} else {
				res[j] = int(math.Min(float64(res[j]), float64(res[j+1]))) + triangle[i][j]
			}
		}
	}
	return res[0]
}
