package tokens

import (
	"context"
	"hradec/internal/domain"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	expirationClaim = "expiration"
	idClaim         = "id"
	roleClaim       = "role"
	nameClaim       = "name"
	surnameClaim    = "surname"
	groupClaim      = "group"
)

type TokenGenerator struct {
	JWTSecret     string
	JWTExpiration time.Duration
}

func NewTokenGenerator(jwtSecret string, JWTExpiration time.Duration) *TokenGenerator {
	return &TokenGenerator{
		JWTSecret:     jwtSecret,
		JWTExpiration: JWTExpiration,
	}
}

func (au *TokenGenerator) CreateUserJWT(ctx context.Context, usr domain.User) (string, error) {
	log.Println("Creating JWT token")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[idClaim] = usr.ID

	claims[expirationClaim] = time.Now().UTC().Add(au.JWTExpiration).Unix()

	resultToken, err := token.SignedString([]byte(au.JWTSecret))
	if err != nil {
		log.Println("Error creating JWT token")
		return "", err
	}
	log.Println("Successfully created JWT token")
	return resultToken, nil
}

func (au *TokenGenerator) CreateEmailValidateJWT(ctx context.Context, userId int64) (string, error) {
	log.Println("Creating JWT token")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[idClaim] = userId
	claims[expirationClaim] = time.Now().UTC().Add(au.JWTExpiration).Unix()

	resultToken, err := token.SignedString([]byte(au.JWTSecret))
	if err != nil {
		log.Println("Error creating JWT token")
		return "", err
	}
	log.Println("Successfully created JWT token")
	return resultToken, nil
}

func (au *TokenGenerator) ValidateEmailJSON(ctx context.Context, token string) (int64, error) {
	log.Println("Validating JWT token")
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Panicln("Error parsing JWT token")
			return nil, domain.Unauthorized
		}
		return []byte(au.JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if parsedToken == nil {
		log.Println("Error parsing JWT token")
		return 0, domain.Unauthorized
	}

	log.Println("Successfully parsed JWT token")
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Error parsing claims")
		return 0, domain.Unauthorized
	}

	userID, ok := claims[idClaim].(float64)
	if !ok {
		log.Println("Error parsing user ID")
		return 0, domain.Unauthorized
	}

	//check expiration
	expiration, ok := claims[expirationClaim].(float64)
	if !ok {
		log.Println("Error parsing expiration")
		return 0, domain.Unauthorized
	}

	if int64(expiration) < time.Now().UTC().Unix() {
		log.Println("Error parsing expiration")
		return 0, domain.Unauthorized
	}

	log.Println("Successfully parsed JWT token")
	return int64(userID), nil
}

func (au *TokenGenerator) ValidateUserJWT(ctx context.Context, token string) (int64, error) {
	log.Println("Authenticating JWT token")
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, domain.Unauthorized
		}
		return []byte(au.JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if parsedToken == nil {
		return 0, domain.Unauthorized
	}

	log.Println("Successfully parsed JWT token")
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Error parsing claims")
		return 0, domain.Unauthorized
	}

	userID, ok := claims[idClaim].(float64)
	if !ok {
		log.Println("Error parsing user ID")
		return 0, domain.Unauthorized
	}
	return int64(userID), nil
}
