# {name}

Benchmarks including hash algorithms for non-cryptographic purposes.

```text
BenchmarkAkyotoHash_4B-12       1000000000               0.705 ns/op           0 B/op          0 allocs/op
BenchmarkCRC32_4B-12            732247504                1.60 ns/op            0 B/op          0 allocs/op
BenchmarkXXHash_4B-12           697969192                1.70 ns/op            0 B/op          0 allocs/op
BenchmarkSipHash_4B-12          831329154                1.43 ns/op            0 B/op          0 allocs/op

BenchmarkAkyotoHash_8B-12       1000000000               0.520 ns/op           0 B/op          0 allocs/op
BenchmarkCRC32_8B-12            548619034                2.17 ns/op            0 B/op          0 allocs/op
BenchmarkXXHash_8B-12           670126293                1.77 ns/op            0 B/op          0 allocs/op
BenchmarkSipHash_8B-12          680157788                1.75 ns/op            0 B/op          0 allocs/op

BenchmarkAkyotoHash_10KB-12     30696001                37.0 ns/op             0 B/op          0 allocs/op
BenchmarkCRC32_10KB-12          22894020                50.3 ns/op             0 B/op          0 allocs/op
BenchmarkXXHash_10KB-12         11611434               101 ns/op               0 B/op          0 allocs/op
BenchmarkSipHash_10KB-12         2812972               424 ns/op               0 B/op          0 allocs/op
```

{go:footer}
