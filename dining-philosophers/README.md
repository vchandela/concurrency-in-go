### Dining Philosophers

- There are N philosophers and N forks. 
- - A philosopher has 2 actions:
    - think
    - eat: Each philosopher has to pick both forks to eat.

- M-1: Resource Hierarchy
    - Use a unique ID for each fork.
    - Always pick a fork with the lowest id first if it is available, otherwise, wait for your turn to eat.
        - Approach #1: Use a mutex to denote a fork
        - Approach #2: Use a channel of size 1 to denote a fork.
            - Use `chan struct{}` instead of `chan int/bool` as `struct{}` is of size 0. Using `struct{}` implies that we are using the channel for `signalling`.
- M-2: Use condition variable as an arbitrator
    - The mutex associated with condition variable will be a global lock.

- I have only implemented M-1.