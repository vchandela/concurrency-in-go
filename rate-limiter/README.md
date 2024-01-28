### Rate Limiter
- Go supports rate limiting with goroutines, channels, and tickers.
- Naive: 1 request every 200 milliseconds by blocking on a receive from the `limiter` channel
- Bursty: We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our `limiter` channel.