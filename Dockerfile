FROM golang:1.21 AS build-stage

ENV GIN_MODE=release

# Install dependencies
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./
COPY models ./models
COPY repository ./repository

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app .

# ---------------------------------------------

FROM gcr.io/distroless/base-debian12 AS release-stage

# Copy binary
COPY --from=build-stage /api/app /app

# Run
EXPOSE 8080
ENTRYPOINT ["./app"]