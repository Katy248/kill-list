FROM golang:1.22
LABEL authors="katy"

ENTRYPOINT ["top", "-b"]