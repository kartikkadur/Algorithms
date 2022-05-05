package twopointers

// find two numbers that add up to target
func TwoSum(arr []int, target int) []int {
    mapping := make(map[int]int)
    for i:=0; i<len(arr); i++ {
        _, first_number := mapping[arr[i]]
        if !first_number {
            mapping[arr[i]] = i
        } 
        index, second_number := mapping[target - arr[i]]
        if second_number {
            return []int {i, index}
        }
    }
    return []int {-1, -1}
}

/* Given hight of the walls as array values find the max
 volume of water it can hold. Function takes in an array
 height which has integers representing the wall heights
 and returns the max volume and the array index of the walls*/
func MaxVolume(height []int) (int, [] int) {
    low, high := 0, len(height) - 1
    volume := 0
    current_volume := 0
    for low < high {
        if height[low] < height[high] {
            current_volume = (high - low) * height[low]
            low++
        } else {
            current_volume = (high - low) * height[high]
            high--
        }
        if volume < current_volume {
            volume = current_volume
        }
    }
    return volume, []int {low-1, high}
}

/* Remove duplicates */
func RemoveDuplicates(arr []int) int {
    insert := 0
    for j:= 1; j<len(arr); j++ {
        if arr[j] != arr[j-1] {
            insert++
            arr[insert] = arr[j]
        }
    }
    return insert+1
}

/* Remove given value */
func RemoveValue(arr []int, val int) int {
    insert := 0
    for j:=0; j<len(arr); j++ {
        if arr[j] != val {
            arr[insert] = arr[j]
            insert++
        }
    }
    return insert
}

/* find substring in a string and return its starting index */
func FindNeedle(haystack string, needle string) int {
    for i:=0; i<len(haystack) - len(needle) + 1; i++ {
        if haystack[i:i+len(needle)] == needle {
            return i
        }
    }
    return -1
}