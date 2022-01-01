package firstJan

import (
	"fmt"
	"reflect"
)

type firstJan struct {
	TestCases []testCase
}

type testCase struct {
	Num            []int
	ExpectedResult []int
}

func NewProblem() *firstJan {
	return &firstJan{
		TestCases: buildTestCases(),
	}
}

/*
	1920. Build Array from Permutation
	https://leetcode.com/problems/build-array-from-permutation/
*/
func (p firstJan) Solution() {
	for _, v := range p.TestCases {
		actualResult := solver(v.Num)
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
			Num:            []int{0, 2, 1, 5, 3, 4},
			ExpectedResult: []int{0, 1, 2, 4, 5, 3},
		},
		{
			Num:            []int{5, 0, 1, 2, 3, 4},
			ExpectedResult: []int{4, 5, 0, 1, 2, 3},
		},
	}
	return testCases
}

func solver(num []int) []int {
	result := make([]int, len(num))
	for i, v := range num {
		result[i] = num[v]
	}
	return result
}

// Its using O(1) space
func optimizedSolver(nums []int) []int {
	n := len(nums)
	for i, v := range nums {
		nums[i] = nums[i] + n*(nums[v]%n)
	}

	for i, v := range nums {
		nums[i] = v / n
	}
	return nums
}
