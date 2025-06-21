#!/bin/bash
set -e  # Exit immediately if any command fails
scp .env.development deployer@172.18.0.2:~/investment
ssh deployer@172.18.0.2 <<EOF
cp /home/deployer/.env /home/deployer/investment/
EOF
