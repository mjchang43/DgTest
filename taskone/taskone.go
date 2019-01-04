package taskone

import (

)

func Solution(A []int) int {
	max := 0
	turnOnList := make([]bool, len(A))
	shineMomentList := []int{}
	for i, value := range A {
		turnOnList[value - 1] = true
		if value > max {
			max = value
		}
		for j := 0; j < max; j++ {
			if !turnOnList[j] {
				break
			} else if j == max-1 {
				shineMomentList = append(shineMomentList, i+1)
			} 
		}
	}
	return len(shineMomentList)
}


