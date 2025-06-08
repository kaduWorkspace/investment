#!/bin/bash
set -e  # Exit immediately if any command fails
ssh -t deployer@172.17.0.2 <<EOF
    docker image prune -a
EOF
