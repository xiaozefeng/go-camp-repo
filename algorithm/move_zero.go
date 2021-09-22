package algorithm

// 遍历数组，遇到0 的元素就跟数组末尾的元素交换
func moveZeroes(nums []int) {

	var nonZero = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZero], nums[i] = nums[i], nums[nonZero]
			nonZero++
		}
	}
}
