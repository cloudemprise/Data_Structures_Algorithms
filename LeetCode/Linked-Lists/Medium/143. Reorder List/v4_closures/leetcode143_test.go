package main

import "testing"

func TestReorderList(t *testing.T) {
	// Test case 1: Reorder a linked list with even number of nodes
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reorderList(head1)
	if head1.Val != 1 || head1.Next.Val != 4 || head1.Next.Next.Val != 2 || head1.Next.Next.Next.Val != 3 {
		t.Errorf("Test case 1 failed: expected [1, 4, 2, 3], but got %v", getListValues(head1))
	}

	// Test case 2: Reorder a linked list with odd number of nodes
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reorderList(head2)
	if head2.Val != 1 || head2.Next.Val != 5 || head2.Next.Next.Val != 2 || head2.Next.Next.Next.Val != 4 || head2.Next.Next.Next.Next.Val != 3 {
		t.Errorf("Test case 2 failed: expected [1, 5, 2, 4, 3], but got %v", getListValues(head2))
	}

	// Test case 3: Reorder a linked list with only one node
	head3 := &ListNode{Val: 1}
	reorderList(head3)
	if head3.Val != 1 {
		t.Errorf("Test case 3 failed: expected [1], but got %v", getListValues(head3))
	}

}

// Helper function to get the values of a linked list
func getListValues(head *ListNode) []int {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
