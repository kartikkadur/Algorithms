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

// count connected componenets or islands
func FindConnectedComponents(graph [][]int) int {
	visited := make([][]bool, len(graph), len(graph))
	count := 0
	for i:=0; i<len(graph); i++ {
	}
	return count
}