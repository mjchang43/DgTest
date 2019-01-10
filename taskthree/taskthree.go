package taskthree

import (
	"fmt"
	"time"
)

func Solution(s string) (string) {
	start := time.Now()
	repeat := true	
	for repeat {
		for i := (len(s)/2); i < len(s); i++ {
			if s[i-1] == s[i] {
				s = s[:(i-1)] + s[i+1:]
				break
			}else {
				if i == (len(s)-1) {
					repeat = false
				}	
			}
		}
		if len(s) == 0 || time.Now().Sub(start) >= (5 * time.Second){
			fmt.Println(s)
			break
		}
	}

	return s
}

