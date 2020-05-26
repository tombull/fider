package jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/errors"
)

var jwtSecret = env.Config.JWTSecret

// Metadata is the basic JWT information
type Metadata = jwtgo.StandardClaims

const (
	//TeamdreamClaimsOriginUI is assigned to Teamdream claims when the Auth Token is generated through the UI
	TeamdreamClaimsOriginUI = "ui"
	//TeamdreamClaimsOriginAPI is assigned to Teamdream claims when the Auth Token is generated through the API
	TeamdreamClaimsOriginAPI = "api"
)

// TeamdreamClaims represents what goes into JWT tokens
type TeamdreamClaims struct {
	UserID    int    `json:"user/id"`
	UserName  string `json:"user/name"`
	UserEmail string `json:"user/email"`
	Origin    string `json:"origin"`
	Metadata
}

// OAuthClaims represents what goes into temporary OAuth JWT tokens
type OAuthClaims struct {
	OAuthID       string `json:"oauth/id"`
	OAuthProvider string `json:"oauth/provider"`
	OAuthName     string `json:"oauth/name"`
	OAuthEmail    string `json:"oauth/email"`
	Metadata
}

// Encode creates new JWT token with given claims
func Encode(claims jwtgo.Claims) (string, error) {
	jwtToken := jwtgo.NewWithClaims(jwtgo.GetSigningMethod("HS256"), claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.Wrap(err, "failed to encode the requested claims")
	}
	return token, nil
}

// DecodeTeamdreamClaims extract claims from JWT tokens
func DecodeTeamdreamClaims(token string) (*TeamdreamClaims, error) {
	claims := &TeamdreamClaims{}
	err := decode(token, claims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode Teamdream claims")
	}
	return claims, nil
}

// DecodeOAuthClaims extract OAuthClaims from given JWT token
func DecodeOAuthClaims(token string) (*OAuthClaims, error) {
	claims := &OAuthClaims{}
	err := decode(token, claims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode OAuth claims")
	}
	return claims, nil
}

func decode(token string, claims jwtgo.Claims) error {
	jwtToken, err := jwtgo.ParseWithClaims(token, claims, func(t *jwtgo.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err == nil {
		err = claims.Valid()
	}

	if err == nil && jwtToken.Valid {
		return nil
	}

	return err
}
