package benchmark

import (
	"testing"
	"fmt"
)

func BenchmarkMapInsert(b *testing.B) {
	for _, n := range []int{10000, 100000, 500000, 1000000, 5000000} {
		n := n

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			keys := makeKeys(n)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				m := make(map[string]int)
				for j, k := range keys {
					m[k] = j
				}
				sinkMap = m
			}
		})
	}
}

func BenchmarkMapInsertCap(b *testing.B) {
	for _, n := range []int{10000, 100000, 500000, 1000000, 5000000} {
		n := n
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			keys := makeKeys(n)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				m := make(map[string]int, n)
				for j, k := range keys {
					m[k] = j
				}
				sinkMap = m
			}
		})
	}
}