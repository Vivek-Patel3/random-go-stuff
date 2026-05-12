package benchmark

import "testing"

func BenchmarkInsertKeysInMap(b *testing.B) {
	keys := makeKeys(n)

	for b.Loop() {
		InsertKeysInMap(keys)
	}
}

func BenchmarkInsertKeysInMapCap(b *testing.B) {
	keys := makeKeys(n)

	for b.Loop() {
		InsertKeysInMap(keys)
	}	
}