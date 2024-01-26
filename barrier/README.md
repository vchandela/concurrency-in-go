### Barrier in go
- Currently, Go doesn't have a Barrier construct like Python/Java
- We need 4 things to create a barrier:
  - total: total no. of threads participating in the barrier
  - count: no. of threads that still have to call Wait()
  - mutex: protect `total` and `count` from multiple threads
  - condition Var: allows a thread to wait for other threads and wake up when barrier is unblocked
- As blue thread is last to arrive, it will release all other threads that are waiting on condition variable
    ```
    green running
    red running
    blue running
    red waiting on condition
    green waiting on condition
    blue releases all waiting threads
    blue running
    green running
    red running
    ```