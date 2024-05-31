package jwt

import (
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/go-jose/go-jose/v4/jwt"

	"github.com/go-jose/go-jose/v4"
	bgocrypto "github.com/ibiscum/Building-Microservices-with-Go/crypto"
)

var rsaPrivate *rsa.PrivateKey
var rsaPublic *rsa.PublicKey

func init() {
	var err error
	rsaPrivate, err = bgocrypto.UnmarshalRSAPrivateKeyFromFile("../keys/sample_key.priv")
	if err != nil {
		log.Fatal("Unable to parse private key", err)
	}

	rsaPublic, err = bgocrypto.UnmarshalRSAPublicKeyFromFile("../keys/sample_key.pub")
	if err != nil {
		log.Fatal("Unable to parse public key", err)
	}
}

// GenerateJWT creates a new JWT and signs it with the private key
func GenerateJWT() []byte {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: rsaPrivate}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		fmt.Printf("making signer: %s\n", err)

	}

	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "issuer",
		NotBefore: jwt.NewNumericDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
		Audience:  jwt.Audience{"leela", "fry"},
	}

	raw, err := jwt.Signed(sig).Claims(cl).Serialize()
	if err != nil {
		fmt.Printf("signing JWT: %s\n", err)
	}

	// claims := jws.Claims{}
	// claims.SetExpiration(time.Now().Add(2880 * time.Minute))
	// claims.Set("userID", "abcsd232jfjf")
	// claims.Set("accessLevel", "user")

	// jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	// b, _ := jwt.Serialize(rsaPrivate)

	return []byte(raw)
}

// ValidateJWT validates that the given slice is a valid JWT and the signature matches
// the public key
func ValidateJWT(token []byte) error {
	raw := string(token)
	tok, err := jwt.ParseSigned(raw, []jose.SignatureAlgorithm{jose.RS256})
	if err != nil {
		return fmt.Errorf("unable to parse token: %v", err)
	}

	out := jwt.Claims{}
	if err := tok.Claims(rsaPublic, &out); err != nil {
		return fmt.Errorf("unable to validate token: %v", err)
	}

	return nil
}
