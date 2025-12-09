package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	vertexesFirst := []*TreeNode{p}
	vertexesSecond := []*TreeNode{q}

	for len(vertexesFirst) != 0 || len(vertexesSecond) != 0 {
		var newVertexesFirst []*TreeNode
		var newVertexesSecond []*TreeNode
		if len(vertexesFirst) != len(vertexesSecond) {
			return false
		}

		for i := 0; i < len(vertexesFirst); i++ {
			if vertexesFirst[i] != nil && vertexesSecond[i] != nil {
				if vertexesFirst[i].Val == vertexesSecond[i].Val {
					newVertexesFirst = append(newVertexesFirst, vertexesFirst[i].Left)
					newVertexesFirst = append(newVertexesFirst, vertexesFirst[i].Right)
					newVertexesSecond = append(newVertexesSecond, vertexesSecond[i].Left)
					newVertexesSecond = append(newVertexesSecond, vertexesSecond[i].Right)
				} else {
					return false
				}
			} else if vertexesFirst[i] == nil && vertexesSecond[i] == nil {
				// do nothing
			} else if vertexesFirst[i] == nil && vertexesSecond[i] != nil {
				return false
			} else {
				return false
			}
		}
		vertexesFirst = newVertexesFirst
		vertexesSecond = newVertexesSecond
	}

	return true
}
