// https://leetcode.com/problems/pass-the-pillow/

package main

//lint:ignore U1000 I
func passThePillow(n int, time int) int {
	mod := 2*n - 2
	rt := time % mod //cycle repeats per rt
	if rt < n {
		return rt + 1 // if it doesnt overflow just return
	} else {
		return n - (rt % (n - 1)) //if overflow go back from n reverse
	}
}

// func main() {
// 	print(passThePillow(5, 10))
// }
