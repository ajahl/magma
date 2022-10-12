#!/bin/bash

docker-compose events | grep --line-buffered 'container restart.*name=sctpd' | while read REPLY;
do 
    docker-compose stop sctpd mobilityd pipelined sessiond oai_mme ;
    sudo su -c '/usr/bin/env python3 /usr/local/bin/config_stateless_agw.py sctpd_pre' ;
    docker-compose start sctpd mobilityd pipelined sessiond oai_mme ;
done