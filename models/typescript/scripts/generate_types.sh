#!/bin/bash

# This script generates typescript interfaces from the gbfs schemas (npx quicktype)
# It also generates files that help with testing the generated interfaces (npx ts-interface-builder)
# It takes the gbfs version number as an argument
# Using this script can be done as follows: ./scripts/generate_types.sh 3.0
# Just make sure to have the gbfs schemas in the correct folder ./vX.X

gbfs_version="v$1" #$1 is the first argument passed to the script (the version number)
parent_dir="$(dirname "$(dirname "$0")")"
folder_path="../../$parent_dir/$gbfs_version/"
output_path="$parent_dir/$gbfs_version/"
test_path="$parent_dir/$gbfs_version/tests/ti/"
copyright_file="$parent_dir/scripts/copyright.txt"

# Iterate over all the files in the folder of the gbfs version
for file in "$folder_path"/*
do
    # Extract the file name from the path and add it to the array
    file_name=$(basename "$file")
    file_name_no_extension="${file_name%.*}"
    echo $file_name_no_extension

    output_file="$output_path$file_name_no_extension"

    # generates model files from schema files
    npx quicktype -s schema "$folder_path$file_name" -o "$output_file.ts" --prefer-unions --just-types

    # JSON cannot represent Date objects, so manually change Date to string
    sed -i "" "s/Date;/string;/g" "$output_file.ts"

    # generates files for typescript testing
    npx ts-interface-builder "$output_file.ts" -o $test_path

    # renames the generated file to .d.ts as it's a declaration file
    # cannot be done right away to .d.ts as it messes up the name of the interfaces
    mv "$output_file.ts" "$output_file.d.ts"

    # Inject copyright text at the top of the file
    temp_file=$(mktemp)
    cat "$copyright_file" "$output_file.d.ts" > "$temp_file"
    mv "$temp_file" "$output_file.d.ts"
done