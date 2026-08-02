func isPalindrome(s string) bool {
	var filtered []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}
	var i int = 0
	for i < len(filtered)/2 {
		if filtered[i] == filtered[len(filtered)-1-i]{
			i++
		} else {
			return false
		}
	}
	return true
}

