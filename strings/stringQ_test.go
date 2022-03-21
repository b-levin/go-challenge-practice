package strings

import (
	"fmt"
	"testing"
)

func TestReserveWords(t *testing.T) {
	var tests = []struct {
		sentence string
		want     string
	}{
		{"i.like.this.program.very.much", "much.very.program.this.like.i"},
		{"pqr.mno", "mno.pqr"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf(tt.sentence)
		t.Run(testName, func(t *testing.T) {
			ans := reverseWords(tt.sentence)
			if ans != tt.want {
				t.Errorf("Got %s, want: %s", ans, tt.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		word string
		want string
	}{
		{"aaaabbaa", "aabbaa"},
		{"abc", "a"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf(tt.word)
		t.Run(testName, func(t *testing.T) {
			ans := longestPalindrome(tt.word)
			if ans != tt.want {
				t.Errorf("Got %s, want: %s", ans, tt.want)
			}
		})
	}
}

func TestRomanNumToInt(t *testing.T) {
	var tests = []struct {
		word string
		want int
	}{
		{"V", 5},
		{"III", 3},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf(tt.word)
		t.Run(testName, func(t *testing.T) {
			ans := romanNumToInt(tt.word)
			if ans != tt.want {
				t.Errorf("Got %d, want: %d", ans, tt.want)
			}
		})
	}
}
