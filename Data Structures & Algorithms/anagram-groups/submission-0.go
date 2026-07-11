func groupAnagrams(strs []string) [][]string {
	symbolsToGroup := map[[26]int][]string{}
	for _, word := range strs {
		symbolCounter := [26]int{}
		for _, symbol := range word {
			symbolCounter[symbol-'a']++
		}
		symbolsToGroup[symbolCounter] = append(symbolsToGroup[symbolCounter], word)
	}

	anagramGroups := [][]string{}
	for _, group := range symbolsToGroup {
		anagramGroups = append(anagramGroups, group)
	}

	return anagramGroups
}