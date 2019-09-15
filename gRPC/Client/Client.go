package main;

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"fmt"
	"time"
	"github.com/arma29/mid-mid-perf/gRPC/fibonacci"
	"github.com/arma29/mid-mid-perf/shared"
)

func main() {
	var i int32

	conn, err := grpc.Dial(":"+strconv.Itoa(shared.PORT), grpc.WithInsecure())
	shared.CheckError(err)

	defer conn.Close()

	fib := fibonacci.NewFibonacciClient(conn)

	// Contacta o servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i = 0; i < shared.SAMPLE_SIZE; i++ {
		t1 := time.Now()

		// Invoca operação remota
		fib.GetFibo(ctx, &fibonacci.FibRequest{ Number: 5})

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Printf("%f \n", x)
	}
}
