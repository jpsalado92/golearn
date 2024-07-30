./create_random_files.sh createDirectories --dirDepth 3 --numDirs 10 --targetDir test_data
./create_random_files.sh createFiles --targetDir test_data/ --fileSizes 10M --numFiles 1000
./create_random_files.sh duplicateFiles --targetDir test_data/ --numTargetFiles 4 --copiesPerFile 1