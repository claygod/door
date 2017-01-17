package door

// The multiplexer `Door`
// Node
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// Nodes are stored in the index. Used for route search.
type node struct {
	child map[typeHash]*node
	route *route
}

func newNode() *node {
	nNode := &node{}
	nNode.child = make(map[typeHash]*node)
	return nNode
}
