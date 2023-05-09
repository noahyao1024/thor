#!/bin/bash

PWD=$(pwd)

# Directory to iterate through
dir=$PWD/html

# Check if openssl command exists
if ! command -v openssl &>/dev/null; then
    echo "openssl command not found. Exiting..."
    exit
fi

routes='package server\n
import (
	"encoding/base64"\n
	"golib/fakehtml"\n
	"net/http"\n
	"github.com/gin-gonic/gin"\n
)\n

func stage2(e *gin.Engine) *gin.Engine {
'

# Loop through each file in the directory
for file in $dir/*; do
    # Check if the file is a regular file (i.e. not a directory, symbolic link, etc.)
    if [ -f "$file" ]; then
        # The base64-encoded contents of the file
        base64_content=$(openssl base64 -in $file -A)

        # The name of the new file
        new_file="${file%.*}.go"

        # Split file by /, and save the last one
        relative_file=$(echo $file | awk -F/ '{print $NF}')

        func_name="$(tr '[:lower:]' '[:upper:]' <<<${relative_file:0:1})${relative_file:1}"

        func_name=${func_name//./}

        # Format the contents of the new file using the template

        # Format the contents of the new file using the template
        new_file_content="package fakehtml\n\nfunc $func_name() string {\n\treturn \"$base64_content\"\n}"

        # Check the suffix of the file
        suffix="${relative_file##*.}"
        content_type="text/html; charset=utf-8"
        if [ "$suffix" == "js" ] || [ "$suffix" == "css" ]; then
            content_type="application/javascript"
        fi

        routes+="\n\te.GET(\"/fakehtml/${relative_file}\", func(c *gin.Context) {\n\t\tdecodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.${func_name}())\n\t\tc.Data(http.StatusOK, \"$content_type\", decodedBytes)\n\t})"

        # Write the formatted contents to the new file
        echo -e "$new_file_content" >fakehtml/$func_name.go

        # Print the name of the new file to the console
        echo "Encoded $file to $new_file"
    fi
done

routes+="\n\treturn e\n}"

echo -e $routes >$PWD/server/entry_generated.go

SQLITE_PATH="/Users/morysky/repos/thor/go" go run main.go
