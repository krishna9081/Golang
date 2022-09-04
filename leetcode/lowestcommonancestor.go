/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {

	pDepth := getDepth(p)
	qDepth := getDepth(q)

	// Level up the Nodes such that they both nodes are on the same level

	if pDepth > qDepth {
		for i := 0; i < pDepth-qDepth; i++ {
			p = p.Parent
		}
	} else if pDepth < qDepth {
		for i := 0; i < qDepth-pDepth; i++ {
			q = q.Parent
		}
	}
	//Now both are on the same level
	for p != q {
		p = p.Parent
		q = q.Parent
	}

	return p

}

func getDepth(p *Node) int {

	depth := 0

	for p != nil {
		p = p.Parent
		depth += 1
	}

	return depth

}