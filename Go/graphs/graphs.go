package graphs

import "fmt"
import "algo/datastruct"

// Depth first graph traversal
func DepthFirstTraversal(graph [][]int, root int, visited []bool) {
	fmt.Println(root)
	visited[root] = true	// keep track of visited node to avoide infinite loops
	for i:=0; i<len(graph); i++ {
		if graph[root][i]==1 && !visited[i] {
			DepthFirstTraversal(graph, i, visited)	//utilize the function stack to get DFS traversal
		}
	}
}

// Depth first graph traversal iterative implementation
func DepthFirstSearchIter(graph [][]int, root int, visited []bool) {
	var stack datastruct.Stack	// initialize a stack to hold values
	stack.Push(root)
	visited[root] = true
	for (!stack.IsEmpty()) { // loop unitil stack is empty
		/*inside the loop pop the top of the stack and push all the
		adjacent nodes of the poped nodes into the stack. Simultaneously
		mark the node as visited
		*/
		current := stack.Pop()
		for i:=0; i<len(graph[current]); i++ {
			if graph[current][i] == 1 && current != i && !visited[i] {
				stack.Push(i)
				visited[i] = true
			}
		}
	}
}

// Berdth first traversal
func BredthFirstSeatchIter(graph [][]int, root int) {
	var queue datastruct.Queue // same as DFS but utilizes Queue for storing nodes
	visited := make([]bool, len(graph), len(graph)) 
	queue.Push(root)
	visited[root] = true
	for (!queue.IsEmpty()) {
		current := queue.Pop()
		fmt.Println(current)
		for i:=0; i<len(graph[current]); i++ {
			if graph[current][i] == 1 && current != i && !visited[i] {
				queue.Push(i)
				visited[i] = true
			}
		}
	}
}

// check if the path exists between two nodes
func HasPath(graph [][]int, src int, dst int) bool {
	var stack datastruct.Stack
	visited := make([]bool, len(graph), len(graph))
	if src == dst {
		return true
	}
	stack.Push(src)
	for !stack.IsEmpty() {
		current := stack.Pop()
		for i:=0; i<len(graph[current]); i++ {
			if graph[current][i] == 1 && current != i && !visited[i] {
				if i == dst {
					return true
				} else {
					stack.Push(i)
					visited[i] = true
				}
			}
		}
	}
	return false
}

func isValied(graph [][]int, row int, col int, visited [][]bool) bool {
	// checks if the index is valied and if the node is not visited.
	// Returns true if both the conditions pass else returns false.
	return (row >= 0) && (row < len(graph)) && (col >=0 ) && (col < len(graph)) && (graph[row][col] == 1) && !visited[row][col]
}

// perform DFS by considering each cell in a 2d grid as a node
func DepthFirstSearchGrid(graph [][]int, row int, col int, visited [][]bool) {
	// check if the cells adjacent to (row, col) is connected
	adjRows := []int {-1, -1, -1, 0, 0, 1, 1, 1}
	adjCols := []int {-1, 0, 1, -1, 1, -1, 0, 1}
	// set current row and col as visited
	visited[row][col] = true
	// loop through all the adjacent cell to recursively find connected components
	for k:=0; k<8; k++ {
		if isValied(graph, row + adjRows[k], col + adjCols[k], visited) {
			DepthFirstSearchGrid(graph, row + adjRows[k], col + adjCols[k], visited) // perform recursion
		}
	}
}

// count connected componenets or islands
func FindConnectedComponents(graph [][]int) int {
	visited := make([][]bool, len(graph))
	for i := range visited {
		visited[i] = make([]bool, len(graph))
	}
	count := 0
	for i:=0; i<len(graph); i++ {
		for j:=0; j<len(graph[i]); j++ {
			// perform DFS traversal and mark visited nodes
			if graph[i][j] == 1 && !visited[i][j] {
				count++
				DepthFirstSearchGrid(graph, i, j, visited)
			}
		}
	}
	return count
}