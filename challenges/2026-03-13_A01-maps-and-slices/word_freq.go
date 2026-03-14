package main

func WordFrequency(inp []string) map[string]int {
	seen := make(map[string]int)
	for _, word := range inp {
		key, exists := seen[word]
		if exists {
			seen[word] = key + 1
		} else {
			seen[word] = 1
		}
	}
	return seen
}
