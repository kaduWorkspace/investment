#!/bin/bash
docker run -d --name fin-dev \
  -e POSTGRES_PASSWORD=postgres \
  -p 5434:5432 \
  fin-dev

