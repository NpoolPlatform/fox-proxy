package router

import (
	"fmt"
)

type treeNode struct {
	nodes map[int]*treeNode
	val   interface{}
}

func (r *treeNode) registerRouter(
	val interface{},
	pathList ...int,
) {
	currentNode := r
	for _, v := range pathList {
		if currentNode.nodes == nil {
			currentNode.nodes = make(map[int]*treeNode)
		}
		if _, ok := currentNode.nodes[v]; !ok {
			currentNode.nodes[v] = &treeNode{}
		}
		currentNode = currentNode.nodes[v]
	}
	currentNode.val = val
}

func (r *treeNode) getVal(pathList ...int) (val interface{}, accurate bool, err error) {
	currentNode := r
	val = currentNode.val
	accurate = true
	for _, v := range pathList {
		accurate = false
		if currentNode.nodes[v] == nil {
			break
		}
		currentNode = currentNode.nodes[v]
		if currentNode.val != nil {
			val = currentNode.val
			accurate = true
		}
	}
	if val == nil {
		return nil, false, fmt.Errorf("cannot match router")
	}

	return val, accurate, nil
}
