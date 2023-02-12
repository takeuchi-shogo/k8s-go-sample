package services

import (
	"strconv"
	"strings"

	"github.com/takeuchi-shogo/k8s-go-sample/utils"
)

func CreatePINCode(max int) string {
	randomInts := []string{}
	for i := 0; i < max; i++ {
		randomInt := utils.RandomInt(9)
		randomInts = append(randomInts, strconv.Itoa(randomInt))
	}
	return strings.Join(randomInts, ",")
}
