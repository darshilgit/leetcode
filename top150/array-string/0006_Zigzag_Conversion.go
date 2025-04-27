package arraystring

import "strings"

func convert(s string, numRows int) string {

	size := len(s)
	if numRows >= size || numRows == 1 {
		return s
	}
	var sb strings.Builder
	for i := 0; i < numRows; i++ {
		ncIndex := i
		// fmt.Printf("\n\ncurrent row : %d\n",i)
		sb.WriteByte(s[ncIndex])
		// fmt.Printf("char: %c\n",s[ncIndex])
		for ncIndex < size {
			if i > 0 && i < numRows-1 {
				ncIndex += 2 * (numRows - 1 - i)
				// fmt.Printf("%v ncIndex : %d\n",i,ncIndex)
				if ncIndex < size {
					// fmt.Printf("char: %c\n",s[ncIndex])
					sb.WriteByte(s[ncIndex])
				}
				ncIndex += 2 * i
				// fmt.Printf("%v ncIndex : %d\n",i,ncIndex)
				if ncIndex < size {
					// fmt.Printf("char: %c\n",s[ncIndex])
					sb.WriteByte(s[ncIndex])
				}
			} else if i == 0 {
				ncIndex += 2 * (numRows - 1 - i)
				// fmt.Printf("%v ncIndex : %d\n",i,ncIndex)
				if ncIndex < size {
					// fmt.Printf("char: %c\n",s[ncIndex])
					sb.WriteByte(s[ncIndex])
				}
			} else if i == numRows-1 {
				ncIndex += 2 * i
				// fmt.Printf("%v ncIndex : %d\n",i,ncIndex)
				if ncIndex < size {
					// fmt.Printf("char: %c\n",s[ncIndex])
					sb.WriteByte(s[ncIndex])
				}
			}
		}
	}

	return sb.String()
}

//not for 0 and N-1
// for i in range(1,N-2):
//     untilNextCharIndex goes out of bounds//
//     nextcharIdx := 2(N-1-i) = 2N-2-2i
//     writeChar
//     nextCharIdx := 2(i)
//     writeChar

// for i = 0 || i == N-1
//     n:4 i:0 nc: 8-2 = 6
//     curridx = 6
//     n:4 i:0 nc: 12

// 0   1   2   3   4   5   6   7   8   9   10   11  12  13
// P   A   Y   P   A   L   I   S   H   I   R    I   N   G

// -----------2-------------
// P0  Y2  A4   I6   H8   R10   N12
// A1  P3  L5   S7   I9   I11   G13
// -----------3-------------
// [0] P0        A4       H8         N12
// [1] A1   P3   L5  S7   I9   I11   G13
// [2] Y2        I6       R10
// -----------4-------------
// [0] P0           I6            N12
// [1] A1      L5   S7       I11  G13
// [2] Y2  A4       H8   R10
// [3] P3          I19
// -----------5-------------
// P0               H8
// A1          S7   I9
// Y2      I6       R10
// P3  L5           I11   G13
// A4               N12
// -----------6-------------
// P0                   R10
// A1               I9  I11
// Y2          H8       N12
// P3       S7          G13
// A4  I6
// L5
