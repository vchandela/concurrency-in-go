### Bounded Buffer
- We have a limited size buffer which can have items added to it or removed from it by different producer and consumer threads
- A blocking queue is defined as a queue which blocks the caller of the enqueue method if there's no more capacity to add the new item being enqueued
- Similarly, the queue blocks the dequeue caller if there are no items in the queue. 
- Also, the queue notifies a blocked enqueuing thread when space becomes available and a blocked dequeuing thread when an item becomes available in the queue.

- M-1: conditional variable
  - Create implementation of queue.

- M-2: channels
  - Use atomic operations to update size of channel. 

#### Learnings
- Do we need `Signal()`/`Broadcast()` in our implementation?
    - Consider a situation with two producer threads and one consumer thread all working with a queue of size one. It's possible that when an item is added to the queue by one of the producer threads, the other two threads are blocked waiting on the condition variable. If the producer thread after adding an item invokes `Signal()` it is possible that the other producer thread is chosen by the system to resume execution. The woken-up producer thread would find the queue full and go back to waiting on the condition variable, causing a deadlock. Invoking `Broadcast()` assures that the consumer thread also gets a chance to wake up and resume execution.

- What about non-blocking implementation?
    - If a consumer or a producer can successfully enqueue or dequeue an item, it is considered non-blocking. However, if the queue is full or empty then a producer or a consumer (respectively) need not wait until the queue can be added to or taken from.
    - Trivial Solution:  Return a boolean value indicating the success of an operation. If the invoker of either Enqueue() or Dequeue() receives False, then it is the responsibility of the invoker to retry the operation at a later time.