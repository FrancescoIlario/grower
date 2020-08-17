#!/bin/bash

function main() {
    docker-compose -f docker-compose.yaml -p ci up -d --build itest rabbitmq
    docker logs -f ci_itest_1
    docker wait ci_itest_1
    res=$?
    docker-compose -f docker-compose.yaml -p ci down

    return $res
}

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
pushd $SCRIPTPATH > /dev/null
main
popd > /dev/null
