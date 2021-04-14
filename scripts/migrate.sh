#!/bin/bash

cont=$(docker ps | grep database | awk '{ print $1 }')

docker cp "$(pwd)"/migrations/0-CREATE-DATABASE.sql "$cont":/docker-entrypoint-initdb.d/
docker cp "$(pwd)"/migrations/1-CREATE-USER.sql "$cont":/docker-entrypoint-initdb.d/
docker cp "$(pwd)"/migrations/2-CREATE-JOBS.sql "$cont":/docker-entrypoint-initdb.d/

docker exec -u postgres -it "$cont" psql postgres -f docker-entrypoint-initdb.d/0-CREATE-DATABASE.sql
docker exec -u postgres -it "$cont" psql jobs_db postgres -f docker-entrypoint-initdb.d/1-CREATE-USER.sql
docker exec -u postgres -it "$cont" psql jobs_db postgres -f docker-entrypoint-initdb.d/2-CREATE-JOBS.sql

echo 'migrations finished'