package test

type Node struct {
	Value int
	Next  *Node
}

func delete(root *Node, deleted *Node) {
	if deleted == root {
		p := root.Next
		if p == nil {
			root = nil
			return
		}
		root.Value = p.Value
		root.Next = p.Next
	} else {
		current := root
		for current.Next != nil {
			next := current.Next
			if next == deleted {
				current.Next = next.Next
				break
			} else {
				current = next
			}
		}
	}
}
