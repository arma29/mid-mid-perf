# Official golang as parent image
FROM golang:1.12

# Set working directory to /go/src/app
WORKDIR /go/src/app

# Copy everything to container (only need dockerfile)
COPY . .

# Installing requirements 
RUN go get github.com/arma29/mid-mid-perf/...

# # Execute the binary when the container launches
# CMD ["ServerUDP"]