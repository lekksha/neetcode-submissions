func hasDuplicate(nums []int) bool {
    var dict = map[int]int{}
	for _, v := range nums {
		dict[v]++
		if dict[v] > 1 {
			return true
		}
	}
	return false
}
