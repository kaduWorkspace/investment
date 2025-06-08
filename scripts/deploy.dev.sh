#!/bin/bash
ssh -t deployer@172.17.0.2 <<EOF
    newgrp docker
    docker stop cdb || true
    docker rm cdb || true
    docker run --name cdb -d -p 3000:3000 kaduhod/fin
EOF
