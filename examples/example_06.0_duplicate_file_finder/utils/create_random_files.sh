#!/bin/bash

# Color definitions for logging
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BOLD_WHITE='\033[1;37m'
NC='\033[0m' # No Color

# Function to display general help
show_help() {
  echo
  echo -e "${GREEN}Usage:${NC}"
  echo -e "  ${YELLOW} createDirectories${NC}  Create nested directories."
  echo -e "  ${YELLOW} createFiles${NC}        Create files of specified sizes in the target directory."
  echo -e "  ${YELLOW} duplicateFiles${NC}     Duplicate existing files in the target directory."
  echo -e "${GREEN}Options:${NC}"
  echo -e "  ${YELLOW}--help${NC}  Show this help message and exit"
  echo
}

# Function to display help for createDirectories
show_help_createDirectories() {
  echo
  echo -e "${GREEN}createDirectories:${NC} Create nested directories."
  echo -e "${GREEN}Usage:${NC}"
  echo -e "  ${YELLOW} createDirectories [--dirDepth <depth>] [--numDirs <number>] [--targetDir <directory>]${NC}"
  echo -e "${GREEN}Parameters:${NC}"
  echo -e "  ${BOLD_WHITE}--dirDepth${NC}    (Optional) Depth of nested directories. Default is 3."
  echo -e "  ${BOLD_WHITE}--numDirs${NC}     (Optional) Maximum number of directories at each level. Default is 10."
  echo -e "  ${BOLD_WHITE}--targetDir${NC}   (Optional) Target directory. Default is 'test_data_<random>'."
  echo -e "${GREEN}Examples:${NC}"
  echo -e "  'createDirectories --dirDepth 3 --numDirs 10 --targetDir my_test_data$'"
  echo
}

# Function to create nested directories
create_nested_dirs() {
  local current_dir=$1
  local depth=$2
  local max_dirs=$3

  if [ $depth -le 0 ]; then
    return
  fi
  if [ $max_dirs -le 0 ]; then
    return
  fi

  for i in $(seq 1 $((RANDOM % max_dirs + 1))); do
    local new_dir="$current_dir/dir_$RANDOM"
    if [ ! -d "$new_dir" ]; then
      mkdir -p "$new_dir"
      create_nested_dirs "$new_dir" $((depth - 1)) $max_dirs
    fi
  done
}

# Entry point for creating directories
createDirectories() {
  local dirDepth=3
  local numDirs=10
  local targetDir="test_data_$RANDOM"

  while [[ "$#" -gt 0 ]]; do
    case $1 in
    --dirDepth)
      dirDepth="$2"
      shift
      ;;
    --numDirs)
      numDirs="$2"
      shift
      ;;
    --targetDir)
      targetDir="$2"
      shift
      ;;
    --help)
      show_help_createDirectories
      exit 0
      ;;
    *)
      echo -e "${RED}Unknown parameter passed: $1${NC}"
      show_help_createDirectories
      exit 1
      ;;
    esac
    shift
  done

  if [ -d "$targetDir" ]; then
    echo -e "${RED}Error: Target directory '$targetDir' already exists.${NC}"
    exit 1
  fi

  mkdir -p "$targetDir"
  create_nested_dirs "$targetDir" "$dirDepth" "$numDirs"

  local total_dirs=$(find "$targetDir" -type d | wc -l)
  echo -e "${GREEN}Created $total_dirs directories under '$targetDir'.${NC}"
}

