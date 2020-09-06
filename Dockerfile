# Latest golang base image
FROM golang:latest

# Maintainer Info
LABEL maintainer="Rajender Koppula <koppularajender18@gmail.com>"

# Seting the Current Working Directory inside the container
WORKDIR /go/src/DuplicateFilesCheck

# Copying the Code
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Command to run the executable
ENTRYPOINT [ "DuplicateFilesCheck" ]

