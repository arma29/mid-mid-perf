package application

func CalcFibonacci(n int32) int32 {
	if n == 0{
		return 0
	} else if n <= 2 {
		return n
	} else {
		return CalcFibonacci(n-1) + CalcFibonacci(n-2)
	}
}

