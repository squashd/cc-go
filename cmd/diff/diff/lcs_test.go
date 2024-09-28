package diff

import "testing"

func TestLargestCommonSequence(t *testing.T) {
	type testCase struct {
		s1       string
		s2       string
		expected string
	}

	testCases := []testCase{
		{
			s1:       "ABCDEF",
			s2:       "ABCDEF",
			expected: "ABCDEF",
		},
		{
			s1:       "ABC",
			s2:       "XYZ",
			expected: "",
		},
		{
			s1:       "AABCXY",
			s2:       "XYZ",
			expected: "XY",
		},
		{
			s1:       "",
			s2:       "",
			expected: "",
		},
		{
			s1:       "ABCD",
			s2:       "AC",
			expected: "AC",
		},
	}

	for _, tc := range testCases {
		o := LargestCommonSequence(tc.s1, tc.s2)
		if o != tc.expected {
			t.Errorf("Expected %s but got %s", tc.expected, o)
		}
	}
}
