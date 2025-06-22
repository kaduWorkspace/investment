#!/bin/bash
docker run -d --name fin-dev \
  -e POSTGRES_PASSWORD=postgres \
  --hostname fin-dev \
  --network fin \
  -p 5434:5432 \
  fin-dev

