package ex1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if idx, exist := m[target-num]; exist {
			return []int{idx, i}
		}

		m[num] = i
	}

	return []int{}
}
