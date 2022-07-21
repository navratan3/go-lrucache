package models

type NodeData struct {
	Value interface{}
	Key   int
}
type Node struct {
	Prev *Node
	Next *Node
	Data NodeData
}

type Queue struct {
	capacity int
	size     int
	front    *Node
	rear     *Node
}

func (q *Queue) SetCapacity(capacity int) {
	q.capacity = capacity
}

func (q *Queue) Insert(key int, data interface{}) (*Node, bool) {
	if q.size >= q.capacity {
		return nil, false
	}
	node := new(Node)
	node.Data = NodeData{
		Value: data,
		Key:   key,
	}
	if q.front == nil && q.rear == nil {
		q.front = node
		q.rear = node
	} else {
		node.Prev = nil
		node.Next = q.front
		q.front.Prev = node
		q.front = node
	}
	q.size++
	return node, true
}

func (q *Queue) Delete() (*Node, bool) {
	if q.size == 0 {
		return nil, false
	}
	deletedNode := q.rear
	temp := q.rear.Prev
	temp.Next = nil
	q.rear = temp
	q.size--
	return deletedNode, true
}

func (q *Queue) GetFront() *Node {
	return q.front
}

func (q *Queue) GetRear() *Node {
	return q.rear
}

func (q *Queue) MoveToFront(val *Node) {
	if q.front == val {
		return
	} else if q.rear == val {
		q.rear = val.Prev
		q.rear.Next = nil

	} else {
		temp := val.Prev
		temp.Next = val.Next
		val.Next.Prev = temp
	}
	q.front.Prev = val
	val.Prev = nil
	val.Next = q.front
	q.front = val
}
