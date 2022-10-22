package merge_sorted_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	var mergedList *ListNode

	fistListElement := list1
	secondListElement := list2

	for ; fistListElement != nil; fistListElement = fistListElement.Next {

		for ; secondListElement != nil; secondListElement = secondListElement.Next {
			if fistListElement.Val < secondListElement.Val {
				break
			}

			mergedList = Insert(mergedList, secondListElement.Val)
		}

		mergedList = Insert(mergedList, fistListElement.Val)
	}

	if secondListElement != nil {
		for ; secondListElement != nil; secondListElement = secondListElement.Next {
			mergedList = Insert(mergedList, secondListElement.Val)
		}
	}
	if fistListElement != nil {
		for ; fistListElement != nil; fistListElement = fistListElement.Next {
			mergedList = Insert(mergedList, fistListElement.Val)
		}
	}

	return mergedList
}

func Insert(head *ListNode, val int) *ListNode {
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
