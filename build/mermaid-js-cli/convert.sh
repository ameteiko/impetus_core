#!/bin/sh

# This script receives an optional path variable, then searches
#
# Usage:
#    conv.sh                     # Converts all *.mmd files in the current directory and all subdirectories.
#    conv.sh ./path/to/images    # Converts all *.mmd files in a specified directory.

# Conversion modes. It can be either converting a whole directory (dir mode), or a specific file (file mode).
MODE_DIR="dir"
MODE_FILE="file"
RESULT_EXT=".svg"
FILE_EXT=".mmd"

# Default behaviour is to convert current working directory.
MODE=$MODE_DIR
CONVERTDIR=.
CONVERTFILE=

# Process the command line arguments.
if [ $# -eq 1 ]; then
    if [ -d $1 ]; then
        CONVERTDIR="$1"
    elif [ -e $1 ]; then
        CONVERTFILE=$1
        MODE=$MODE_FILE
    else
        echo "Incorrect argument. $1 is neither existing file, not existing directory."
        exit 1
    fi
fi

# Process file conversion.
if [ $MODE = $MODE_FILE ]; then
    file=$(basename -s $FILE_EXT "$CONVERTFILE")
    mmdc -p /puppeteer-config.json -i "$CONVERTFILE" -o "$file$RESULT_EXT"
    exit 0
fi

# Process directory conversion.
for file in $(echo "$CONVERTDIR/**/*$FILE_EXT") $(echo "$CONVERTDIR/*$FILE_EXT"); do
    echo processing "$file"
    name=$(basename -s $FILE_EXT "$file")
    dir=$(dirname "$file")
    mmdc -p /puppeteer-config.json -i "$file" -o "$dir/$name$RESULT_EXT"
done