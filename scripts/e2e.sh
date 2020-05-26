#!/bin/bash

FIDER_CONTAINER="teamdream_e2e"
PG_CONTAINER="teamdream_pge2e"
PORT=3000

start_teamdream () {
  echo "Starting Teamdream (HOST_MODE: $1)..."
  docker rm -f $FIDER_CONTAINER $PG_CONTAINER || true
  docker run -d -e POSTGRES_USER=teamdream_e2e -e POSTGRES_PASSWORD=teamdream_e2e_pw --name $PG_CONTAINER postgres:9.6.8
  docker run --link $PG_CONTAINER:pg waisbrot/wait
  docker run \
    -d \
    -p 3000:3000 \
    -e SSL_CERT=development.crt \
    -e SSL_CERT_KEY=development.key \
    -e HOST_MODE=$1 \
    -e DATABASE_URL=postgres://teamdream_e2e:teamdream_e2e_pw@$PG_CONTAINER:5432/teamdream_e2e?sslmode=disable \
    -v `pwd`/etc:/app/etc \
    --env-file .env \
    --link $PG_CONTAINER \
    --name $FIDER_CONTAINER getteamdream/teamdream:e2e
}

run_e2e () {
  start_teamdream $1
  echo "Running e2e tests ..."
  npx jest ./tests/$1.spec.ts
}

if [[ $1 == 'build' ]] || [ -z $1 ]
then
  mage build:docker
  docker tag getteamdream/teamdream getteamdream/teamdream:e2e
fi

if [[ $1 == 'single' ]] || [ -z $1 ]
then
  run_e2e single
fi

if [[ $1 == 'multi' ]] || [ -z $1 ]
then
  run_e2e multi
fi

echo "Stopping Containers ..."
docker rm -f $FIDER_CONTAINER $PG_CONTAINER || true

echo "Killing Chromium..."
ps -A | grep '[c]hromium' | awk '{print $1}' | xargs kill || true
