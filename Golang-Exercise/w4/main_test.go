package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	nodes := []Node{{8}, {9}, {3}, {7}, {6}, {4}, {1}, {5}, {0}, {2}}
	Init(nodes)
	expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range expect {
		var node Node
		node, nodes = Pop(nodes)
		if node.Value != expect[i] {
			t.Errorf("node.Value != sorted[i], expected %d, found %d", expect[i], node.Value)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	var node Node
	nodes := []Node{}
	nodes = Push(Node{3}, nodes)
	nodes = Push(Node{2}, nodes)
	nodes = Push(Node{1}, nodes)
	node, nodes = Pop(nodes)
	if node.Value != 1 {
		t.Errorf("expected %d, found %d", 1, node.Value)
	}
	nodes = Push(Node{3}, nodes)
	nodes = Push(Node{4}, nodes)
	node, nodes = Pop(nodes)
	if node.Value != 2 {
		t.Errorf("expected %d, found %d", 2, node.Value)
	}
}

func TestRemove(t *testing.T) {
	nodes := []Node{{8}, {9}, {3}, {7}, {6}, {4}, {1}, {5}, {0}, {2}}
	Init(nodes)
	nodes = Remove(nodes, Node{3})
	expect := []int{0, 1, 2, 4, 5, 6, 7, 8, 9}
	for i := range expect {
		var node Node
		node, nodes = Pop(nodes)
		if node.Value != expect[i] {
			t.Errorf("node.Value != sorted[i], expected %d, found %d", expect[i], node.Value)
		}
	}
}
