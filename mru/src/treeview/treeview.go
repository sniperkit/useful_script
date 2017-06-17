package treeview

import (
	"strings"
)

type Node struct {
	Key      string  `json:"text"`
	Children []*Node `json:"nodes"`
}

func New(root string) *Node {
	return &Node{
		Key:      root,
		Children: make([]*Node, 0, 1),
	}
}

func (n *Node) AddChild(key string) {
	if !strings.Contains(key, ":::") {
		n.Children = append(n.Children, &Node{
			Key:      key,
			Children: make([]*Node, 0, 1),
		})
		return
	}

	levels := strings.Split(key, ":::")
	if len(levels) == 0 {
		return
	}

	var inserted = false
	for _, c := range n.Children {
		if levels[0] == c.Key {
			inserted = true
			c.AddChild(key[len(levels[0])+len(":::"):])
		}
	}

	if !inserted {
		newnode := &Node{
			Key:      levels[0],
			Children: make([]*Node, 0, 1),
		}

		n.Children = append(n.Children, newnode)
		newnode.AddChild(key[len(levels[0])+len(":::"):])
	}
}
