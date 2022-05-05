package recursion

// Sum all non negative integers up to n
func AddPositiveNums(n int) int {
	if n <= 0 {
		return 0
	}
	return n + AddPositiveNums(n-1)
}

// find unique paths in a n x m grid
func UniquePaths(n int, m int) int {
	if n == 1 || m == 1 {
		return 1
	}
	return UniquePaths(n, m-1) + UniquePaths(n-1, m)
}