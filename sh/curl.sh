#!/bin/bash

# Author      : Balaji Pothula <balan.pothula@gmail.com>,
# Date        : Wednesday, 23 July 2025,
# Description : curl commands.

# get index page
curl --location --request GET --url 'https://ahooooy.fly.dev'

# insert new song
curl \
  --location \
  --request POST \
  --url 'https://ahooooy.fly.dev/insert/song' \
  --header 'Content-Type: application/json' \
  --data '{
    "artist": "Linkin Park",
    "title": "In the End",
    "difficulty": 3.7,
    "level": 5,
    "released": "2000-10-24"
  }'

# select all songs
curl --location --request GET --url 'https://ahooooy.fly.dev/select/songs'

# select song by id
curl --location --request GET --url 'https://ahooooy.fly.dev/select/song/1'

# update song by id
curl \
  --location \
  --request PUT \
  --url 'https://ahooooy.fly.dev/update/song/1' \
  --header 'Content-Type: application/json' \
  --data '{
    "artist": "Linkin Park",
    "title": "Numb",
    "difficulty": 4.0,
    "level": 6,
    "released": "2003-03-25"
  }'

# delete song by id
curl --location --request DELETE --url 'https://ahooooy.fly.dev/delete/song/1'
