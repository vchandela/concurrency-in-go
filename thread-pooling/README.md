### Thread pooling in Go
For this, the main thread shares buffered channel(s) with worker threads.
In this example, we have 3 workers and 5 tasks.

For a single worker thread, using a boolean `done` channel is preferred.
For >1 workers, using WaitGroups is better if we don't want to receive any results. 

#### Key Concepts
- Iterating over a channel using for-loop is blocking by default.
-  It is possible to close a non-empty channel but still have the remaining values be received.
