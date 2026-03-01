package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU   ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "hello\tworld\n",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  multiple   spaces  between  words  ",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("String length does not match")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word does not match")
				return
			}
		}
	}
}
