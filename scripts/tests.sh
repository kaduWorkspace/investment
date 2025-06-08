#!/bin/bash
set -e  # Exit immediately if any command fails
ssh deployer@172.17.0.2 <<EOF
    cd investment
    /usr/local/go/bin/go test -v -short ./...
EOF
