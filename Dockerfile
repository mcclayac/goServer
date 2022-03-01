# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

# copy dependencies references
COPY go.mod ./
# COPY go.sum ./

# Downloading the dependencies
RUN go mod download

# copy the application code
COPY *.go ./

# build the application
RUN go build -o /docker-hostname

expose 9090

CMD [ "/docker-hostname" ]





