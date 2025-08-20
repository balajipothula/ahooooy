#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Wednesday, 23 July 2025,
# Description : cURL commands.

curl --location 'https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/'


curl -X POST http://localhost:3000/songs \
  -H "Content-Type: application/json" \
  -d '{
    "artist": "Linkin Park",
    "title": "In the End",
    "difficulty": 3.7,
    "level": 5,
    "released": "2000-10-24T00:00:00Z"
  }'

curl http://localhost:3000/songs

curl http://localhost:3000/songs/1

curl -X PUT http://localhost:3000/songs/1 \
  -H "Content-Type: application/json" \
  -d '{
    "artist": "Linkin Park",
    "title": "Numb",
    "difficulty": 4.0,
    "level": 6,
    "released": "2003-03-25T00:00:00Z"
  }'

curl -X DELETE http://localhost:3000/songs/1



curl \
  --location \
  --request POST 'https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/insert/song' \
  --header 'Content-Type: application/json' \
  --data '{
    "artist": "Linkin Park",
    "title": "In the End",
    "difficulty": 3.7,
    "level": 5,
    "released": "2000-10-24T00:00:00Z"
  }'
