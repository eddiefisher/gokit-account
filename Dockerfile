############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
# Fetch dependencies.
# Using go get.
#RUN go get -d -v
# Build the binary.
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o ./build/app -v ./
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
COPY --from=builder /app/build/app /app/app
# Run the hello binary.
ENTRYPOINT ["/app/app"]
