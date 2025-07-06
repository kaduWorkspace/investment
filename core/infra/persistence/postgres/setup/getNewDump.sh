#!/bin/bash
docker exec fin-dev bash -c 'pg_dump -U postgres -Fc dev > /tmp/dev.dump'
docker cp fin-dev:/tmp/dev.dump ./dump.sql
docker exec fin-dev rm /tmp/dev.dump
echo "Dump created at ./dump.sql"
