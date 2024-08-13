# Use the official Golang image as a base image
FROM golang:1.18-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o hello-world .

# Command to run the executable
CMD ["./hello-world"]