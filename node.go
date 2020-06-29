package platypus

import (
	"strings"
)

type action struct {
	handler Handler
}

type node struct {
	children      []*node
	key           string
	isParam       bool
	isLeaf        bool
	action        Handler
	transformFunc Transformer
}

func (n *node) insertNode(path string, handler Handler, ts Transformer) {
	keys := strings.Split(path, "*")[1:]
	count := len(keys)

	if ts == nil {
		ts = NilTransformer
	}

	for {
		fNode, key := n.traverse(keys, nil)
		if fNode.key == key && count == 1 {
			fNode.action = handler
			fNode.transformFunc = ts
			return
		}

		var leaf = false

		if key[len(key)-1] == '#' {
			leaf = true
		}

		nNode := node{key: key, transformFunc: ts, isParam: false, isLeaf: leaf}

		if len(key) > 0 && key[0] == ':' { // check if it is a named param.
			nNode.isParam = true
		}

		if count == 1 { // this is the last component of the url resource, so it gets the handler.
			nNode.action = handler
			nNode.transformFunc = ts
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
			if child.key == child.transformFunc(key) || child.isParam {
				if child.isParam && params != nil {
					ckey := child.key
					switch child.isLeaf {
					case true:
						params.Add(ckey[1:len(ckey)-1], key[:len(key)-1])
					case false:
						params.Add(ckey[1:], child.transformFunc(key))
					}
					// params.Add(child.key[1:], child.transformFunc(key))
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
