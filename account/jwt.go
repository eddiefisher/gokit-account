package account

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
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

// AuthJSONMiddleware ...
func AuthJSONMiddleware(token string) endpoint.Middleware {
	kf := func(token *jwtgo.Token) (interface{}, error) { return []byte("secret-key"), nil }
	return jwt.NewParser(kf, jwtgo.SigningMethodES256, jwt.StandardClaimsFactory)
}
