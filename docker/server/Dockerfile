# Telling to use Docker's golang ready image
FROM golang

# Set pwd to the go folder
WORKDIR /go/src/app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Copy the entire project directory into the container
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

EXPOSE 8080 443 80
CMD ["./myapp"]