#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Tuesday, 26 August 2025,
# Description : docker compose commands.

# Create and start containers
#
# --file   : Compose configuration files
# --detach : Detached mode - Run containers in the background
docker compose --file=./docker-compose/docker-compose-go1.21.0.yml up --detach
