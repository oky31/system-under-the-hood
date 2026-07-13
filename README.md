# Systems Under The Hood

> A collection of practical experiments, benchmarks, and deep dives into core computer science, database internals, networking, and memory management.

---

## Why This Repository Exists

In an era saturated with framework hype, AI wrappers, and rapid application development, writing high-level code has become a commodity. However, building **resilient, high-throughput, and cost-efficient distributed systems** requires understanding what happens below the abstraction layer.

This repository serves as my personal lab and documentation for exploring **how systems actually work under the hood**—from OS syscalls and memory allocation to database indexing engines and network protocols.

No bloated business logic. No full-stack CRUD apps. Just pure engineering fundamentals, micro-benchmarks, and low-level implementations.

---

## Focus Areas & Core Concepts

This repository is structured into focused, reproducible experiments:

- **Memory & Concurrency:** Garbage collection profiling, allocation optimizations (`sync.Pool`), and lock-free concurrency patterns in Go.
- **Database Internals:** B-Trees vs. LSM-Trees, Write-Ahead Logging (WAL), indexing efficiency, and execution plan profiling.
- **Networking & Web Architecture:** TCP/UDP sockets, custom HTTP/1.1 parsers, connection pooling, and latency optimization.
- **Systems & OS:** Linux syscalls, memory mapping (`mmap`), CPU cache line behavior, and low-level assembly basics.

---


