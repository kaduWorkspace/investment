#!/bin/bash
ssh -t deployer@172.18.0.2 <<EOF
    newgrp docker
    docker stop cdb || true
    docker rm cdb || true
    docker run --name cdb -d -p 3000:8989 kaduhod/fin
EOF
