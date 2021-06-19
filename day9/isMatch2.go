func isMatch(s string, p string) bool {
	si, pi := 0, 0
	for si < len(s) && pi < len(p) {
		// promise: p[pi] cannt be '*'
		if p[pi] != '.' {
			if pi+1 < len(p) && p[pi+1] == '*' {
				cutSKnife := p[pi]
				pi++
				for pi < len(p) {
					if p[pi] == '*' {
						pi++
					}
					break
				}
				childP := p[pi:]
				reCall := []bool{isMatch(s[si:], childP)}
				for si < len(s) {
					if s[si] != cutSKnife {
						break
					}
					si++
					reCall = append(reCall, isMatch(s[si:], childP))
				}
				for _, ans := range reCall {
					if ans {
						return true
					}
				}
				return false
			}
			if s[si] != p[pi] {
				return false
			}
			si++
			pi++
		} else {
			if pi+1 < len(p) && p[pi+1] == '*' {
				pi++
				for pi < len(p) {
					if p[pi] == '*' {
						pi++
					}
					break
				}
				childP := p[pi:]
				reCall := []bool{isMatch(s[si:], childP)}
				for si < len(s) {
					si++
					reCall = append(reCall, isMatch(s[si:], childP))
				}
				for _, ans := range reCall {
					if ans {
						return true
					}
				}
				return false
			}
			si++
			pi++
		}
	}
	for pi < len(p) {
		if pi+1 < len(p) && p[pi+1] == '*' {
			pi++
			for pi < len(p) {
				if p[pi] == '*' {
					pi++
				}
				break
			}
		} else {
			break
		}
	}
	return len(s)-si == len(p)-pi
}
