func hasDuplicate(nums []int) bool {
	numMap := make(map[int]int)

	for i, num := range nums {
		_, exists := numMap[num]
		if exists {
			return true
		} else {
			numMap[num] = i
		}
	}
	return false
}
