func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	
	res := make(map[string][]string)

	for _, s := range strs {
		sortedS := sortStringCharacters(s)
		res[sortedS] = append(res[sortedS], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}

	return result
}

func sortStringCharacters(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
