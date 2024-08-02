./utils/create_random_files.sh createDirectories --dirDepth 4 --numDirs 10 --targetDir test_data
./utils/create_random_files.sh createFiles --targetDir test_data/ --fileSizes 10M --numFiles 1000
./utils/create_random_files.sh duplicateFiles --targetDir test_data/ --numTargetFiles 1 --copiesPerFile 1

echo "Benchmarking VER1..."
time go run 1.0_sequential_approach/main.go test_data/
echo

echo "Benchmarking VER2..."
time go run 2.0_concurrent_approach/main.go test_data/
echo

echo "Benchmarking VER3..."
time go run 3.0_concurrent_approach_improved/main.go test_data/
echo

echo "Benchmarking VER4..."
time go run 4.0_concurrent_approach_semaphore/main.go test_data/
echo