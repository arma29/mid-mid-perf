package shared

import (
	"fmt"
	"os"
)

const SAMPLE_SIZE = 10000
const PORT = 7272

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
