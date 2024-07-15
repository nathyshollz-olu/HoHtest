#Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

#Set the working dir in the container
WORKDIR /app

#Copy the source code into the container
COPY /src /app

#Build the Go binnary
RUN go build -o /bin/go/api ./main.go

#Stage 2:Create  a lightweight Alpine linux Container and copy the binary
FROM alpine:latest

#Create the directory for the binary
RUN mkdir -p /bin/go/api

#Copy the binary from the builder stage
COPY --from=builder /bin/go/api /bin/go/api

#Set the entrypoint to the compiled binary
ENTRYPOINT ["/bin/go/api"]

