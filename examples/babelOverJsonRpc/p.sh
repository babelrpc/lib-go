#!/bin/bash
mkdir proxy
mkfifo proxy/proxypipe
cat proxy/proxypipe | nc -l 8444 | tee -a proxy/inflow | nc localhost 8222 | tee -a proxy/outflow 1>proxy/proxypipe
