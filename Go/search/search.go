package search

// binary search iterative approach
func binarySearchIterative(arr []int, target int, low int, high int) int{
	for low <= high {
		mid := (low + high) / 2
		if target >= arr[mid] {
			low = mid+1
		} else {
			high = mid-1
		}
	}
	if arr[high] == target {
		return high
	} else {
		return high + 1
	}
}

// binary search recursive approach
func binarySearchRecursive(arr []int, target int, low int, high int) int {
	mid := low + (high - low) / 2
	if low >= high {
		if arr[mid] == target {
			return mid
		} else if target > arr[mid]{
			return mid + 1
		} else {
			return mid - 1
		}
	}
	if target < arr[mid] {
		return binarySearchRecursive(arr, target, low, mid)
	} else {
		return binarySearchRecursive(arr, target, mid+1, high)
	}
}

// An exposed function to call search algorithm
func Search(arr []int, target int, use_recursion bool = false) int {
	low, high := 0, len(arr)-1
	if use_recursion {
		return binarySearchRecursive(arr, target, low, high)
	} else {
		return binarySearchIterative(arr, target, low, high)
	}
}