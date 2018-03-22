package ex06

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount24(x uint64) int {
	var ans int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			ans += int(byte(x >> uint(i*8+j) & 1))
		}
	}
	return ans
}

func PopCount25(x uint64) int {
	var ans int
	for int(x) != 0 {
		x = uint64(byte(x & (x - 1)))
		ans++
	}
	return ans
}
