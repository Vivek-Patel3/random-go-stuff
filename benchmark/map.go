package benchmark

import "strconv"

const n = 50000000

var sinkMap map[string]int

func makeKeys(n int) []string {
	keys := make([]string, n)
	for i := range keys {
		keys[i] = "key_" + strconv.Itoa(i)
	}
	return keys
}

// making the map without any capacity pre-given
func InsertKeysInMap(keys []string) {
	m := make(map[string]int)

	for i,v := range keys {
		m[v] = i
	}

	sinkMap = m
}

// pre-giving the capacity
func InsertKeysInMapCap(keys []string) {
	m := make(map[string]int, n)

	for i, v := range keys {
		m[v] = i
	}

	sinkMap = m
}