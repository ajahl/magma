#!/bin/bash

docker-compose events | grep --line-buffered 'container restart.*name=oai_mme' |  while read REPLY;
do
    docker-compose restart --time 1 mobilityd pipelined sessiond;
done