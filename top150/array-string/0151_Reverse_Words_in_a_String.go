package arraystring

import "strings"

//iterate through the string from behind, until you find a char
//if you find a char thats end , keep going back again, if you find space, the char in front of it is the start
//from start to end write all chars to strings.Builder

func reverseWords(s string) string {
	f, l := -1, -1 //char indexes
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if !isSpace(s[i]) && l == -1 {
			l = i
		}
		if isSpace(s[i]) && l != -1 {
			f = i + 1
		}
		if i == 0 && !isSpace(s[i]) && l != -1 {
			f = i
		}
		if f != -1 && l != -1 {
			if sb.Len() > 0 {
				sb.WriteByte(' ')
			}
			for f <= l {
				sb.WriteByte(s[f])
				f++
			}
			f, l = -1, -1
		}
	}
	return sb.String()
}

func isSpace(b byte) bool {
	return b == ' '
}
