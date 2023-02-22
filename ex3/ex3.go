package ex3

func lengthOfLongestSubstring(s string) int {
	start, longest := 0, 0
	m := make(map[int32]int)

	for cur, c := range s {
		if pos, ok := m[c]; ok && (pos >= start) {
			start = pos + 1
		}

		length := cur - start + 1
		if length > longest {
			longest = length
		}

		m[c] = cur
	}

	return longest
}
