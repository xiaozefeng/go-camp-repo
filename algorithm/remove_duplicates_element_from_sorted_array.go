package algorithm

// 快慢指针
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	var slow = 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
