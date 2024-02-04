#!/bin/bash

# Loop through each .sol file in the src directory
find src -name "*.sol" | while read -r sol_file; do
    # Compute directory paths and filenames
    dir_path=$(dirname "$sol_file" | sed "s/src\///")
    base_name=$(basename "$sol_file" .sol)

    # Create corresponding directory in artifacts
    mkdir -p "artifacts/$dir_path"

    # Compile the Solidity file
    solc --pretty-json --optimize --overwrite --base-path ./src --include-path ./lib \
         --bin "$sol_file" --abi "$sol_file" \
         -o "artifacts/$dir_path"

    # Create corresponding directory in contracts
    mkdir -p "contracts/$dir_path/$base_name"

    # Generate Go bindings
    abigen --bin "artifacts/$dir_path/$base_name.bin" \
           --abi "artifacts/$dir_path/$base_name.abi" \
           --pkg contracts \
           --out "contracts/$dir_path/$base_name/$base_name.go"
done
