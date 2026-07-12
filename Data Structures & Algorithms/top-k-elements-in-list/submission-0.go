func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	freq := make([][]int, len(nums)+1)

	for _, number := range nums {
		count[number]++
	}
	
	for number, timesMet := range count {
		freq[timesMet] = append(freq[timesMet], number)
	}
	
	res := []int{}
	for i:= len(nums); i > 0; i-- {
		for j := 0; j < len(freq[i]); j++ {
			res = append(res, freq[i][j])
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
