package main;

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"fmt"
	"time"
	"os"

	"github.com/arma29/mid-mid-perf/application"
	"github.com/arma29/mid-mid-perf/shared"
)

func main() {
	// Get Argument from command Line
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	ipContainer := os.Args[1]
	var i int32

	conn, err := grpc.Dial(ipContainer + ":" + 
		strconv.Itoa(shared.GRPC_PORT), grpc.WithInsecure())
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
		fmt.Println(x)
	}
}
