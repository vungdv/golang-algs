package main

import "testing"

// func BenchmarkGenerate2(b *testing.B) {
// 	generateTrees2(16)
// }

func BenchmarkGenerateNoCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// It's not fair but it showed that cache in a right way is very valuable.
		generateTreesNoCache(8)
	}
}
func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateTreesCache(8)
	}
}

func BenchmarkGenerateNoCacheNoLoop(b *testing.B) {

	generateTreesNoCache(8)
}
func BenchmarkGenerateNoLoop(b *testing.B) {
	generateTreesCache(8)
}

/*
goos: darwin
goarch: arm64
pkg: generateBST
cpu: Apple M1 Pro
BenchmarkGenerateNoCache-8          2884            364064 ns/op
BenchmarkGenerate-8             93910182                12.58 ns/op
PASS
ok      generateBST     2.871s


goos: darwin
goarch: arm64
pkg: generateBST
cpu: Apple M1 Pro
BenchmarkGenerateNoCache-8                  2613            383020 ns/op
BenchmarkGenerate-8                     94358472                13.39 ns/op
BenchmarkGenerateNoCacheNoLoop-8        1000000000               0.0003645 ns/op
BenchmarkGenerateNoLoop-8               1000000000               0.0000002 ns/op
*/
