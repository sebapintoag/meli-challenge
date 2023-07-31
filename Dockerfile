# Set go version
FROM golang:1.19
# No need to use GOPATH
ENV GO111MODULE=on
# Set work directory
WORKDIR /meli
# Copy *.mod and *.sum files to /meli
COPY go.mod go.sum ./
# Download go dependencies
RUN go mod download
# Copy *.go files
COPY *.go ./
# Build app
#RUN go build -o /cmd/meli-challenge-bin
# Expose port
EXPOSE 8080
# Run binary
#CMD [ "/cmd/meli-challenge-bin" ]
# Run code (developer)
CMD ["go", "run", "cmd/main.go"]