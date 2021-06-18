func myAtoi(str string) int {
    

    const MaxInt = 2147483647 
    const MinInt = -2147483648
    
    str = strings.TrimLeft(str, " ")
    
    if len(str) == 0 {
        return 0
    }
    
    res := 0
    
    runes := []rune(str)
    
    isNegative := false
    
    if runes[0] == '-' {
       isNegative = true
        runes = runes[1:]
    } else if runes[0] == '+' {
        runes = runes[1:]
    } 
    
    if len(runes) == 0 {
        return 0
    }
    
    if runes[0] - '0' < 0 || runes[0] - '0' > 9 {
        return 0
    }
    
    for _, v := range runes {
        
        currNum := int(v - '0')
        
        if currNum < 0 || currNum > 9 {
            break
        } else {
            
            validRes := int64(res) * int64(10) + int64(currNum)
            
            if validRes > int64(MaxInt) && !isNegative {
                return MaxInt
            } else if validRes * int64(-1) < int64(MinInt) {
                return MinInt
            }
            
            res = res * 10 + currNum
        }
    }
    
    if isNegative {
        res = res * -1
    }
    
    return res
    
}
