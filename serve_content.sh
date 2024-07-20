#!/bin/bash

# Function to list directories
list_directories() {
    echo "Which directory do you want to host?:"
    select dir in */; do
        if [ -n "$dir" ]; then
            echo "You selected '$dir'"
            break
        else
            echo "Invalid selection. Please choose a valid directory."
        fi
    done
}

# List directories and prompt user for selection
list_directories

# Remove trailing slash from the directory name
TARGET_DIR=${dir%/}

# Define the port for the local server
PORT=8080

# Start a simple HTTP server to serve the files
echo "Serving files from '$TARGET_DIR' on http://localhost:$PORT"
python3 -m http.server $PORT --directory "$TARGET_DIR"
    