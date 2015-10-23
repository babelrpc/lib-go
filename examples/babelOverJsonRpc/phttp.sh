#!/bin/bash
mkdir proxy
mkfifo proxy/hproxypipe
cat proxy/hproxypipe | nc -l 8555 | tee -a proxy/hinflow | nc localhost 8333 | tee -a proxy/houtflow 1>proxy/hproxypipe
