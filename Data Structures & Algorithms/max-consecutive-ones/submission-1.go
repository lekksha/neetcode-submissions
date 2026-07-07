func findMaxConsecutiveOnes(nums []int) int {
    var count, max int = 0, 0
	for _, value := range nums {
        if value == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
    }
	return max
}
