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
