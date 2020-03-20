# Middleware - Middleware - Performance Evaluation

A simple project to evaluate the gRPC and RabbitMQ performance. The client request a number and the server returns the respective fibonacci element. Implementations built in Go. 

You can either check/run the code or run the tests to generate .csv files into your machine. Feel free to analyse the data any way you want.

## To run the code

### Prerequisites 

- Go installed
- RabbitMQ installed

From [Golang - Getting Started](https://golang.org/doc/install)

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

```bash
$ export PATH=$PATH:/usr/local/go/bin
```

For convenience, add the workspace's bin subdirectory to your `PATH`:

```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

Also a good idea add these lines:

```bash
$ export GOPATH=$(go env GOPATH);
$ export PATH=$PATH:$GOPATH/bin;
```

### Install via `go get`

```bash
$ go get github.com/arma29/mid-mid-perf/...
```

This command will build and install all go programs under `mid-mid-perf/` directory. It then installs that binaries to the workspace's bin directory as the folder name.

### Usage

Now, you can run the program by typing its full path at the command line:

```bash
$ $GOPATH/bin/bin_name
```

Or, as you have added `$GOPATH/bin` to your `PATH`, just type the binary name:

```bash
$ bin_name
```
For the RabbitMQ, we need firstly start the rabbitmq server, then server, finally the client. For gRPC only need the server (first) and client.

- RabbitMQ Server listening port: 5672
- gRPC Server listening port: 7272

To run the application servers:
```bash
$ $GOPATH/bin/ServerR 127.0.0.1
$ $GOPATH/bin/ServerGRPC
```
To run the clients, we need to pass the server IP and the request number as argument e.g:

```bash
$ $GOPATH/bin/ClientGRPC 127.0.0.1 5
$ $GOPATH/bin/ClientR 127.0.0.1 5
```

This will print 10K lines containing the fibonacci number requested, the answer and the measured time between the request and response, in milliseconds.

## To run the performance tests

### Prerequisites

- Docker
- docker-compose

### Usage

Inside `docker-test/` directory, simply run:

```bash
$ ./suit.sh
```

It takes 40 minutes to run everything and will output in `/home/$USER/Output` the .csv files in this format:

| Fibonacci | Answer | Time |
| -------- | -------- | -------- |
| requested number     | Answer number     | RTT time     |
