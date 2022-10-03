```bash
go get -u github.com/labi-le/fast-int-conversion 
```

### Benchmarks
```
cpu: AMD Ryzen 5 4500U with Radeon Graphics         

BenchmarkStrToInt-6             155217518                7.983 ns/op           0 B/op          0 allocs/op
BenchmarkStrToInt64-6           83933821                14.53 ns/op            0 B/op          0 allocs/op
BenchmarkStrToUint-6            159441013                7.553 ns/op           0 B/op          0 allocs/op
BenchmarkStrToUint64-6          79895065                14.94 ns/op            0 B/op          0 allocs/op
BenchmarkStrToUintptr-6         82947150                14.97 ns/op            0 B/op          0 allocs/op

```