### Barrier in go
- Currently, Go doesn't have a Barrier construct like Python/Java
- We need 4 things to create a barrier:
  - total: total no. of threads participating in the barrier
  - count: no. of threads that still have to call Wait()
  - condition Var: allows a thread to wait for other threads and wake up when barrier is unblocked
    - It is bound to a mutex that protects `total` and `count` from multiple threads
- As blue thread is last to arrive, it will release all other threads that are waiting on condition variable. Output looks like the following:
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

#### Learnings:
- In case of conditional variable, the associated lock must be acquired before a thread can invoke Wait()/Signal()/Broadcast() on the condition variable.
- When a thread invokes Wait(), it simultaneously gives up the lock associated with the condition variable. Only when the sleeping thread wakes up again on a Signal()/Broadcast(), will it reacquire the lock.
- A peculiarity of condition variables is the possibility of `spurious wakeups`. It means that a thread might wakeup as if it has been signaled even though nobody called Signal()/Broadcast() on the condition variable in question.
  -  The thread must test the conditions again for validity before moving forward.
  -  We must always check for conditions in a loop and Wait() inside it. 