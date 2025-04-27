package arraystring

import "strings"

func fullJustify(words []string, maxWidth int) []string {

	var answer []string
	var subList []string
	// create a subList []string
	// keep adding words to it until it exceeds the maxWidth
	i := 1
	currLen := len(words[0])
	subList = []string{words[0]}
	n := len(words)

	var subString string

	for i < n { //until you reach end of list
		if currLen+len(words[i]) < maxWidth {
			currLen += 1
			subList = append(subList, words[i])
			currLen += len(words[i])
		} else {
			currLen -= len(subList) - 1 //subtract extra added spaces
			spaces := maxWidth - currLen
			// fmt.Printf("spaces %v\n",spaces)
			// fmt.Printf("subList %v\n",subList)
			if len(subList) > 1 {
				//completeJustify
				subString = completeJustify(subList, spaces)
			} else {
				//leftJustify
				subString = leftJustify(subList, spaces)
			}
			answer = append(answer, subString)
			//reset
			subList = []string{words[i]}
			currLen = len(words[i])
			subString = ""
		}
		i++
	}

	currLen -= len(subList) - 1 //subtract extra added spaces
	spaces := maxWidth - currLen
	// fmt.Printf("spaces %v\n",spaces)
	// fmt.Printf("subList %v\n",subList)
	// //handle the last set of entries
	subString = leftJustify(subList, spaces)
	answer = append(answer, subString)

	return answer
}

func completeJustify(subList []string, spaces int) string {
	for spaces > 0 {
		for i := 0; i < len(subList)-1; i++ {
			subList[i] += " "
			spaces--
			if spaces == 0 {
				break
			}
		}
	}

	return getStringFromList(subList)
}

func leftJustify(subList []string, spaces int) string {
	//add one space after every word
	if spaces > 0 {
		for i := range subList {
			subList[i] += " "
			spaces--
			if spaces == 0 {
				break
			}
		}
	}

	// add remaining spaces at the end
	for spaces > 0 {
		subList[len(subList)-1] += " "
		spaces--
		if spaces == 0 {
			break
		}
	}

	return getStringFromList(subList)
}

func getStringFromList(subList []string) string {
	var sb strings.Builder
	for _, subWord := range subList {
		sb.WriteString(subWord)
	}

	return sb.String()
}
