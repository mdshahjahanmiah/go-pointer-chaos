package linkedlist

import (
	"log/slog"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data string) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head

	slog.Info("starting list traversal", "node", current.data)
	for current.next != nil {
		slog.Info("traversing list", "node", current.data)
		current = current.next
	}
	slog.Info("attaching new node", "new_node", data, "at_node", current.data)

	current.next = newNode
}

func (l *LinkedList) CorruptedAppend(data string) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("operation panicked", "error", r, "operation", "CorruptedAppend", "data", data, "consequence", "Node not properly added")
			slog.Info("attempting recovery", "action", "attaching to head as fallback")
			l.head.next = &Node{data: data, next: nil}
		}
	}()

	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head

	slog.Info("starting traversal", "node", current.data)
	for current.next == nil { // incorrect condition
		slog.Warn("incorrect traversal detected", "detail", "only processes last node")
		current = current.next
		slog.Error("traversal error", "current", current)
	}
	slog.Info("attaching node", "new_node", data, "at_node", current.data)

	current.next = newNode
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		if current.next != nil {
			print(current.data + " -> ")
		} else {
			print(current.data + " -> END\n")
		}
		current = current.next
	}
	if l.head == nil {
		print("Empty list -> END\n")
	}
}
