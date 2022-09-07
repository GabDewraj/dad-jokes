package secrets

import (
	"errors"
	"log"

	"github.com/lithammer/shortuuid"
)

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
