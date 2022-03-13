package main

import (
	"fmt"
	"sort"
)

func main() {
	counts := map[string]int{"Barry": 13, "Clean": 98, "Lisi": 99}
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	// 从小到大
	sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] < counts[keys[j]] })
	// 从大到小
	// sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] > counts[keys[j]] })
	for _, key := range keys {
		fmt.Printf("%s, %d\n", key, counts[key])
	}
}
