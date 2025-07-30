#!/bin/bash

# This script generates golang interfaces from the gbfs schemas (npx quicktype)
# It takes the gbfs version number as an argument
# Using this script can be done as follows: ./scripts/generate_go_models.sh 3.0
# Just make sure to have the gbfs schemas in the correct folder ./vX.X

# The script relative path. This make sure the script can be executed from any folder within the repository.
SCRIPT_PATH="$(dirname -- "${BASH_SOURCE[0]}")"

gbfs_version="v$1" #$1 is the first argument passed to the script (the version number)
gbfs_version_no_decimal=$(echo "$1" | tr -d '.' | tr -d -)
folder_path="$SCRIPT_PATH/../$gbfs_version/"
go_folder="$SCRIPT_PATH/../models/golang/"
output_path="$go_folder/$gbfs_version/"
copyright_file="$SCRIPT_PATH/copyright.txt"

# Iterate over all the files in the folder of the gbfs version
for file in "$folder_path"/*
do
    # Extract the file name from the path and add it to the array
    file_name=$(basename "$file")
    file_name_no_extension="${file_name%.*}"
    echo $file_name_no_extension

    output_path_directory="$output_path/$file_name_no_extension/"
    output_file="$output_path_directory/$file_name_no_extension.go"

    # # generates model files from schema files
    # npx quicktype -s schema "$folder_path$file_name" -o "$output_file.go" --prefer-unions --just-types
    mkdir -p $output_path_directory
    npx quicktype -s schema "$folder_path/$file_name" -o "$output_file" --package $file_name_no_extension

    # JSON cannot represent time.Time objects, so manually change time.Time to string and removes the import "time"
    # perl over sed because of compatibility issues with MacOS
    perl -pi -e "s/time\.Time/string/g" "$output_file"  
    perl -pi -e "s/import \"time\"//g" "$output_file"

    # Inject copyright text at the top of the file
    temp_file=$(mktemp)
    cat "$copyright_file" "$output_file" > "$temp_file"
    mv "$temp_file" "$output_file"
done
