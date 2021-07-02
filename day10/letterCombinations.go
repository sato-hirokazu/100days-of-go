func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	num2chars := [...]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := strings.Split(num2chars[uint8(digits[0]) - '2'], "")
	for _, v := range digits[1:] {
		temp := make([]string, 0)
		for _, c := range num2chars[uint8(v) - '2'] {
			for _, prev := range res {
				temp = append(temp, prev+string(c))
			}
		}
		res = temp
	}
	return res
}
