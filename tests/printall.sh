#!/bin/bash

# Loop through all regular files in the current directory
for file in *; do
    if [ -f "$file" ]; then
        echo "----- Contents of: $file -----"
        cat "$file"
        echo "------------------------------"
        echo
    fi
done

