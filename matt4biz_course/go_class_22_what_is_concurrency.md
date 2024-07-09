# [Go Class: 22 What is Concurrency?](https://www.youtube.com/watch?v=A3R-4ZYBqvE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=24)

## Meaning of concurrency

- Execution happens in a non-deterministic order.
- Out of order execution.
- Non sequential execution.
- Parts of a program execute out of order or in partial order.

## Concurrency vs Parallelism

- Concurrency: Multiple tasks are making progress simultaneously.
- Parallelism: Multiple tasks are running at the same time.
- Concurrency is about dealing with lots of things at once.
- Parallelism is about doing lots of things at once.
- Concurrency is about structure.
- Parallelism is about execution.

## Race Condition

Possibility that non-deterministic order of execution can cause unexpected results.

Combating race conditions:

- Do not share data between threads.
- Make the shared things read-only.
- Allow only one thread to write at a time.
- Make the read-modify-write operations atomic.
