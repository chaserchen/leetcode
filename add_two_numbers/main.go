package main

import (
	"fmt"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	count := 0

	fir := l1.Val + l2.Val
	if fir >= 10 {
		count = 1
		fir = fir - 10
	} else {
		count = 0
	}

	newC := &ListNode{
		Val:  fir,
		Next: &ListNode{},
	}

	var next = newC.Next

	index1 := 0
	map1 := make(map[int]int)
	for i := l1.Next; i != nil; i = i.Next {
		map1[index1] = i.Val
		index1++
	}

	index2 := 0
	map2 := make(map[int]int)
	for i := l2.Next; i != nil; i = i.Next {
		map2[index2] = i.Val
		index2++
	}

	if index1 == 0 && index2 == 0 && count == 0 {
		newC.Next = nil
		return newC
	}

	m := int(math.Max(float64(index1), float64(index2)))
	for i := 0; i < m; i++ {
		sum := map1[i] + map2[i] + count
		if sum >= 10 {
			sum = sum - 10
			count = 1
		} else {
			count = 0
		}
		node := next
		if i != m-1 || count != 0 {
			next = &ListNode{}
		} else {
			next = nil
		}
		node.Val = sum
		node.Next = next
	}
	if count != 0 {
		next.Val = 1
	}

	return newC
}

func main() {
	y := &ListNode{
		Val:  2,
		Next: nil,
	}
	z := &ListNode{
		Val:  5,
		Next: y,
	}
	k := &ListNode{
		Val:  8,
		Next: z,
	}

	yy := &ListNode{
		Val:  2,
		Next: nil,
	}
	zz := &ListNode{
		Val:  5,
		Next: yy,
	}
	kk := &ListNode{
		Val:  8,
		Next: zz,
	}

	value := addTwoNumbers(k, kk)
	fmt.Println(value)
}
