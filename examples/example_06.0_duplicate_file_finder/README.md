# Duplicate file finder

This project consist of a duplicate file finder written in Go.
The tool compares the md5 hash of the files in a given locaitonto find duplicates.

## Versions

As a learning project, it also introduces different approaches. Listed from the most basic to the most advanced:

1. Sequential processing DONE

2. Concurrent approach

- Concurrent file discovery finder with wait groups
- Channels to pass down files to hash calculation
- Channel for reciving file calculate hash and pass hash
- Channel for putting hash in a map, wait for every hash to be calculated, and send done.

## Benchmark

In the end a benchmark is performed to compare the different approaches.
It can be easily run by executing the `run.sh` script.

## Additional resources

The project comes with a bash script `create_random_files.sh` to create random files for testing purposes.
