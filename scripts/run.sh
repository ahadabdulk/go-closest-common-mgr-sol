#!/bin/bash
docker build -t aahad_solution:1.0 ../src

docker run -it --name aahad_unittest aahad_solution:1.0