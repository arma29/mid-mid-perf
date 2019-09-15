package application

func CalcFibonacci(n int32) int32 {
	if n <= 2 {
		return 1
	} else {
		return CalcFibonacci(n-1) + CalcFibonacci(n-2)
	}
}
