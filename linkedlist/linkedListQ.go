package linkedlist

import (
	"strconv"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func initList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) Len() int {
	return l.length
}

func (l *linkedList) add(value int) {
	node := &node{data: value}
	if l.Len() == 0 {
		l.head = node
		l.tail = node
		l.length++
	} else {
		l.tail.next = node
		l.tail = node
		l.length++
	}
}

func (l *linkedList) get(pos int) int {
	currentNode := l.head
	for i := 0; i < pos; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (l *linkedList) toString() string {
	out := ""
	current := l.head
	for current != nil {
		out += strconv.Itoa(current.data) + " "
		current = current.next
	}
	return out
}

func listFromArr(arr []int) *linkedList {
	list := initList()
	for _, v := range arr {
		list.add(v)
	}
	return list
}

//https://practice.geeksforgeeks.org/problems/finding-middle-element-in-a-linked-list/1
func findMiddle(list linkedList) int {
	return list.get(list.length / 2)
}
