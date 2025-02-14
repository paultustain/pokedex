package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this is a sentence",
			expected: []string{"this", "is", "a", "sentence"},
		},
		{
			input:    "hello, world",
			expected: []string{"hello,", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"Different lengths got %d when %d was expected for %s",
				len(actual),
				len(c.expected),
				c.input,
			)
			continue

		}

		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			if word != expected_word {
				t.Errorf("Different words")
			}
		}
	}
}
