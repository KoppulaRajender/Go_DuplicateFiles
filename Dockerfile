# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Rajender Koppula <koppularajender18@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/duplicates
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
# Command to run the executable
ENTRYPOINT [ "duplicates" ] 