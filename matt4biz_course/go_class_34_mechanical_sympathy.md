# [Go Class: 34 Mechanical Sympathy](https://www.youtube.com/watch?v=7QLoOd9HinY&t=10s)

- [Go Class: 34 Mechanical Sympathy](#go-class-34-mechanical-sympathy)
  - [The concept](#the-concept)
  - [Memory caching](#memory-caching)
    - [Locality](#locality)
    - [Cache efficiency](#cache-efficiency)
  - [Access patterns](#access-patterns)
  - [Synchronization costs](#synchronization-costs)
  - [Other costs](#other-costs)
  - [Optimization in Go](#optimization-in-go)
  - [Citations](#citations)

## The concept

Make code that works *with* the machine, not against it.

## Memory caching

"Computational" cost is often dominated by **memory access cost**.

Caching takes advantage of access patterns to keep frequently-used code and data "close" to the CPU to reduce access time.

Cahcing imposes some costs of its own:
- Memory access by the **cache line**, typically 64 bytes.
- **Cache coherency** to manage cache line ownership.
 
### Locality

 **Locality in space:** Access to one thing implies access to another nearby thing.
 **Locality in time:** Access implies we are likely to access it again soon.


### Cache efficiency

Cache is effective when we use (and reuse) entire cache lines.

Caching is effective when we access memory in predictable patterns (but only sequential access is predictable).

We get our best performance when we
- **Keep things in contiguous memory**
- **Access them sequentially**


Things that make the cache **less efficient:**
- Synchronization between CPUs
- Copying blocks of data around in memory
- Non sequential access patterns (calling functions, chasing pointers) **A little copying is better than a lot of pointer chasing!**

Things that make the cache **more efficient:**
- Keeping code or data in cache longer
- Keeping data together (so all of a cache line is used)
- Processing memory in sequential order (code or data)


## Access patterns

A slice of objects beats a list with pointers.
A struct with contiguous fields beats a class  with pointers.
Calling lots short methods via dynamic dispatch is very expensive.

## Synchronization costs

Synchroniaation has two costs
- The actual cost of synchronization (lock and unlock)
- The impact of contention if we create a "hot spot" that many threads are trying to access.

In the worst case, synchronization can make a program sequential.

Amdahl's Law: Total speedup is limited by the fraction of the program that runs sequentially.



## Other costs
- False sharing: Cores fight over cache lines for different variables.
- Disk access
- Garbage collection, solve it by:
  - Reduce the unnecessary allocations
  - Reduce embedded pointers in objects
  - Paradoxically, you may want a larger heap to reduce garbage collection frequency.

- Virtual memory & its cache
- Context switching between processes


## Optimization in Go

Go encourages good desing, you can choose:

- To allocate contiguo  
- To copy or not to copy
- To allocate on the stack or heap (sometimes)
- To be synchronous or asynchronous
- To avoid unnecessary abstraction layers
- To avoid short/forwarding methods

Go does not get between you and the machine.

Good code in Go does not hide the costs involved.


## Citations
**Don Knuth**: "Premature optimization is the root of all evil."
**Michael Fromberger**: "There are only three kinds of optimizations: do less, do it less often, and do it faster. The largest gains comes from the first, but we spend most of our time on the third."