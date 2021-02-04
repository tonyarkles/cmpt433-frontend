#!/bin/bash
set -ex

tag=$1

if [ ${tag} == "" ]; then
    echo "need to specify a tag"
    exit 1
fi

docker build -t tonypersea/cmpt433-frontend:${tag} .
docker push tonypersea/cmpt433-frontend:${tag}
