# cachify

> A high-performance reverse proxy built completely from scratch in Go using raw TCP sockets.

Unlike traditional reverse proxies that rely on Go's `net/http` package, ProxyForge is implemented directly on top of the `net` package to gain a deeper understanding of HTTP, TCP, concurrency, and systems programming.

---

## Features

- Raw TCP socket implementation
- Custom HTTP request parser
- Reverse proxy
- Connection pooling
- Persistent TCP connections (Keep-Alive)
- Concurrent request handling using goroutines
- Efficient memory management
- Designed for high throughput and low latency

---

## Why?

This project was built to understand what happens underneath Go's HTTP stack.

Instead of relying on high-level abstractions, ProxyForge manually handles:

- TCP connections
- HTTP parsing
- Request forwarding
- Response streaming
- Connection reuse
- Concurrent clients

The goal was educational first and performance second.

---

## Architecture

```
             Client
                │
                ▼
      ┌─────────────────┐
      │   ProxyForge    │
      │                 │
      │ HTTP Parser     │
      │ Connection Pool │
      │ Reverse Proxy   │
      └─────────────────┘
                │
                ▼
         Origin Server
```

---

## Tech Stack

- Go
- TCP Sockets
- Goroutines
- Channels
- Synchronization Primitives

---

## Benchmarks

Benchmarks were performed against a simple backend server and compared with Caddy.

| Server | Requests/sec |
|---------|-------------:|
| Backend | XXXXX |
| ProxyForge | XXXXX |
| Caddy | XXXXX |

> Replace the numbers with your benchmark results.

---

## Project Structure

```
.
├── cmd/
├── proxy/
├── pool/
├── parser/
├── internal/
├── benchmark/
└── README.md
```

---

## Future Improvements

- [ ] HTTP/2 support
- [ ] HTTPS termination
- [ ] Load balancing
- [ ] Health checks
- [ ] Rate limiting
- [ ] Metrics endpoint
- [ ] Compression
- [ ] Caching
- [ ] WebSocket support

---

## Motivation

Building infrastructure software from scratch is one of the best ways to understand how modern backend systems work.

This project explores networking, concurrency, and systems design by implementing a production-inspired reverse proxy from the ground up.

---

## License

MIT
