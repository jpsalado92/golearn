Create a bash script that:

1. Has an entry point `createDirectories`
- Optional parameters `dirDepth` (default to 3),`numDirs` (default to 10), `targetDir` (defaults to some random name containing the string `test_data` in the same location the executable is located)
- Validations: If targetDir is provided, and it exists, error saying that the target dir should not exist for safety reasons. Also, when creating dirs, make sure a dir with the same name does not exist in the same location.
- Calls a function `createNestedDirs` which populates a `targetDir` with a random number of directories (max `numDir`) with random names. These directories should also carry more directories the same way, up until a depth of 3 from the root folder. Be ware that this operation should not lead to an infinite loop.
- Logging: Once it is finished, it prints the amount of directories created under the targetDir location.


2. Has an entry point called `createFiles`
- Required parameter `targetDir` (path in which to create files); `fileSizes` a comma separated string of values indicating the different possible lengths for each file. As an example, it could be 3B,1KB,50MB,1GB
- Validations: `targetDir` should exist. Also, when creating files, make sure a file with the same name does not exist in the same location.
- Optional parameter `numFiles`, total amount of files to be created.
- It should cache the different possible directories within `targetDir` and create files with random names and random size selected from `fileSizes`, in a random directory selected from the cached resource.
- Logging: Once it is finished, it outputs the amount of different size files created under the given targetDir. It also prints the total size of everything and the total amount of files.

3. Has an entry point called `duplicateFiles`
- Required parameter `targetDir` (path in which to create files); `numTargetFiles` (the amount of times a target file should be picked to be copied) `copiesPerFile` (the amount of times a given target file should be copied in different locations)
- Validate for name colision.
- Caches every dir location and file location in `targetDir`. Then it picks a random file from the cached list and copies it with another random name to another location in `targetDir`. This operation is done an amount of times specified by `copiesPerfile` for the same file, so different locations have to be picked for each file. There should be a safeguard to avoid colision of names, even if the names are random.
- Logging. Outputs the total sum of duplicates created.

4. Integrates concise and well explained --help entries for each of these features.

5. Use appropiate coloring for the logging messages, making the script look good.

6. Uses caching techniques to ensure everything runs as fast as possible.
