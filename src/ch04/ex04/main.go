package main

// wakaranai
func main() {
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	gcd(b, a%b)
}
