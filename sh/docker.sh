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
  --tag=balajipothula/go:1.21.0 \
  --file=./docker/Dockerfile.go1.21.0 .

# Upload an image to a registry
docker image push balajipothula/go:1.21.0

# Create and run a new container from an image
#
# --name    : Assign a name to the container
# --detach  : Run container in background and print container ID
# --publish : Publish a container's port(s) to the host
# --restart : Restart policy to apply when a container exits
# --volume  : Bind mount a volume
docker container run \
  --name=go_fiber_app \
  --detach=true \
  --publish=127.0.0.1:3000:3000/tcp \
  --restart=unless-stopped \
  --volume=$HOME/gofiber:/gofiber \
  balajipothula/go:1.21.0

# Execute a command in a running container
#
# --interactive : Keep STDIN open even if not attached
# --tty         : Allocate a pseudo-TTY
docker container exec --interactive --tty go_fiber_app ash
