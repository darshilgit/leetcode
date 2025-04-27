package arraystring

func romanToInt(s string) int {
	curr_sum := 0
	n := len(s) - 1
	for i := range s {
		switch s[i] {
		case 'I': //can come before V & X
			// fmt.Printf("I\n")
			if i+1 <= n {
				if s[i+1] == 'V' || s[i+1] == 'X' {
					curr_sum -= 1
					continue
				}
			}
			curr_sum += 1
		case 'V':
			// fmt.Printf("V\n")
			curr_sum += 5
		case 'X': //can come before L & C
			// fmt.Printf("X\n")
			if i+1 <= n {
				if s[i+1] == 'L' || s[i+1] == 'C' {
					curr_sum -= 10
					continue
				}
			}
			curr_sum += 10
		case 'L':
			// fmt.Printf("L\n")
			curr_sum += 50
		case 'C': //can come before D & M
			// fmt.Printf("C\n")
			if i+1 <= n {
				if s[i+1] == 'D' || s[i+1] == 'M' {
					curr_sum -= 100
					continue
				}
			}
			curr_sum += 100
		case 'D':
			// fmt.Printf("D\n")
			curr_sum += 500
		case 'M':
			// fmt.Printf("M\n")
			curr_sum += 1000
		}
		// fmt.Printf("%v\n",curr_sum)
	}

	return curr_sum

}
