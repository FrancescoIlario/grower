#!/bin/bash

function main() {
    docker-compose -f docker-compose.yaml -p ci up -d --build itest mongo # mongo-express
    docker logs -f ci_itest_1
    docker wait ci_itest_1
    res=$?
    docker-compose -f docker-compose.yaml -p ci down

    return $res
}

main
