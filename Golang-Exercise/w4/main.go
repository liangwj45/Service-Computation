package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	n := len(nodes)
	for i := n/2 - 1; i >= 0; i-- {
		down(nodes, i, n)
	}
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	for {
		left, right, min := 2*i+1, 2*i+2, i
		if left < n && nodes[left].Value < nodes[min].Value {
			min = left
		}
		if right < n && nodes[right].Value < nodes[min].Value {
			min = right
		}
		if min == i {
			break
		}
		nodes[i], nodes[min] = nodes[min], nodes[i]
		i = min
	}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	for {
		k := (j - 1) / 2
		if j <= 0 || nodes[j].Value > nodes[k].Value {
			break
		}
		nodes[j], nodes[k] = nodes[k], nodes[j]
		j = k
	}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	n := len(nodes) - 1
	nodes[0], nodes[n] = nodes[n], nodes[0]
	down(nodes, 0, n)
	return nodes[n], nodes[0:n]
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes)-1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	var temp []Node
	var top Node
	for {
		top, nodes = Pop(nodes)
		if top.Value == node.Value {
			break
		}
		temp = append(temp, top)
	}
	for _, ele := range temp {
		nodes = Push(ele, nodes)
	}
	return nodes
}

func main() {
	nodes := []Node{{8}, {4}, {5}, {1}, {2}, {5}, {1}, {6}}
	Init(nodes)

	nodes = Push(Node{0}, nodes)
	n, nodes := Pop(nodes)
	fmt.Println(n) // 0

	n, nodes = Pop(nodes)
	fmt.Println(n) // 1

	nodes = Remove(nodes, Node{4})
	n, nodes = Pop(nodes)
	fmt.Println(n) // 1
}
