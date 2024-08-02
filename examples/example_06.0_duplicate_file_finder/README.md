# Duplicate file finder

This project consist of a duplicate file finder written in Go.
The tool compares the md5 hash of the files in a given locaitonto find duplicates.

## Versions

As a learning project, it also introduces different approaches. Listed from the most basic to the most advanced:

### 1. Sequential processing

In this approach the files are processed sequentially, no concurrency is used.

### 2. Concurrent approach

Channels are used to pass data across goroutines and the main-thread.

- `filePathMetaDataChan`
- `pairChan`
- `doneChan`
- `resultChan`

The goroutines are:

#### `getPairs`

- It gets started as many times as the number of workers.
- Gets a file location from `filePathMetaDataChan`, calculates the hash and sends it to `pairChan` along with the `filePathMetadata`.
- When all files are processed, each goroutine sends a done signal to `doneChan`.

#### `collectHashMap`

- Collects the pairs from `pairChan` and puts them in a map.
- When the `pairChan` is closed, results are sent back to `resultChan`.

Explanation:

1. A number of workers are created based on the number of CPUs.
1. The main-thread starts the goroutines (getPairs started as many times as workers are available), which are on "wait" until the `filePathMetaDataChan` gets some input.
1. Then it starts the `searchTree` function to send the file locations to `filePathMetaDataChan`.
1. Once all the files have been sent, the `filePathMetaDataChan` channel is closed.
1. This causes the `getPairs` goroutines to finish their work and send the done signal to `doneChan`.
1. Once a number of done signals equal to the number of workers are received, the `pairChan` is closed.
1. This causes the `collectHashMap` goroutine to finish its work and send the results to `resultChan`.

- Concurrent file discovery finder with wait groups
- Channels to pass down files to hash calculation
- Channel for reciving file calculate hash and pass hash
- Channel for putting hash in a map, wait for every hash to be calculated, and send done.

### 3. Improved concurrent approach

In this approach some improvements are made to the previous concurrent approach:

- The `filePathMetaDataChan` becomes a buffered channel, so that the main-thread is not blocked until the next hash is calculated.
- `searchTree` starts goroutines per subdirectory. As the amount of directories is unknown,
we close the channel after all those routines have been finished, by using a `sync.WaitGroup`.

### 4. Improved concurrent approach with many more workers
This approach is similar to the previous one, but with a higher number of workers.
- A channel `limitChan` is implemented so that the number of workers is limited.

## Benchmark

In the end a benchmark is performed to compare the different approaches.
It can be easily run by executing the `run.sh` script.

## Additional resources

The project comes with a bash script `create_random_files.sh` to create random files for testing purposes.
