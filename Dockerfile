# Golang version
FROM golang:1.21

# Set destination for COPY
WORKDIR /sup-monitor

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-sup-monitor

WORKDIR /sup-monitor

CMD ["/docker-sup-monitor"]
