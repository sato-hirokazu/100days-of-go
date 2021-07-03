func fourSum(nums []int, target int) (ret [][]int) {
	lnums := len(nums)
	if lnums < 4 {
		return nil
	}

	sort.Ints(nums)

	for i := 0; i < lnums-3; i++ {
		if 4*nums[i] > target {
			return
		}
		if nums[i]+nums[lnums-1]+nums[lnums-2]+nums[lnums-3] < target {
			continue
		}

		smallTarget := target - nums[i]
		if threes := threeSumHelper(nums[i+1:], smallTarget); len(threes) != 0 {
			for _, three := range threes {
				ret = append(ret, append([]int{nums[i]}, three...))
			}
		}
		for i < lnums-3 && nums[i+1] == nums[i] {
			i++
		}
	}
	return
}

func threeSumHelper(nums []int, target int) (ret [][]int) {
	lnums := len(nums)
	if lnums < 3 {
		return nil
	}

	for i := 0; i < lnums-2; i++ {
		if 3*nums[i] > target {
			return
		}
		if nums[i]+nums[lnums-1]+nums[lnums-2] < target {
			continue
		}

		j, k := i+1, lnums-1

		for j < k {
			a, b, c := nums[i], nums[j], nums[k]
			sum := a + b + c
			if sum == target {
				ret = append(ret, []int{a, b, c})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > target {
				k--
			} else {
				j++
			}
		}

		for i < lnums-2 && nums[i+1] == nums[i] {
			i++
		}

	}
	return
}
