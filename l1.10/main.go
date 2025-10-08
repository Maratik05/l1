package main

import "fmt"

func groupTemperature(arr []float64) map[int][]float64 {
	group := make(map[int][]float64)
	gr := 0
	for _, el := range arr {
		gr = int(el/10) * 10
		group[gr] = append(group[gr], el)
	}
	return group
}

func main() {
	var degrees = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Print(groupTemperature(degrees))
}
