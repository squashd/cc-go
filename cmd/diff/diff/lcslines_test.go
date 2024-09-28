package diff

import "testing"

func TestFindCommonLines(t *testing.T) {
	type testCase struct {
		original []string
		new      []string
		expected []string
	}

	testCases := []testCase{
		{
			original: []string{"This is a test which contains:", "this is the lcs"},
			new:      []string{"this is the lcs", "we're testing"},
			expected: []string{"this is the lcs"},
		},
		{
			original: []string{
				"Coding Challenges helps you become a better software engineer through that build real applications.",
				"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
				"I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
				"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
			},
			new: []string{
				"Helping you become a better software engineer through coding challenges that build real applications.",
				"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
				"These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
				"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
			},
			expected: []string{
				"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
				"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
			},
		},
		{
			original: []string{
				"Start of text",
				"Line A",
				"Unrelated line",
				"Line B",
				"End of text",
			},
			new: []string{
				"Line A",
				"Some different line",
				"Line B",
			},
			expected: []string{"Line A", "Line B"},
		},
		{
			original: []string{
				"Start of text",
				"Line A",
				"Unrelated line",
				"where did this come from?",
				"Line B",
				"End of text",
			},
			new: []string{
				"Line A",
				"Some different line",
				"where did this come from?",
				"Line B",
			},
			expected: []string{"Line A", "where did this come from?", "Line B"},
		},
		{
			original: []string{
				"Intro line",
				"Line 1",
				"Line 2",
				"Line 3",
				"Conclusion",
			},
			new: []string{
				"Intro line",
				"Line 3",
				"Line 2",
				"Line 1",
				"Conclusion",
			},
			expected: []string{"Intro line", "Line 1", "Conclusion"},
		},
	}

	for _, tc := range testCases {
		output := LCSLines(tc.original, tc.new)
		if !compareLines(output, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, output)
		}
	}
}

func compareLines(lines1, lines2 []string) bool {
	if len(lines1) != len(lines2) {
		return false
	}
	for i := range lines1 {
		if lines1[i] != lines2[i] {
			return false
		}
	}
	return true
}
