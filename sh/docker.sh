#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Tuesday, 26 August 2025,
# Description : docker commands.

# Log in to a registry
#
# --username : Username
docker login --username balajipothula

# Start a build
#
# --file  : Name of the Dockerfile
# --quiet : Suppress the build output and print image ID on success
# --tag   : Name and optionally a tag
docker buildx build \
  --quiet \
  --tag=balajipothula/mini_httpd:1.30-r5 \
  --file=./docker/Dockerfile.mini_httpd1.30-r5 .

# Upload an image to a registry
docker image push balajipothula/mini_httpd:1.30-r5

# Create and run a new container from an image
#
# --name    : Assign a name to the container
# --detach  : Run container in background and print container ID
# --restart : Restart policy to apply when a container exits
# --publish : Publish a container's port(s) to the host
docker container run \
  --name=go_fiber_app \
  --detach=true \
  --restart=unless-stopped \
  --publish=127.0.0.1:3000:3000/tcp \ 
  balajipothula/mini_httpd:1.30-r5
