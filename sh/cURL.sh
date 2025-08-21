#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Wednesday, 23 July 2025,
# Description : cURL commands.

curl --location --request GET --url 'https://go-fiber-app.fly.dev/'

curl \
  --location \
  --request POST \
  --url https://go-fiber-app.fly.dev/insert/song \
  --header "Content-Type: application/json" \
  --data '{
    "artist": "Linkin Park",
    "title": "In the End",
    "difficulty": 3.7,
    "level": 5,
    "released": "2000-10-24T00:00:00Z"
  }'

curl --location --request GET --url https://go-fiber-app.fly.dev/select/songs

curl --location --request GET --url https://go-fiber-app.fly.dev/select/song/1

curl \
  --location \
  --request PUT \
  --url https://go-fiber-app.fly.dev/update/song/1 \
  --header "Content-Type: application/json" \
  --data '{
    "artist": "Linkin Park",
    "title": "Numb",
    "difficulty": 4.0,
    "level": 6,
    "released": "2003-03-25T00:00:00Z"
  }'

curl --location --request DELETE --url https://go-fiber-app.fly.dev/delete/song/1
