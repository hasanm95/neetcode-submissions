func twoSum(nums []int, target int) []int {
    seenMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, found := seenMap[complement]; found {
            return []int{index, i}
        } else {
            seenMap[num] = i
    }
    }
    return nil
}
