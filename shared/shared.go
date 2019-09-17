package shared

import (
	"fmt"
	"os"
)

const SAMPLE_SIZE = 14000 // Warm-up: 30% , Post: 10%
const GRPC_PORT = 7272
const RABBITMQ_PORT = 5672

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}

// same as fib parameter
type Request struct {
	Req int32
}

// same as fib response
type Response struct {
	Res int32
}
