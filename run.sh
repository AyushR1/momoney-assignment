#!/bin/bash

echo "Building server"
sudo docker build -t server .

echo "Running server"
sudo docker run -p 8080:8080 -t server
