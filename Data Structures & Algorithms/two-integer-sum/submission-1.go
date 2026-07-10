func twoSum(nums []int, target int) []int {
    prevElems := map[int]int{}
    for i, value := range nums {
        prevElemIndex, ok := prevElems[target-value]
        if ok {
            return []int{prevElemIndex, i}
        } else {
            prevElems[value] = i
        }
    }
    return []int{}
}
