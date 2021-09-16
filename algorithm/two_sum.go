package algorithm

// 1. brute force
// 2. map twice
// 3. map once
func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	var res []int
	var m = make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		var temp = target - v
		if index, ok := m[temp]; ok && index != i {
			return []int{i, index}
		}
	}
	return res
}

func twoSum3(nums []int, target int) []int {
	var res []int
	var m = make(map[int]int)
	for i, v := range nums {
		var temp = target-v
		if index,ok := m[temp] ; ok {
			return []int{index, i}
		}
		m[v] =i
	}
	return res
}