package arraystring

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	ret, up, down, peak := 1, 0, 0, 0

	for i := 0; i < len(ratings)-1; i++ {
		prev, curr := ratings[i], ratings[i+1]

		if prev < curr {
			up++
			down = 0
			peak = up
			ret += 1 + up
		} else if prev == curr {
			up, down, peak = 0, 0, 0
			ret += 1
		} else {
			up = 0
			down++
			ret += 1 + down
			if peak >= down {
				ret--
			}
		}
	}

	return ret
}
