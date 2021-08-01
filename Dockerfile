# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

WORKDIR /valet-parking

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source to image
COPY ./ .

# Build the project
RUN go build -o .

# Run the program
CMD [ "/valet-parking/main" ]