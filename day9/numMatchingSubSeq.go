func numMatchingSubseq(s string, words []string) int {
	strMap := make(map[rune][]int)

	// Initialise strMap
	for i, c := range s {
		if val, ok := strMap[c]; ok {
			strMap[c] = append(val, i)
		} else {
			strMap[c] = []int{i}
		}
	}

	solCount := 0

	for _, word := range words {
		i := 0
		var c rune
		prevIdx := -1

		for i, c = range word {
			// Check if letter exists in map
			if idxs, ok := strMap[c]; ok {
				idxFound := false

				// Find index of letter > than index of previously found letter
				for _, idx := range idxs {
					if idx > prevIdx {
						prevIdx = idx
						idxFound = true
						break
					}
				}

				// If no valid index is found, break
				if !idxFound {
					i = 0
					break
				}
			} else {
				i = 0
				break
			}

		}

		// Word is a subsequence if all letters were successfully found in strMap
		if i == len(word)-1 {
			solCount += 1
		}
	}

	return solCount
}
