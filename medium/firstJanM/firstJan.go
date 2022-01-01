package firstJanM

import (
	"fmt"
	"reflect"
)

type firstJanM struct {
	TestCases []testCase
}

type testCase struct {
	FirstNumber    ListNode
	SecondNumber   ListNode
	ExpectedResult ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewProblem() *firstJanM {
	return &firstJanM{
		TestCases: buildTestCases(),
	}
}

/*
	2. Add Two Numbers
	https://leetcode.com/problems/add-two-numbers/
*/
func (p firstJanM) Solution() {
	for _, v := range p.TestCases {
		actualResult := solver(&v.FirstNumber, &v.SecondNumber)
		if !reflect.DeepEqual(v.ExpectedResult, actualResult) {
			fmt.Println(fmt.Errorf("Test Case Failed. Expected result: %v, Acutal result: %v", v.ExpectedResult, actualResult))
			continue
		}
		fmt.Println("Test case passed.")
	}
}

func buildTestCases() []testCase {
	testCases := []testCase{
		{
			FirstNumber:    ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			SecondNumber:   ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			ExpectedResult: ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
		{
			FirstNumber:    ListNode{Val: 0, Next: nil},
			SecondNumber:   ListNode{Val: 0, Next: nil},
			ExpectedResult: ListNode{Val: 0, Next: nil},
		},
		{
			FirstNumber:    ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}},
			SecondNumber:   ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}},
			ExpectedResult: ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}},
		},
	}
	return testCases
}

func solver(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	sentinel := &ListNode{Val: 0, Next: nil}
	d := sentinel
	sum := 0
	for c1 != nil || c2 != nil {
		sum /= 10
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}

		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}
		d.Next = &ListNode{Val: sum % 10, Next: nil}
		d = d.Next
	}

	if sum/10 == 1 {
		d.Next = &ListNode{Val: 1, Next: nil}
	}

	return sentinel.Next
}
