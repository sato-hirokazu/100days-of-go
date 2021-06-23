func numMatchingSubseq(S string, words []string) int {
	var res int
	exists := map[string]bool{}
	for _, w := range words {
		if exists[w] {
			res ++
			continue
		}
		flag := true
		p, q := -1, 1
		for _, c := range w {
			p := strings.Index(S[p+q:], string(c))
			if p == -1 {
				flag = false
				break
			}
			p, q = 0, q+p+1

		}
		if flag == true {
			exists[w] = true
			res++
		}
	}
	return res
}
