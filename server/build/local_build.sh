#! /bin/bash

# Setup the db First 
docker compose up -d jokeDb
cd ./db/cmd
go run main.go


# Compose Local Dev ENV
docker compose up -d server


