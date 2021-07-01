func threeSumClosest(tmp_nums []int, target int) int {
    var nums ArrInt = tmp_nums
    sort.Sort(nums)
    
    n := len(nums)
    
    result := nums[n-1] + nums[n-2] + nums[n-3]
    smallest := abs(result - target)
        
    k, s := 0, 0
    
    for i:=0; i < n-2; i++ {
        k = n - 1
        for j := i+1; j < k; j++ {
            for j < k {
                if s = nums[i] + nums[j] + nums[k] - target; s >= 0 {
                    if smallest > s {
                        smallest, result = s, nums[i] + nums[j] + nums[k]
                    }
                    k--
                } else {
                    break
                }
            }
        }
    }
    
    for i := n-1; i>=0; i-- {
        k = 0
        for j := i-1; j > k; j-- {
            for j > k {
                if s = target - nums[i] - nums[j] - nums[k]; s >= 0 {
                    if smallest > s {
                        smallest, result = s, nums[i] + nums[j] + nums[k]
                    }
                    k++
                } else {
                    break
                }
            }
        }
    }
    
    return result
}

func abs(i int) int {
    if i < 0 {
        return -i
    }
    
    return i
}

type ArrInt []int

func (a ArrInt) Len() int           { return len(a) }
func (a ArrInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArrInt) Less(i, j int) bool { return a[i] < a[j] }
