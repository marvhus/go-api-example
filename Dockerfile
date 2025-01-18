# Builds the project so that it can be used within a docker container.

FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o server ./src
