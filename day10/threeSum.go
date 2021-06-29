func threeSum(nums []int) [][]int {
	results := make([][]int, 0)

	sort.Ints(nums)
	pivot := 0
	size := len(nums)

	for pivot < size-2 {
		if pivot > 0 && nums[pivot] == nums[pivot-1] {
			pivot++
			continue
		}
		l := pivot + 1
		r := size - 1
		// loop for each pivot element
		for l < r {
			if nums[pivot]+nums[l]+nums[r] == 0 {
				results = append(results, []int{nums[pivot], nums[l], nums[r]})
			}
			//if pivot is greater than left + right, then moveleft pointer one step
			if nums[pivot]+nums[l]+nums[r] < 0 {
				currL := l
				for nums[l] == nums[currL] && l < r {
					l++
				}
			} else {
				currR := r
				for nums[r] == nums[currR] && l < r {
					r--
				}
			}
		}
		pivot++
	}

	return results
}
