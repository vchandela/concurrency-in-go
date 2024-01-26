### Matrix multiplication

- Benchmark Results
```
goos: darwin
goarch: amd64
pkg: matmul/utils
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMatMul-12              1000000000               0.04260 ns/op
BenchmarkMatMulMutex-12         1000000000               0.08166 ns/op
BenchmarkMatMulRWMutex-12       1000000000               0.1076 ns/op
PASS
ok      matmul/utils    2.401s
```

- In matrix multiplication,  each individual element of the result matrix is updated independently in the loop.
- There is no sharing of data between rows.
- So, Mutex/RWMutex and WaitGroup slow the program down due to more synchronization overhead.