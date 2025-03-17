package main

import (
	"fmt"

	"github.com/mdshahjahanmiah/go-pointer-chaos/linkedlist"
)

func main() {
	// Correct linked list
	fmt.Println("Correct Append:")
	list1 := linkedlist.LinkedList{}
	list1.Append("Miah")
	list1.Append("Md")
	list1.Append("Shahjahan")
	list1.Print()

	// Corrupted linked list
	fmt.Println("\nCorrupted Append:")
	list2 := linkedlist.LinkedList{}
	list2.CorruptedAppend("Miah")
	list2.CorruptedAppend("Md")
	list2.CorruptedAppend("Shahjahan")
	list2.Print()
}
