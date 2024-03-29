package main

import "fmt"

func decodeMessage(key string, message string) string {
	d := [128]byte{}
	d[' '] = ' '
	for i, j := 0, 0; i < len(key); i++ {
		if d[key[i]] == 0 {
			d[key[i]] = byte('a' + j)
			j++
		}
	}

	ans := []byte(message)
	for i, c := range ans {
		ans[i] = d[c]
	}

	return string(ans)
}

func main() {
	fmt.Println(' ')
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}
