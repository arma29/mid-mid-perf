package main

import (
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"
	"fmt"
	"github.com/arma29/mid-mid-perf/gRPC/fibonacci"
	"github.com/arma29/mid-mid-perf/shared"
)

type fibonacciServer struct{}

func (s *fibonacciServer) GetFibo(ctx context.Context, req *fibonacci.FibRequest) (*fibonacci.FibResponse, error) {
	return &fibonacci.FibResponse{ Number: fibonacci.CalcFibonacci(req.Number) }, nil
}

func main() {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(shared.PORT))
	shared.CheckError(err)

	servidor := grpc.NewServer()
	fibonacci.RegisterFibonacciServer(servidor, &fibonacciServer{})

	fmt.Println("Servidor pronto ...")

	// Register reflection service on gRPC servidor.
	reflection.Register(servidor)

	err = servidor.Serve(conn);
	shared.CheckError(err)
}