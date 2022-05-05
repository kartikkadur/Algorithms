package search

func binarySearchIterative(arr []int, target int) int{
	var low, high int = 0, len(arr)-1
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

func Search(arr []int, target int) int {
	//low, high := 0, len(arr)-1
	return binarySearchIterative(arr, target)
}