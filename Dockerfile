# Dockerfile
FROM golang:1.19-alpine

WORKDIR /app
EXPOSE 3000
ARG env
ENV MODE $env

COPY . ./

RUN go mod download

RUN go build -o /wishlist

CMD [ "/wishlist" ]
