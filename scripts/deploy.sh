#!/bin/bash
ssh -t deployer <<EOF
    newgrp docker
    docker pull kaduhod/fin
    docker stop cdb || true
    docker rm cdb || true
    docker run --name cdb -d -p 3000:8989 kaduhod/fin
EOF
