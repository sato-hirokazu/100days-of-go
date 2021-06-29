func lengthOfLongestSubstring(s string) int {
    m := make(map[rune]int)
    i := 0
    maxLen := 0
    ok := false // this asved 4 ms for total run
    ss := []rune(s)
    for j := 0; j < len(ss); j++ {
        _, ok = m[ss[j]]
        if ok { 
           i =  Max(m[ss[j]], i)
        }
        maxLen = Max(maxLen, j-i+1)
        m[ss[j]] = j + 1
    }
    return maxLen
}

func Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

