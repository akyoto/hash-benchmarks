package benchmark_test

import (
	"bytes"
	"hash/crc32"
	"testing"

	"github.com/OneOfOne/xxhash"
	"github.com/akyoto/hash"
	"github.com/dchest/siphash"
)

var (
	data4B   = []byte("1234")
	data8B   = []byte("12345678")
	data10KB = bytes.Repeat([]byte("HelloWorld"), 1000)
)

func BenchmarkAkyotoHash_4B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.Bytes(data4B)
		}
	})
}

func BenchmarkCRC32_4B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			crc32.ChecksumIEEE(data4B)
		}
	})
}

func BenchmarkXXHash_4B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		h := xxhash.New64()

		for pb.Next() {
			h.Reset()
			_, _ = h.Write(data4B)
			h.Sum64()
		}
	})
}

func BenchmarkSipHash_4B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			siphash.Hash(13, 10, data4B)
		}
	})
}

func BenchmarkAkyotoHash_8B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.Bytes(data8B)
		}
	})
}

func BenchmarkCRC32_8B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			crc32.ChecksumIEEE(data8B)
		}
	})
}

func BenchmarkXXHash_8B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		h := xxhash.New64()

		for pb.Next() {
			h.Reset()
			_, _ = h.Write(data8B)
			h.Sum64()
		}
	})
}

func BenchmarkSipHash_8B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			siphash.Hash(13, 10, data8B)
		}
	})
}

func BenchmarkAkyotoHash_10KB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.Bytes(data10KB)
		}
	})
}

func BenchmarkCRC32_10KB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			crc32.ChecksumIEEE(data10KB)
		}
	})
}

func BenchmarkXXHash_10KB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		h := xxhash.New64()

		for pb.Next() {
			h.Reset()
			_, _ = h.Write(data10KB)
			h.Sum64()
		}
	})
}

func BenchmarkSipHash_10KB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			siphash.Hash(13, 10, data10KB)
		}
	})
}
