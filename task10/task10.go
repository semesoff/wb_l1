package main

import (
	"fmt"
)

func task10() map[int][]float32 {
	degree := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	hashDegree := make(map[int][]float32)

	for _, v := range degree {
		step := int(v/10) * 10
		hashDegree[step] = append(hashDegree[step], v)
	}

	return hashDegree
}

func main() {
	fmt.Println(task10())
}
