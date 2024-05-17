#!/bin/bash

# This script copies the gbfs schemas to subfolders in the resources folder of the gbfs-java-model
# Reason is to create separate packages for each gbfs file
# Runs from models/java folder using 'mvn install'

# brings the script to the root directory
cd ../../..

# clears old resources
rm -rf ./models/java/gbfs-java-model/resources

output_path="./models/java/gbfs-java-model/resources/"
path_to_schemas="."

versions=("v3.0" "v2.3" "v2.2" "v2.1" "v2.0" "v1.1" "v1.0")

echo "Start Schema to Subfolders Script"

for version in "${versions[@]}"; do
    folder_path="$path_to_schemas/$version/"

    # Iterate over all the files in the folder of the gbfs version
    for file in "$folder_path"/*
    do
        # Extract the file name from the path and add it to the array
        file_name=$(basename "$file")
        file_name_no_extension="${file_name%.*}"

        output_file="$output_path$file_name_no_extension"

        mkdir -p "${output_path}/${version}/${file_name_no_extension}/"

        # Copy the file to the destination directory
        cp "$folder_path/$file_name" "${output_path}/${version}/${file_name_no_extension}/${file_name}"
    done
done

echo "End Schema to Subfolders Script"

