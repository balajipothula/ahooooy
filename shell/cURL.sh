#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Wednesday, 23 July 2025,
# Description : cURL commands.

curl --location 'https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/'

curl --location --request PUT 'https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/insert/song' \
--header 'Content-Type: application/json' \
--data '{
  "artist": "Linkin Park",
  "title": "In the End",
  "difficulty": 3.7,
  "level": 5,
  "released": "2000-10-24"
}'

curl --location 'https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/select/songs'

# Insert song record into webapp_db.public."Song" table.
# Compatible with Power Shell and Unix shell.
curl -X PUT https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/insert/song -H "Content-Type: application/json" -d '{ "artist": "Linkin Park", "title": "In the End", "difficulty": 3.7, "level": 5, "released": "2000-10-24" }'

# 
curl -X GET https://vvjxyyap33.execute-api.eu-central-1.amazonaws.com/select/songs -H "Accept: application/json"