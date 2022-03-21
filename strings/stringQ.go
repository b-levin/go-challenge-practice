package strings

import (
	"strings"
)

//https://practice.geeksforgeeks.org/problems/reverse-words-in-a-given-string5459/1
func reverseWords(sentence string) string {
	sentArr := strings.Split(sentence, ".")
	out := ""
	for i := len(sentArr) - 1; i >= 0; i-- {
		out += sentArr[i]
		if i-1 >= 0 {
			out += "."
		}
	}
	return out
}

type Stack []rune

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func reverse(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//Doesn't cover palindromes in middle of word
//https://practice.geeksforgeeks.org/problems/longest-palindrome-in-a-string3411/1
func longestPalindrome(word string) string {
	out := string(word[0])
	// start to end
	for i := 2; i < len(word); i++ {
		reverse := reverse(word[:i])
		if reverse == word[:i] && len(reverse) > len(out) {
			out = reverse
		}
	}
	// end to start
	for i := len(word) - 2; i >= 0; i-- {
		reverse := reverse(word[i:])
		if reverse == word[i:] && len(reverse) > len(out) {
			out = reverse
		}
	}
	// no palindrome
	return out
}

//https://practice.geeksforgeeks.org/problems/roman-number-to-integer3201/1
func romanNumToInt(number string) int {
	key := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	out := 0
	for _, r := range number {
		out += key[r]
	}
	return out
}
