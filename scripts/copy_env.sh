#!/bin/bash
set -e  # Exit immediately if any command fails
scp .env.development deployer@172.17.0.2:~/investment
scp .env deployer@172.17.0.2:~/investment
