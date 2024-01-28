### Multithreaded web-crawler
- Solution to [Exercise: Web Crawler](https://go.dev/tour/concurrency/10)

- We will use a map to keep track of visited web pages.
- Why do we need to keep a RWMutex{}?
  - Let's say we only use a Mutex for updating urlMap[url] = true. There is no lock anywhere else.
        - e.g webpage A -> webpages B, C and D; C -> B and D -> B
        - If we acquire writer's lock for B and simultaneously C checks urlMap[B] and it is false. So, it goes in and prints it.
        - B also writes urlMap[B] with true once more -> we get B 2 times.
        - As a result, we need a reader's lock while checking urlMap[B]