FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY models ./models
COPY repository ./repository

RUN go build -o ./app .

EXPOSE 8080

ENTRYPOINT ["./app"]