package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOtpCode() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(10000))

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%04d", n), nil
}
