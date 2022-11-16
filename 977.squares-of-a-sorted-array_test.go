package leetcode

import "testing"

type testData struct {
	input    []int
	expected []int
}

var scenarios = []testData{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
}

func TestSortedSquares(t *testing.T) {
	for _, test := range scenarios {
		output := sortedSquares(test.input)

		for i, num := range output {
			if num != test.expected[i] {
				t.Errorf("Output %d not equal to expected %d", num, test.expected[i])
			}
		}
	}
}
