func feedChannel(feeder string, inp, out chan string) {
	for s := range inp {
		for _, r := range feeder {
			out <- string(r) + s
		}
	}
	close(out)
}


func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
	m := map[rune]string{
		'2' : "abc",
		'3' : "def",
		'4' : "ghi",
		'5' : "jkl",
		'6' : "mno",
		'7' : "pqrs",
		'8' : "tuv",
		'9' : "wxyz",
	}
	rc := make(chan string)
	nc := rc
	for _,r := range digits {
		N := make(chan string)
		s := m[r]
		go feedChannel(s,N,rc)
		rc = N
	}
	rc <- ""
	close(rc)
	var ret []string
	for r := range nc {
		ret = append(ret, r)
	}
	return ret
}
