#!/bin/bash

process_value() {
    local value=$1
    if [[ $value =~ IP$ ]]; then
        processed_value=$(echo "$value" | sed -E 's/.*([A-Z][a-z]+[A-Z][a-z]*IP)$/\1/')
    else
        processed_value=$(echo "$value" | sed -E 's/.*([A-Z][a-z]*)$/\1/')
    fi
    echo "$processed_value" | tr '[:upper:]' '[:lower:]'
}

find . -name "*.go" | while read -r file; do
    grep 'Use: *".*",' "$file" | awk -F'"' '{print $2}' | while read -r value; do
        processed_value=$(process_value "$value")

        quoted_value="\"$value\""
        quoted_processed_value="\"$processed_value\""
        find . -name "*.go" -exec sed -i '' "s/$quoted_value/$quoted_processed_value/g" {} +
    done
done
