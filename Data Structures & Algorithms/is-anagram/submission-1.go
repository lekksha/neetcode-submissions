func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var gottem = map[rune]int{}
	for _, v := range(s) {
		gottem[v]++
	}
	for _, v := range(t) {
		gottem[v]--
	}
	for _, value := range gottem {
		if value != 0 {
			return false
		}
	}
	return true
}
