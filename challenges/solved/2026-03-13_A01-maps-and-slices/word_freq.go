package main

func WordFrequency(inp []string) map[string]int {
	seen := make(map[string]int)
	for _, word := range inp {
		seen[word]++
	}
	return seen
}
