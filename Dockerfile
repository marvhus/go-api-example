FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o server ./src

ENV GIN_MODE=release
ENV API_ENV=prod
EXPOSE 443
CMD [ "/app/server" ]
