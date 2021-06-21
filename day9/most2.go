func maxArea(height []int) int {
    
    if len(height)==0{
        return 0
    }
    
    var(
        max = 0
        left = 0
        right = len(height)-1
        curr = 0
    )
    
    for left<right{
        //adjust maxArea
        curr = getArea(left, right, height)
        if curr>max{
            max=curr
        }
        
        //move on
        if height[left]<height[right]{
            left++
        }else{
            right--   
        }
    }
    return max
}

// calculate the area 
func getArea(left, right int, height []int) int{
    if height[left]> height[right] {
        return height[right]*(right-left)
    }    
    return height[left]*(right-left)
}

