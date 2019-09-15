package fibonacci

func CalcFibonacci(n int32) int32 {
	if (n <= 2) {
		return 1;
	}
	
	return CalcFibonacci(n-1) + CalcFibonacci(n-2);
}