FROM golang:1.16

# Select working directory
WORKDIR /test

# Copy all files and directories into the working directory
COPY . .

# Download necessary Go modules
RUN go mod download

# Run Go tests (includes sub directories)
CMD [ "go", "test", "./..." ]
