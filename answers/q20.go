package answers

import "fmt"

func Day20() []interface{} {
	data := ReadInputAsInt(20)
	return []interface{}{q20part1(data), q20part2(data)}
}

type Node struct {
	prev  *Node
	next  *Node
	value int64
	list  *LinkedList
}

type LinkedList struct {
	start  *Node
	end    *Node
	length int64
}

func (l *LinkedList) Print() {
	pos := l.start
	for i := int64(0); i < l.length; i++ {
		fmt.Printf("%d,", pos.value)
		pos = pos.next
	}
	fmt.Printf("\n")
}

func (L *LinkedList) AddToEnd(key int64) *Node {
	node := &Node{
		value: key,
		list:  L,
	}
	L.length++
	if L.start == nil && L.end == nil {
		L.start = node
		L.end = node
		node.next = node
		node.prev = node
		return node
	}

	end := L.end
	node.prev = end
	node.next = L.start
	L.end.next = node
	L.start.prev = node
	L.end = node
	return node
}

func (L *LinkedList) Close() {
	L.end.next = L.start
	L.start.prev = L.end
}

func (n *Node) Next(num int64) *Node {
	sol := n
	for i := int64(0); i < num; i++ {
		sol = sol.next
	}
	return sol
}

func (n *Node) Previous(num int64) *Node {
	sol := n
	for i := int64(0); i < abs64(num); i++ {
		sol = sol.prev
	}
	return sol
}

func (n *Node) Remove() {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
}

func (pos_n *Node) InsertBefore(new *Node) {
	prev := pos_n.prev
	prev.next = new
	pos_n.prev = new
	new.prev = prev
	new.next = pos_n
}

func (pos_n *Node) InsertAfter(new *Node) {
	next := pos_n.next
	next.prev = new
	pos_n.next = new
	new.next = next
	new.prev = pos_n
}

func (L *LinkedList) Move(node *Node, steps int64) {

	// No need to loop around multiple times
	if abs64(steps) >= (L.length - 1) {
		if steps > 0 {
			steps = steps % (L.length - 1)
		} else {
			steps = steps % (L.length - 1)
		}
	}

	// Remove from existing link
	if steps == 0 {
		return
	}
	node.Remove()

	// If we're the start or the end, change to the next value
	if node == L.start && node.value != 0 {
		L.start = node.next
	}
	if node == L.end && node.value != 0 {
		L.end = node.prev
	}

	if steps < 0 {
		target := node.Previous(steps)
		if L.start == target {
			L.end = node
		}
		target.InsertBefore(node)
	} else if steps > 0 {
		target := node.Next(steps)
		if L.end == target {
			L.start = node
		}
		target.InsertAfter(node)
	}
}

func ParseLinkedList(data []int, key int64) (LinkedList, []*Node) {
	list := LinkedList{}
	nodes := make([]*Node, len(data))
	for i, row := range data {
		nodes[i] = list.AddToEnd(int64(row) * int64(key))
	}
	list.Close()
	return list, nodes
}

func q20part1(data []int) int64 {
	list, nodes := ParseLinkedList(data, 1)
	var originNode *Node
	for _, node := range nodes {
		list.Move(node, node.value)
		if node.value == 0 {
			originNode = node
		}
	}
	sum := int64(0)
	for i := 0; i < 3; i++ {
		originNode = originNode.Next(1000)
		sum += originNode.value
	}

	return sum
}

func q20part2(data []int) int64 {
	decryptionKey := int64(811589153)
	list, nodes := ParseLinkedList(data, decryptionKey)
	var originNode *Node
	for iter := 0; iter < 10; iter++ {
		for _, node := range nodes {
			list.Move(node, node.value)
			if node.value == 0 {
				originNode = node
			}
		}
	}

	sum := int64(0)
	for i := 0; i < 3; i++ {
		originNode = originNode.Next(1000)
		sum += originNode.value
	}
	return sum
}
