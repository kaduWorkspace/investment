#!/bin/bash
set -e  # Exit immediately if any command fails
BRANCH=${1:-master}
ssh -t deployer@172.18.0.2 <<EOF
    REPO_DIR="investment"

    if [ -d "investment" ]; then
        if [ -d "$REPO_DIR/.git" ]; then
            echo "O repositório '$REPO_DIR' existe e é um repositório Git."
        else
            echo "Repositorio existe mas nao tem a pasta .git"
            rm -rf ./*
            echo "git clone -b $BRANCH git@github.com:kaduWorkspace/investment.git"
            git clone -b $BRANCH git@github.com:kaduWorkspace/investment.git
        fi
    else
        echo "O repositório '$REPO_DIR' nao existe."
        git clone git@github.com:kaduWorkspace/investment.git
    fi
    cd investment
    echo "branch: $BRANCH"
    git checkout $BRANCH
    git pull
EOF
