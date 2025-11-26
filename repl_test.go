package main

import "testing"

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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "    one two    three four  five",
			expected: []string{"one", "two", "three", "four", "five"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("\n"+"expected: %#v\n"+" but got: %#v\n", c.expected, actual)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("\n"+
				"expected: %#v\n"+
					" but got: %#v\n"+
					"index %v is incorrect", c.expected, actual, i)
				break
			}
		}
	}
}
