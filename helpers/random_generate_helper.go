package helpers

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type randomGenerator struct {
	rand *rand.Rand
}

func NewRandomGenerator() *randomGenerator {
	return &randomGenerator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rg *randomGenerator) RandomInt(min, max int64) int64 {
	return min + rg.rand.Int63n(max-min+1)
}

func (rg *randomGenerator) RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = alphabet[rg.rand.Intn(len(alphabet))]
	}
	return string(result)
}

func (rg *randomGenerator) RandomValue(fieldName string) interface{} {
	switch fieldName {
	case "Username":
		return rg.RandomString(8)
	case "Email":
		return rg.RandomString(8) + "@" + rg.RandomString(6) + ".com"
	case "Phone":
		return rg.RandomString(int(rg.RandomInt(10, 13)))
	case "Name":
		return rg.RandomString(10)
	case "Address":
		return rg.RandomString(10)
	case "Age":
		return uint(rg.RandomInt(1, 100))
	case "Gender":
		genders := []string{"Male", "Female", "Other"}
		return genders[rg.rand.Intn(len(genders))]
	default:
		return nil
	}
}
