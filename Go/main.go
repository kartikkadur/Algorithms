package main

import "fmt"
import _ "algo/sort"
import _ "algo/search"
import _ "algo/recursion"
import _ "algo/twopointers"
import _ "algo/datastruct"
import "algo/graphs"

func main() {
	arr := [][]int{{1, 1, 0, 0, 0},
					{0, 1, 0, 0, 1},
					{1, 0, 0, 1, 1},
					{0, 0, 0, 0, 0},
					{1, 0, 1, 0, 1}}
	fmt.Println(graphs.FindConnectedComponents(arr))
}