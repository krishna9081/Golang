func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = abs(n)
	}
	return pow(x, n)

}

func pow(x float64, n int) float64 {
	res := 1.0
	for n != 0 {
		if n%2 == 1 {
			res *= x
		} //If n is odd
		n >>= 1 // Rightshift divide the power by 2
		x *= x
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}