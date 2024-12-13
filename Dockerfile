# Thank you https://betterstack.com/community/guides/scaling-go/dockerize-golang/ for the tutorial
# Parent image
FROM golang:1.23-bookworm AS build

# Use /build as workdir
WORKDIR /build

# Copy go.mod and go.sum to work directory (/build)
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source code from . (current dir) to . (work dir --> /build)
COPY . .

# Build the application
RUN go build -o server

# Listen to port 8000
EXPOSE 8000

# Start the application
CMD ["/build/server"]
