package interview_preparation_kit

// https://www.hackerrank.com/challenges/counting-valleys?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func countingValleys(n int32, s string) int32 {
	var valleys int32 = 0
	var level = 0
	for _, path := range s {
		if path == 'U' {
			if level == -1 {
				valleys++
			}
			level++
		} else {
			level--
		}
	}
	return valleys
}

// https://www.hackerrank.com/challenges/jumping-on-the-clouds?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func jumpingOnClouds(c []int32) int32 {
	var jumps int32 = 0
	for i := 0; i < len(c)-1; i += 2 {
		if c[i] == 1 {
			i--
		}
		jumps++
	}
	return jumps
}

// https://www.hackerrank.com/challenges/repeated-string?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func repeatedString(s string, n int64) int64 {
	var c int64 = 0
	var l = int64(len(s))
	var m = int64(n % l)
	var r int64 = n / l
	var a int64 = 0
	var as int64 = 0
	for i := int64(0); i < l; i++ {
		if s[i] == 'a' {
			a++
		}
	}
	c = a * r
	for i := int64(0); i < m; i++ {
		if s[i] == 'a' {
			as++
		}
	}
	c += as
	return c
}

// https://www.hackerrank.com/challenges/sock-merchant?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
func sockMerchant(n int32, ar []int32) int32 {
	var socks = make(map[int32]int32)
	var pairs int32 = 0
	for i := 0; i < len(ar); i++ {
		socks[ar[i]]++
	}
	for _, n := range socks {
		pairs += n / 2
	}
	return pairs
}
