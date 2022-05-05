package sort

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(left []int, right []int) []int {
	var i, j int = 0, 0
	var size int = len(left) + len(right)
	array := make([]int, size, size)
	for k:=0; k < size; k++ {
		if i < len(left) && j<len(right){
			if left[i] < right[j] {
				array[k] = left[i]
				i++
			} else {
				array[k] = right[j]
				j++
			}
		} else if i < len(left) {
			array[k] = left[i]
			i++
		} else if j < len(right) {
			array[k] = right[j]
			j++
		}
	}
	return array
}

func partition(arr []int, low, high int) int {
	pivot := (low + high) / 2
	i := low - 1
	arr[pivot], arr[high] = arr[high], arr[pivot]
	for j:=low; j<high; j++ {
		if arr[j] <= arr[high] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i+1
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func Sort(arr []int, algo string) []int {
	if algo == "merge" {
		arr = mergeSort(arr)
	} else if algo == "quick" {
		quickSort(arr, 0, len(arr)-1)
	}
	return arr
}