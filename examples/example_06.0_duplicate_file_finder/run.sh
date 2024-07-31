# ./utils/create_random_files.sh createDirectories --dirDepth 3 --numDirs 10 --targetDir test_data
# ./utils/create_random_files.sh createFiles --targetDir test_data/ --fileSizes 10M --numFiles 1000
# ./utils/create_random_files.sh duplicateFiles --targetDir test_data/ --numTargetFiles 4 --copiesPerFile 1

echo "Benchmarking sequential approach..."
time go run 1.0_sequential_approach/main.go test_data/

echo "Benchmarking concurrent approach..."
time go run 2.0_concurrent_approach/main.go test_data/
