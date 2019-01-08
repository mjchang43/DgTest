package tasktwo

import (

)

type Dispenser struct {
	Liters int
	Filling bool
	LastRecord int
}

func Solution(A []int, X, Y, Z int) int {
	dispenserX := Dispenser{X, false, 0}
	dispenserY := Dispenser{Y, false, 0}
	dispenserZ := Dispenser{Z, false, 0}

	maxWaiting := -1

	for i, value := range A {
		
		if dispenserX.Liters >= value {
			if !dispenserX.Filling {
				dispenserX.Filling = true
				dispenserX.LastRecord = value
			}else {
				if dispenserX.LastRecord > maxWaiting {
					maxWaiting = dispenserX.LastRecord
				}
			}
			dispenserX.Liters -= value
			continue
		}

		if dispenserY.Liters >= value {
			if !dispenserY.Filling {
				dispenserY.Filling = true
				dispenserY.LastRecord = value
			}else {
				if dispenserY.LastRecord > maxWaiting {
					maxWaiting = dispenserY.LastRecord
				}
			}
			dispenserY.Liters -= value
			continue
		}

		if dispenserZ.Liters >= value {
			if !dispenserZ.Filling {
				dispenserZ.Filling = true
				dispenserZ.LastRecord = value
			}else {
				if dispenserZ.LastRecord > maxWaiting {
					maxWaiting = dispenserZ.LastRecord
				}
			}
			dispenserZ.Liters -= value
			continue
		}
		
		if i == 0 {
			break
		}
	}

	return maxWaiting
}
