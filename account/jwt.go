package account

import (
	"fmt"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// ServiceClaims ...
type ServiceClaims struct {
	Service string `json:"service"`
	ID      string `json:"id"`
	jwtgo.StandardClaims
}

// createToken create test token
func createToken() (string, error) {
	newToken := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, ServiceClaims{
		"account",
		"some-account-id",
		jwtgo.StandardClaims{
			Issuer: "test",
		},
	})
	signingKey := []byte("gokit-secret-key")
	token, err := newToken.SignedString(signingKey)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func parseToken(token string) (*jwtgo.Token, error) {
	return jwtgo.Parse(token, func(token *jwtgo.Token) (interface{}, error) {
		return []byte("gokit-secret-key"), nil
	})
}

func validateToken(token *jwtgo.Token, err error) error {
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}

	if ve, ok := err.(*jwtgo.ValidationError); ok {
		if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
			return fmt.Errorf("That's not even a token")
		} else if ve.Errors&(jwtgo.ValidationErrorExpired|jwtgo.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return fmt.Errorf("Timing is everything")
		} else {
			return fmt.Errorf("Couldn't handle this token: %v", err)
		}
	} else {
		return fmt.Errorf("Couldn't handle this token: %v", err)
	}
}
