#!/bin/bash
set -e  # Exit immediately if any command fails
ssh -t deployer@172.17.0.2 <<EOF
    REPO_DIR="chi-version"

    if [ -d "chi-version" ]; then
        if [ -d "$REPO_DIR/.git" ]; then
            echo "O repositório '$REPO_DIR' existe e é um repositório Git."
        else
            echo "Repositorio existe mas nao tem a pasta .git"
            rm -rf ./*
            git clone git@github.com:kaduWorkspace/investment.git
        fi
    else
        echo "O repositório '$REPO_DIR' nao existe."
        git clone git@github.com:kaduWorkspace/chi-version.git
    fi
    cd chi-version
    git checkout master
    git pull
EOF
