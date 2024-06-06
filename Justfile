
default: help
set shell := ["zsh", "-cu"]
# Print help and exit
help:
    @just -l
# Run program
run:
    @go run .
# Build program
build:
    @go build .
    
# Compose
compose:
    sudo docker compose up -d

uncompose:
    source ./.env
    sudo docker compose down
    
test: