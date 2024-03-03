package leetcode

import "testing"

type isPalindromeTestData struct {
	input    int
	expected bool
}

var isPalindromeScenarios = []isPalindromeTestData{
	{13231, true},
	{1337, false},
	{666, true},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range isPalindromeScenarios {
		output := isPalindrome(test.input)
		if output != test.expected {
			t.Errorf("Output for %d  is %t, not equal to expected %t", test.input, output, test.expected)
		}
	}
}
