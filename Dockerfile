# Set go version
FROM golang:1.20.6-alpine

# No need to use GOPATH
ENV GO111MODULE=on

# Copy project files to /meli
ADD . /meli

# Set /meli as work directory
WORKDIR /meli

# Copy *.mod and *.sum files to /meli
COPY go.mod go.sum ./

# Download go dependencies
RUN go mod download

# Copy *.go files
COPY *.go ./

# Build app (production build)
RUN go build -o /meli-challenge cmd/main.go

# Expose port
EXPOSE 8080

#Run binary (production build)
CMD [ "/meli-challenge" ]

# Run code (development)
#CMD ["go", "run", "cmd/main.go"]