#!/bin/bash

# Find all files matching the pattern compro_*

find . -type f -name 'compro_*' | while read file; do
  # Extract directory and filename
  dir=$(dirname "$file")
  base=$(basename "$file")
  # Replace compro_ with banner_
  newbase=$(echo "$base" | sed 's/^compro_/banner_/')
  # Rename the file
  mv "$file" "$dir/$newbase"
done
