#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Sunday, 31 August 2025,
# Description : redis-cli commands.

# -h       : Server hostname.
# -p       : Server port.
# --user   : Used to send ACL style.
# --pass   : Password to use when connecting to the server.
# -n       : Database number.
# --no-raw : Force formatted output even when STDOUT is not a tty.
redis-cli \
  -h 'redis-14589.c311.eu-central-1-1.ec2.redns.redis-cloud.com' \
  -p 14589 \
  --user 'default' \
  --pass "$REDIS_PASSWORD" \
  -n 0 \
  --no-raw

#
GET "otp:balan.pothula@gmail.com"

# 
HGET Email:balan.pothula@gmail.com Code

# 
HGET Email:balan.pothula@gmail.com Code ExpiresAt

# 
HGETALL Email:balan.pothula@gmail.com
