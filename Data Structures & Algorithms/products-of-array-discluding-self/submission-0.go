func productExceptSelf(nums []int) []int {
	prefixProd := make([]int, len(nums))
	postfixProd := make([]int, len(nums))
	curPrefixProd, curPostfixProd := 1, 1
	for i, _ := range(nums) {
		curPrefixProd *= nums[i]
		prefixProd[i] = curPrefixProd

		curPostfixProd *= nums[len(nums)-1-i]
		postfixProd[len(nums)-1-i] = curPostfixProd
	}

	var result []int = make([]int, len(nums))
	for i := range(nums) {
		var prevV, nextV = 1, 1
		if i > 0 {
			prevV = prefixProd[i-1]
		}
		if i < len(nums)-1 {
			nextV = postfixProd[i+1]
		}
		result[i] = prevV * nextV
	}
	fmt.Println(result)
	return result
}
