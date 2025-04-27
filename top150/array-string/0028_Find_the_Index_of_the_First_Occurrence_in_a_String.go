package arraystring

func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			check := checkSubStr(i, haystack, needle)
			if check {
				return i
			}
		}
	}

	return -1
}

func checkSubStr(start int, haystack string, needle string) bool {
	if start+len(needle) <= len(haystack) { //so we dont overrun the idx for haystack
		for i, j := start, 0; i < start+len(needle); i, j = i+1, j+1 {
			if haystack[i] != needle[j] {
				return false
			}
			// fmt.Printf("%d:%c %d:%c\n",i,haystack[i],j,needle[j])
		}
		return true //went through whole loop and all chars matched
	}
	return false
}
