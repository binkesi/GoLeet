package binarytree

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	graph := make([][]int, 500)
	dfs(root, graph)
	type pairT struct {
		val  int
		step int
	}
	queue := []pairT{{target.Val, k}}
	visited := make([]bool, 500)
	visited[target.Val] = true
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		next, step := graph[p.val], p.step
		if step == 0 {
			ans = append(ans, p.val)
		}
		for i := 0; i < len(next); i++ {
			if visited[next[i]] == false {
				visited[next[i]] = true
				queue = append(queue, pairT{next[i], step - 1})
			}
		}
	}
	return ans
}

func dfs(root *TreeNode, graph [][]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		graph[root.Val] = append(graph[root.Val], root.Left.Val)
		graph[root.Left.Val] = append(graph[root.Left.Val], root.Val)
		dfs(root.Left, graph)
	}
	if root.Right != nil {
		graph[root.Val] = append(graph[root.Val], root.Right.Val)
		graph[root.Right.Val] = append(graph[root.Right.Val], root.Val)
		dfs(root.Right, graph)
	}
}
