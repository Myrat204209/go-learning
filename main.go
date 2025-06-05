package main

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if len(name)==0 {
			continue
		}
		runes := []rune(name)
		first:= rune(runes[0])
		if _, ok := counts[first]; !ok {
			counts[first] = make(map[string]int)
		}
		counts[first][name] = counts[first][name]+1
	}
	return counts
}
