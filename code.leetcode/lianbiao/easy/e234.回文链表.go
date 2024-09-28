package lianbiao

func reverseListhere(head *ListNode) *ListNode {
	var curl *ListNode
	var pre *ListNode
	curl = head
	for curl != nil {
		next := curl.Next
		curl.Next = pre
		pre = curl
		curl = next
	}

	return pre

}

func isPalindrome(head *ListNode) bool {
	return reverseListhere(head) == head
}
