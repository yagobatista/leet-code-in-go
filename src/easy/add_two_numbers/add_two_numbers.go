package add_two_numbers

import (
	"fmt"
	"leet-code/src/pkg/stack"
	"log"
	"strconv"
	"strings"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return makeNumber(
		sum(l1) + sum(l2),
	)
}

func sum(list *ListNode) int {
	if list == nil {
		return 0
	}

	var stack stack.Stack[int]
	for node := list; node != nil; node = node.Next {
		stack.Push(node.Val)
	}

	var sum string
	for !stack.IsEmpty() {
		sum = fmt.Sprintf("%s%d", sum, stack.Pop())
	}

	number, err := strconv.Atoi(sum)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func makeNumber(number int) *ListNode {
	digits := strings.Split(fmt.Sprintf("%d", number), "")

	var list *ListNode
	for i := len(digits) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(digits[i])
		if err != nil {
			log.Fatal(err)
		}

		list = appendToList(list, number)
	}

	return list
}

func appendToList(head *ListNode, val int) *ListNode {
	node := &ListNode{
		Val: val,
	}

	if head == nil {
		return node
	}

	for element := head; element != nil; element = element.Next {
		if element.Next != nil {
			continue
		}

		element.Next = node
		return head
	}

	return head
}
