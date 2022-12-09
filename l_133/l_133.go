package leetcode

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)

	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	visited[node] = &Node{Val: node.Val}
	for _, n := range node.Neighbors {
		visited[node].Neighbors = append(visited[node].Neighbors, clone(n, visited))
	}
	return visited[node]
}
