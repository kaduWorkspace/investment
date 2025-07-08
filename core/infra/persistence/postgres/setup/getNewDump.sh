#!/bin/bash
docker exec fin-dev bash -c 'pg_dump -U postgres -Fc dev > /tmp/dev.dump'
docker cp fin-dev:/tmp/dev.dump ../migrations/dump.sql
docker exec fin-dev rm /tmp/dev.dump
DATA=$(date +'%d-%b-%Y' | tr '[:upper:]' '[:lower:]')  # exemplo: 07-jul-2025
cp ../migrations/dump.sql ../migrations/dump.$DATA.sql
echo "Dump criado: ../migrations/dump.sql e ../migrations/dump.$DATA.sql"

