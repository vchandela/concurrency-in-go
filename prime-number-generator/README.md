### Prime-number generator (Consumer-Producer)
- Imagine a scenario where we have two threads working together to find prime numbers and print them. 
  - Say the producer thread finds the prime number and the consumer thread is responsible for printing the found prime. 
  - The producer thread sets a boolean flag `found_prime` whenever it determines an integer is a prime number. 
  - The consumer thread needs to know when the producer thread has hit upon a prime number. 
- Naive Approach: Have the consumer thread do a `busy wait and keep polling` for the boolean value.
  - Con: Consumer thread is busy waiting which wastes CPU cycles.
- Approach #2: Use a conditional variable on `found_prime`
  - Con: We must always check for conditions in a loop and Wait() inside it.
    - This is slow.
- Approach #3: Use a semaphore
  - Acquire(): blocks if there are no more buffers to consume in the channel
  - Release(n): release a buffer in the channel


#### Learnings
- Use Interface to avoid exposing your structs to outside packages. Abstraction is the key!!
- In case of conditional variable, the associated lock must be acquired before a thread can invoke Wait()/Signal()/Broadcast() on the condition variable.
- When a thread invokes Wait(), it simultaneously gives up the lock associated with the condition variable. Only when the sleeping thread wakes up again on a Signal()/Broadcast(), will it reacquire the lock.
- A peculiarity of condition variables is the possibility of `spurious wakeups`. It means that a thread might wakeup as if it has been signaled even though nobody called Signal()/Broadcast() on the condition variable in question.
  -  The thread must test the conditions again for validity before moving forward.
  -  We must always check for conditions in a loop and Wait() inside it. 
- `Mutex allows a single thread exclusive access` to a critical section while `semaphore allows atmost N threads exclusive access`. So, semaphore is a more generalized version of a mutex.
- Go doesn't have built-in semaphore implementation but it can be emulated via `buffered channel`.
  - When buffered channel is full, the channel will lock the goroutine and the thread waits until a buffer is available.
  - Whenever you hear `bounded concurrency`, think of semaphore.
- The struct{} type has zero size, which means it doesn't consume any memory.
  - This is useful when you don't need any data associated with the channel element and you're using the channel purely for signaling purposes.