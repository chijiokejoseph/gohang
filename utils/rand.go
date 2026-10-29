package utils

import (
	"log"
	"math/rand/v2"

)

func RangeN(min, max int) int {
	if max <= min {
		log.Fatalf("max: %d must be greater than min: %d\n", max, min)
	}

	diff := float64(max - min)
	return min + int(rand.Float64()*diff)
}
