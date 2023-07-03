# First stage: build the executable
FROM golang:1.20-alpine as golang

# Enable Go modules and vendor
ENV GO111MODULE=on GOFLAGS=-mod=vendor

# Set the Current Working Directory inside the Container
WORKDIR /app

# Copy go.mod and go.sum to the working directory inside the container
COPY go.mod go.sum ./

# Run go mod vendor
RUN go mod vendor

# Copy all the source the current directory to the working directory inside the container
COPY . .

# Build the Go App
RUN CG0_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o server .

# Final stage: the running container
FROM gcr.io/distroless/static-debian11

# Set the Current Working Directory inside the Container
WORKDIR /app

# Copy binary from first stage to final stage
COPY --from=golang /app/server ./
COPY --from=golang /app/vendor ./vendor

ENTRYPOINT ["./server", "-port=3000"]