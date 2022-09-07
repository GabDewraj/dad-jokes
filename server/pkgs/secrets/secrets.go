package secrets

import (
	"errors"
	"log"

	"github.com/lithammer/shortuuid"
)

// Make shift in memory cache

var vault []string

func AddSecret() string {
	key := shortuuid.New()
	vault = append(vault, key)
	log.Print(vault)
	return key
}

func VerifySecret(keyFromHeader string) error {
	for _, secret := range vault {
		if secret == keyFromHeader {
			return nil
		}
	}
	return errors.New("secret not found for key")
}
