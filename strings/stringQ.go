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

func (s *Stack) peek() rune {
	if s.IsEmpty() {
		return -1
	} else {
		index := len(*s) - 1
		return (*s)[index]
	}
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return -1
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
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

//https://practice.geeksforgeeks.org/problems/remove-duplicates3034/1
func removeDup(word string) string {
	contains := make(map[rune]bool)
	out := ""
	for _, r := range word {
		if !contains[r] {
			contains[r] = true
			out += string(r)
		}
	}
	return out
}

//https://practice.geeksforgeeks.org/problems/form-a-palindrome1455/1
func makePalindrome(word string) int {
	if reverse(word) == word {
		//already a palindrome
		return 0
	} else {
		//not a palindrome (fun part!)
		var wordStack Stack
		for _, r := range word {
			if !wordStack.IsEmpty() && wordStack.peek() == r {
				wordStack.Pop()
			} else {
				wordStack.Push(r)
			}
		}
		if wordStack.Size() == len(word) {
			return wordStack.Size() - 1
		} else {
			return wordStack.Size()
		}
	}
}
