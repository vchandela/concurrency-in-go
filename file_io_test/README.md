## IO Test
- Program to test file I/O using different combinations of Batch Size (`max_poll_records`), Channel Size (`channel_size`) and Goroutine Workers (`num_workers`).
- Flow:
  - We fetch records from a file in a batch of size `max_poll_records`.
  - Put these records one by one into a channel of size `channel_size`. This is done in `ioTest()` and not `ioTest2()`
  - `num_workers` goroutines will process the records for 2 ms each.

### Observation for ioTest() with temp100K (100k records)
1. Keeping Batch size and Channel size same, time taken by program decreases as Goroutine Workers increases.
2. No noticeable change if Goroutine Workers are constant.
3. No noticeable change on changin GOMAXPROCS.
4. Time decreases as we increase Batch size = Channel size = Goroutine workers from 1024 to 4096 

| Benchmark                                 | ns/op         | ms/op       |
|-------------------------------------------|---------------|-------------|
| BenchmarkIoTest_1024_1024_32-10           | 7105128917    | 7105.13     |
| BenchmarkIoTest_1024_1024_64-10           | 3553434167    | 3553.43     |
| BenchmarkIoTest_1024_1024_128-10          | 1782262042    | 1782.26     |
| BenchmarkIoTest_1024_1024_256-10          | 1057335959    | 1057.34     |
| BenchmarkIoTest_1024_1024_512-10          | 568953333     | 568.95      |
| BenchmarkIoTest_1024_1024_1024-10         | 226877383     | 226.88      |
| BenchmarkIoTest_2048_2048_32-10           | 7219402584    | 7219.40     |
| BenchmarkIoTest_2048_2048_64-10           | 3520654583    | 3520.65     |
| BenchmarkIoTest_2048_2048_128-10          | 1765975334    | 1765.98     |
| BenchmarkIoTest_2048_2048_256-10          | 882110938     | 882.11      |
| BenchmarkIoTest_2048_2048_512-10          | 443025917     | 443.03      |
| BenchmarkIoTest_2048_2048_1024-10         | 218604358     | 218.60      |
| BenchmarkIoTest_2048_2048_2048-10         | 116059069     | 116.06      |
| BenchmarkIoTest_4096_4096_32-10           | 7079761833    | 7079.76     |
| BenchmarkIoTest_4096_4096_64-10           | 3515230833    | 3515.23     |
| BenchmarkIoTest_4096_4096_128-10          | 1774125000    | 1774.13     |
| BenchmarkIoTest_4096_4096_256-10          | 884217084     | 884.22      |
| BenchmarkIoTest_4096_4096_512-10          | 443581083     | 443.58      |
| BenchmarkIoTest_4096_4096_1024-10         | 217451417     | 217.45      |
| BenchmarkIoTest_4096_4096_2048-10         | 113388884     | 113.39      |
| BenchmarkIoTest_4096_4096_4096-10         | 64599449      | 64.60       |
| BenchmarkIoTest_2048_1024_32-10           | 7041002958    | 7041.00     |
| BenchmarkIoTest_2048_1024_64-10           | 3539475667    | 3539.48     |
| BenchmarkIoTest_2048_1024_128-10          | 1867630042    | 1867.63     |
| BenchmarkIoTest_2048_1024_256-10          | 883694688     | 883.69      |
| BenchmarkIoTest_2048_1024_512-10          | 441035333     | 441.04      |
| BenchmarkIoTest_2048_1024_1024-10         | 218317092     | 218.32      |
| BenchmarkIoTest_4096_2048_32-10           | 7039505417    | 7039.51     |
| BenchmarkIoTest_4096_2048_64-10           | 3519979334    | 3520.00     |
| BenchmarkIoTest_4096_2048_128-10          | 1761209417    | 1761.21     |
| BenchmarkIoTest_4096_2048_256-10          | 880956020     | 880.96      |
| BenchmarkIoTest_4096_2048_512-10          | 441660014     | 441.66      |
| BenchmarkIoTest_4096_2048_1024-10         | 217538342     | 217.54      |
| BenchmarkIoTest_4096_2048_2048-10         | 113140597     | 113.14      |
| BenchmarkIoTest_4096_1024_32-10           | 7030176167    | 7030.18     |
| BenchmarkIoTest_4096_1024_64-10           | 3525375542    | 3525.38     |
| BenchmarkIoTest_4096_1024_128-10          | 1759467083    | 1759.47     |
| BenchmarkIoTest_4096_1024_256-10          | 881303584     | 881.30      |
| BenchmarkIoTest_4096_1024_512-10          | 454045430     | 454.05      |
| BenchmarkIoTest_4096_1024_1024-10         | 260778750     | 260.78      |
| BenchmarkIoTest_1024_2048_32-10           | 7055822291    | 7055.82     |
| BenchmarkIoTest_1024_2048_64-10           | 3680074792    | 3680.07     |
| BenchmarkIoTest_1024_2048_128-10          | 1801862875    | 1801.86     |
| BenchmarkIoTest_1024_2048_256-10          | 896663042     | 896.66      |
| BenchmarkIoTest_1024_2048_512-10          | 437521736     | 437.52      |
| BenchmarkIoTest_1024_2048_1024-10         | 219238342     | 219.24      |
| BenchmarkIoTest_1024_2048_2048-10         | 110461762     | 110.46      |
| BenchmarkIoTest_1024_4096_32-10           | 7049948292    | 7049.95     |
| BenchmarkIoTest_1024_4096_64-10           | 3521211750    | 3521.21     |
| BenchmarkIoTest_1024_4096_128-10          | 1758892250    | 1758.89     |
| BenchmarkIoTest_1024_4096_256-10          | 880754708     | 880.75      |
| BenchmarkIoTest_1024_4096_512-10          | 441623542     | 441.62      |
| BenchmarkIoTest_1024_4096_1024-10         | 220434775     | 220.43      |
| BenchmarkIoTest_1024_4096_2048-10         | 109599250     | 109.60      |
| BenchmarkIoTest_1024_4096_4096-10         | 63002583      | 63.00       |
| BenchmarkIoTest_2048_4096_32-10           | 7038179500    | 7038.18     |
| BenchmarkIoTest_2048_4096_64-10           | 3522561458    | 3522.56     |
| BenchmarkIoTest_2048_4096_128-10          | 1765253500    | 1765.25     |
| BenchmarkIoTest_2048_4096_256-10          | 892491125     | 892.49      |
| BenchmarkIoTest_2048_4096_512-10          | 440833097     | 440.83      |
| BenchmarkIoTest_2048_4096_1024-10         | 220466733     | 220.47      |
| BenchmarkIoTest_2048_4096_2048-10         | 110565767     | 110.57      |
| BenchmarkIoTest_2048_4096_4096-10         | 60058430      | 60.06       |

