#!/bin/bash

# Download and extract GBFS json schema
validate () {
    NAME=$1
    VALUE=$2

    if [ -z ${VALUE} ];
    then
        echo "$NAME not set"
        exit 1
    fi
}

validate "GITHUB_URL" ${GITHUB_URL}
validate "SCHEMA_VERSION" ${SCHEMA_VERSION}
validate "DESTINATION_PATH" ${DESTINATION_PATH}

ZIP_FILE=downloaded.zip

echo "GBFS JSON schema repo github URL: $GITHUB_URL"

echo "Removing any existing contents in $DESTINATION_PATH"
rm -rf ${DESTINATION_PATH}/*
mkdir -p ${DESTINATION_PATH}

if [ -f ${ZIP_FILE} ]; then
  echo "Removing existing file $ZIP_FILE"
  rm ${ZIP_FILE}
fi

WGET_URL="${GITHUB_URL}"
echo "About to download from $WGET_URL"
wget -q ${WGET_URL} -O ${ZIP_FILE}

if [ -f ${ZIP_FILE} ]; then
    echo "Done"
    {
    echo "Create ${DESTINATION_PATH}" &&
    mkdir -p ${DESTINATION_PATH} &&

    echo "Unzip files from zip file ${ZIP_FILE} to ${DESTINATION_PATH}" &&
    unzip -q ${ZIP_FILE} -d ${DESTINATION_PATH} &&

    echo "Remove zipfile ${ZIP_FILE}" &&
    rm ${ZIP_FILE} &&

    echo "Remove intermediate folder" &&
    mv ${DESTINATION_PATH}/gbfs-json-schema-${SCHEMA_VERSION}/* ${DESTINATION_PATH} &&
    rm -rf ${DESTINATION_PATH}/gbfs-json-schema-${SCHEMA_VERSION}

    echo "Move each schema file into its own subfolder"
    find ${DESTINATION_PATH}/. -type f -exec sh -c 'mkdir "${0%.*}" && mv "$0" "${0%.*}/"' {} \;

    echo "JSON schema extracted to $DESTINATION_PATH"
    } ||
    {
        (>&2 echo "Error extracting zip file $ZIP_FILE from $WGET_URL. See my previous output for details")
        exit 1
    }
else
    (>&2 echo "Error downloading zip from $WGET_URL")
    exit 1
fi