package platypus

import (
	"strings"
)

type action struct {
	handler Handler
}

type node struct {
	children []*node
	key      string
	isParam  bool
	action   Handler
}

func (n *node) insertNode(path string, handler Handler) {

	keys := strings.Split(path, "*")[1:]
	count := len(keys)

	for {
		fNode, key := n.traverse(keys, nil)
		if fNode.key == key && count == 1 {
			fNode.action = handler
			return
		}

		nNode := node{key: key, isParam: false}

		if len(key) > 0 && key[0] == ':' { // check if it is a named param.
			nNode.isParam = true
		}

		if count == 1 { // this is the last component of the url resource, so it gets the handler.
			nNode.action = handler
		}

		fNode.children = append(fNode.children, &nNode)
		count--
		if count == 0 {
			break
		}

	}
}

// TODO:pass the handler func a transformer that adds named params
// type Transformer func(string)string
func (n *node) traverse(keys []string, params Params) (*node, string) {
	key := keys[0]

	if len(n.children) > 0 {
		for _, child := range n.children {
			if child.key == key || child.isParam {
				if child.isParam && params != nil {
					params.Add(child.key[1:], key)
				}

				next := keys[1:]
				if len(next) > 0 {
					return child.traverse(next, params) // tail recursion is it's own reward.
				}
				return child, key
			}
		}
	}

	return n, key
}

func (n *node) isLeaf() bool {
	return len(n.children) <= 2
}
