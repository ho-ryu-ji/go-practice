package main

import "fmt"

func main() {
	var n, taro, hanako, cnt int
	var s_t, s_h string

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s_t, &s_h)

		if len(s_t) <= len(s_h) {
			cnt = len(s_t)
		} else {
			cnt = len(s_h)
		}

		for j := 0; j < cnt; j++ {
			if s_t[j] == s_h[j] {
				if j == cnt-1 {
					if len(s_t) > len(s_h) {
						taro = taro + 3
					} else if len(s_t) < len(s_h) {
						hanako = hanako + 3
					} else {
						taro++
						hanako++
						break
					}
				} else {
					continue
				}
			} else if s_t[j] > s_h[j] {
				taro = taro + 3
				break
			} else {
				hanako = hanako + 3
				break
			}
		}
	}

	fmt.Printf("%d %d\n", taro, hanako)
}
