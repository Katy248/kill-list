
default: help

# Print help and exit
help:
    @just -l
# Run program
run:
    @go run .
# Build program
build:
    @go build .
