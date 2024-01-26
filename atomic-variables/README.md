### Letter frequency program
- We download 200 files and count frequency of each english alphabet
- The results are as follows:
  ```
    Single-threaded Elapsed time: 20.588132676s
    Multi-threaded with Lock Elapsed time: 3.374354584s
    Multi-threaded with Atomic Elapsed time: 1.567141525s
  ```

- Benchmark results
  ```
    goos: darwin
    goarch: amd64
    pkg: atomic/cmd
    cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
    BenchmarkST-12                         1        21409890897 ns/op
    BenchmarkMTLock-12              1000000000               0.0007034 ns/op
    BenchmarkMTAtomic-12            1000000000               0.0009571 ns/op
    PASS
    ok      atomic/cmd      23.169s
  ```

- We can conclude that atomic variables are faster than lock/mutex in this case.