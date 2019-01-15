package main

type ListNode struct{
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil{
		return head
	}

	index := 0
	ptr := head

	for;ptr!=nil;{
		if index % k == 0{
			for i:=0;i<k;i++{
				
			}
		}else{
			ptr = ptr.Next
		}
	}

	return head
}

func main() {
	
}