# Function to display help for createFiles
show_help_createFiles() {
  echo
  echo -e "${GREEN}Usage:${NC}"
  echo -e "  ${YELLOW} createFiles --targetDir <directory> --fileSizes <sizes> [--numFiles <number>]${NC}"
  echo -e "${GREEN}Parameters:${NC}"
  echo -e "  ${BOLD_WHITE}--targetDir${NC}   (Required) Target directory."
  echo -e "  ${BOLD_WHITE}--fileSizes${NC}   (Required) Comma-separated list of file sizes (e.g., 1K,50M,1G)."
  echo -e "  ${BOLD_WHITE}--numFiles${NC}    (Optional) Number of files to create. Default is 10."
  echo -e "${GREEN}Examples:${NC}"
  echo -e "  ${YELLOW} createFiles --targetDir my_test_data --fileSizes 3B,1KB,50MB,1GB --numFiles 100${NC}"
  echo
}
# Entry point for creating files
createFiles() {
  local targetDir=""
  local fileSizes=""
  local numFiles=10

  while [[ "$#" -gt 0 ]]; do
    case $1 in
    --targetDir)
      targetDir="$2"
      shift
      ;;
    --fileSizes)
      fileSizes="$2"
      shift
      ;;
    --numFiles)
      numFiles="$2"
      shift
      ;;
    --help)
      show_help_createFiles
      exit 0
      ;;
    *)
      echo -e "${RED}Unknown parameter passed: $1${NC}"
      show_help_createFiles
      exit 1
      ;;
    esac
    shift
  done

  if [ ! -d "$targetDir" ]; then
    echo -e "${RED}Error: Target directory '$targetDir' does not exist.${NC}"
    exit 1
  fi

  IFS=',' read -r -a sizes <<<"$fileSizes"
  dirs=($(find "$targetDir" -type d))

  for i in $(seq 1 "$numFiles"); do
    local size=${sizes[$RANDOM % ${#sizes[@]}]}
    local dir=${dirs[$RANDOM % ${#dirs[@]}]}
    local file="$dir/file_$RANDOM"

    if [ ! -f "$file" ]; then
      head -c "$size" /dev/urandom >"$file"
    fi
  done

  local total_files=$(find "$targetDir" -type f | wc -l)
  local total_size=$(du -sh "$targetDir" | cut -f1)
  echo -e "${GREEN}Created $numFiles files under '$targetDir'.${NC}"
  echo -e "${GREEN}Total size: $total_size, Total files: $total_files.${NC}"
}

# Function to display help for duplicateFiles
show_help_duplicateFiles() {
  echo
  echo -e "${GREEN}Usage:${NC}"
  echo -e "  ${YELLOW} duplicateFiles --targetDir <directory> --numTargetFiles <number> --copiesPerFile <number>${NC}"
  echo -e "${GREEN}Parameters:${NC}"
  echo -e "  ${BOLD_WHITE}--targetDir${NC}       (Required) Target directory."
  echo -e "  ${BOLD_WHITE}--numTargetFiles${NC}  (Optional) Number of files to duplicate. Default is 10."
  echo -e "  ${BOLD_WHITE}--copiesPerFile${NC}   (Optional) Number of copies per file. Default is 5."
  echo -e "${GREEN}Examples:${NC}"
  echo -e "  'duplicateFiles --targetDir my_test_data --numTargetFiles 10 --copiesPerFile 5'"
  echo
}

# Entry point for duplicating files
duplicateFiles() {
  local targetDir=""
  local numTargetFiles=10
  local copiesPerFile=5

  while [[ "$#" -gt 0 ]]; do
    case $1 in
    --targetDir)
      targetDir="$2"
      shift
      ;;
    --numTargetFiles)
      numTargetFiles="$2"
      shift
      ;;
    --copiesPerFile)
      copiesPerFile="$2"
      shift
      ;;
    --help)
      show_help_duplicateFiles
      exit 0
      ;;
    *)
      echo -e "${RED}Unknown parameter passed: $1${NC}"
      show_help_duplicateFiles
      exit 1
      ;;
    esac
    shift
  done

  if [ ! -d "$targetDir" ]; then
    echo -e "${RED}Error: Target directory '$targetDir' does not exist.${NC}"
    exit 1
  fi

  files=($(find "$targetDir" -type f))
  dirs=($(find "$targetDir" -type d))

  for i in $(seq 1 "$numTargetFiles"); do
    local src_file=${files[$RANDOM % ${#files[@]}]}
    for j in $(seq 1 "$copiesPerFile"); do
      local dest_dir=${dirs[$RANDOM % ${#dirs[@]}]}
      local dest_file="$dest_dir/duplicate_$RANDOM"

      if [ ! -f "$dest_file" ]; then
        cp "$src_file" "$dest_file"
      fi
    done
  done

  local total_duplicates=$(find "$targetDir" -name 'duplicate_*' | wc -l)
  echo -e "${GREEN}Created $total_duplicates duplicate files under '$targetDir'.${NC}"
}

# Main script logic to handle entry points
if [[ "$#" -lt 1 ]]; then
  show_help
  exit 1
fi

case $1 in
createDirectories)
  shift
  createDirectories "$@"
  ;;
createFiles)
  shift
  createFiles "$@"
  ;;
duplicateFiles)
  shift
  duplicateFiles "$@"
  ;;
--help) show_help ;;
*)
  echo -e "${RED}Unknown command: $1${NC}"
  show_help
  exit 1
  ;;
esac
