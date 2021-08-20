package main

func countDigitK(n int, k int) int {
	cnt := 0
	digit := 1
	prev := 0 // 记录低于digit位的数值 如2593 digit=100（百位）那么prev=93
	for n != 0 {
		cur_digit := n % 10
		n /= 10
		cnt += n * digit
		if cur_digit == k {
			cnt += prev + 1
		} else if cur_digit > k {
			cnt += digit
		}
		prev += cur_digit * digit
		digit *= 10
	}
	return cnt
}

func countDigitOne(n int) int {
	return countDigitK(n, 1)
}
