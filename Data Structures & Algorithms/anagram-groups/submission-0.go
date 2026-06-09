func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	
	result := [][]string{}
	visited := make([]bool, len(strs))

	for i, val := range strs {
		if visited[i] {
			continue
		}
		group := []string{val}
		sortedVal := sortStringCharacters(val)
		for j := i + 1; j < len(strs); j++ {
			if visited[j] {
				continue
			}
			sortedVal2 := sortStringCharacters(strs[j])
			if sortedVal == sortedVal2 {
				group = append(group, strs[j])
				visited[j] = true
			}
		}
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