### Observation for ioTest3() with temp100K (100k records)
1. This is slightly slower than ioTest() but doesn't use channels. When we pass structs in channel, they are copied.
So, we can save some memory by removing channels.

| Benchmark                                     | ns/op        | ms/op       |
|-----------------------------------------------|--------------|-------------|
| BenchmarkIoTest_1024_32-10                    | 7072599042   | 7072.599042 |
| BenchmarkIoTest_1024_64-10                    | 3543012958   | 3543.012958 |
| BenchmarkIoTest_1024_128-10                   | 1766818083   | 1766.818083 |
| BenchmarkIoTest_1024_256-10                   | 882565458    | 882.565458  |
| BenchmarkIoTest_1024_512-10                   | 438003389    | 438.003389  |
| BenchmarkIoTest_1024_1024-10                  | 228796983    | 228.796983  |
| BenchmarkIoTest_2048_32-10                    | 7083157916   | 7083.157916 |
| BenchmarkIoTest_2048_64-10                    | 3540968958   | 3540.968958 |
| BenchmarkIoTest_2048_128-10                   | 1774568459   | 1774.568459 |
| BenchmarkIoTest_2048_256-10                   | 888086708    | 888.086708  |
| BenchmarkIoTest_2048_512-10                   | 448281820    | 448.281820  |
| BenchmarkIoTest_2048_1024-10                  | 220163033    | 220.163033  |
| BenchmarkIoTest_2048_2048-10                  | 123541083    | 123.541083  |
| BenchmarkIoTest_4096_32-10                    | 7238411291   | 7238.411291 |
| BenchmarkIoTest_4096_64-10                    | 3931431292   | 3931.431292 |
| BenchmarkIoTest_4096_128-10                   | 1815033041   | 1815.033041 |
| BenchmarkIoTest_4096_256-10                   | 880407875    | 880.407875  |
| BenchmarkIoTest_4096_512-10                   | 442479445    | 442.479445  |
| BenchmarkIoTest_4096_1024-10                  | 219957242    | 219.957242  |
| BenchmarkIoTest_4096_2048-10                  | 113393196    | 113.393196  |
| BenchmarkIoTest_4096_4096-10                  | 88920198     | 88.920198   |

### Note
- ioTest2() should only be used for cron jobs and not jobs running 24X7. It implements `TCP slow start`.

### Machine Specs
```
goos: darwin
goarch: amd64
pkg: io_test/cmd
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
GOMAXPROCS = 4
```