FROM golang:1.23
LABEL authors="katy"

ENTRYPOINT ["top", "-b"]