package auth

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
	"time"
)

func newPasetoToken(userId string) string {
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	token.SetString("user-id", userId)
	key := paseto.NewV4SymmetricKey() // don't share this!!

	encrypted := token.V4Encrypt(key, nil)
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey := secretKey.Public()                // DO share this one

	signed := token.V4Sign(secretKey, nil)

	fmt.Println("encrypted", encrypted)
	fmt.Println("publicKey", publicKey)
	fmt.Println("signed", signed)

	return publicKey.ExportHex()
}
