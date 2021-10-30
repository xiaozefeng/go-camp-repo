package algorithm

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	var n = len(nums)
	if nums == nil || n < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			var l, r = i + 1, n - 1
			for l < r {
				var sum = nums[l] + nums[r] + nums[i]
				if sum == 0 {
					ans = append(ans, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--

				} else if sum < 0 {
					l++
				} else {
					r--
				}

			}

		}
	}
	return ans
}
