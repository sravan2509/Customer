# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application
RUN go build -o main .

# Expose port 8000 for the API
EXPOSE 8000

ENV GOPATH=/app

# Run the command to start the API
CMD ["./main"]



