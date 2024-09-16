package handlers

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator"

	"github.com/DataDog/datadog-go/statsd"
	bogcrypto "github.com/ibiscum/Building-Microservices-with-Go/crypto"

	log "github.com/sirupsen/logrus"
)

// const oneDay time.Duration = (1440 * time.Minute)

var validate = validator.New()
var defaultFields = log.Fields{
	"service": "jwt",
	"handler": "auth",
}

type LoginRequest struct {
	Username string `json:"username" validate:"email"`
	Password string `json:"password" validate:"max=36,min=8"`
}

type JWT struct {
	rsaPrivate *rsa.PrivateKey
	statsd     *statsd.Client
	logger     *log.Logger
}

// generateJWT creates a new JWT and signs it with the private key
// func (j *JWT) generateJWT(request LoginRequest) []byte {
func (j *JWT) generateJWT() []byte {
	// sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: rsaPrivate}, (&jose.SignerOptions{}).WithType("JWT"))
	// if err != nil {
	// 	fmt.Printf("making signer: %s\n", err)

	// }

	// cl := jwt.Claims{
	// 	Subject:   "subject",
	// 	Issuer:    "issuer",
	// 	NotBefore: jwt.NewNumericDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
	// 	Audience:  jwt.Audience{"leela", "fry"},
	// }

	// raw, err := jwt.Signed(sig).Claims(cl).Serialize()
	// if err != nil {
	// 	fmt.Printf("signing JWT: %s\n", err)
	// }

	// claims := jws.Claims{}
	// claims.SetExpiration(time.Now().Add(2880 * time.Minute))
	// claims.Set("userID", "abcsd232jfjf")
	// claims.Set("accessLevel", "user")

	// jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	// b, _ := jwt.Serialize(rsaPrivate)

	var raw string = "N/A"

	return []byte(raw)
}

func (j *JWT) Handle(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := j.statsd.Incr("jwt.badmethod", nil, 1)
		if err != nil {
			log.Fatal(err)
		}

		j.logger.WithFields(defaultFields).Infof("Method: %s, not allowed", r.Method)
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		err = j.statsd.Incr("jwt.badrequest", nil, 1)
		if err != nil {
			log.Fatal(err)
		}

		j.logger.WithFields(defaultFields).Errorf("Error decoding request %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		err := j.statsd.Incr("jwt.badrequest", nil, 1)
		if err != nil {
			log.Fatal(err)
		}

		j.logger.WithFields(defaultFields).Errorf("Error validating request %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	j.logger.WithFields(defaultFields).Infof("Login request from %s", request.Username)
	//jwt := j.generateJWT(request)
	jwt := j.generateJWT()

	err = j.statsd.Incr("jwt.success", nil, 1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = rw.Write(jwt)
	if err != nil {
		log.Fatal(err)
	}
}

func NewJWT(logger *log.Logger, statsd *statsd.Client) *JWT {
	var err error
	rsaPrivate, err := bogcrypto.UnmarshalRSAPrivateKeyFromFile("./sample_key.priv")
	if err != nil {
		logger.WithFields(defaultFields).Fatalf("Unable to parse private key: %v", err)
	}

	return &JWT{
		rsaPrivate: rsaPrivate,
		logger:     logger,
		statsd:     statsd,
	}
}
