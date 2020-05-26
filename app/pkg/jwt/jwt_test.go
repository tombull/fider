package jwt_test

import (
	"testing"
	"time"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/jwt"
)

func TestJWT_Encode(t *testing.T) {
	RegisterT(t)

	claims := &jwt.TeamdreamClaims{
		UserID:    424,
		UserName:  "Jon Snow",
		UserEmail: "jon.snow@got.com",
	}

	token, err := jwt.Encode(claims)
	Expect(token).IsNotEmpty()
	Expect(err).IsNil()
}

func TestJWT_Decode(t *testing.T) {
	RegisterT(t)

	claims := &jwt.TeamdreamClaims{
		UserID:    424,
		UserName:  "Jon Snow",
		UserEmail: "jon.snow@got.com",
	}

	token, _ := jwt.Encode(claims)

	decoded, err := jwt.DecodeTeamdreamClaims(token)
	Expect(err).IsNil()
	Expect(decoded.UserID).Equals(claims.UserID)
	Expect(decoded.UserName).Equals(claims.UserName)
	Expect(decoded.UserEmail).Equals(claims.UserEmail)
}

func TestJWT_DecodeExpired(t *testing.T) {
	RegisterT(t)

	claims := &jwt.TeamdreamClaims{
		UserID:    424,
		UserName:  "Jon Snow",
		UserEmail: "jon.snow@got.com",
		Metadata: jwt.Metadata{
			ExpiresAt: time.Now().Unix(),
		},
	}

	token, err := jwt.Encode(claims)
	Expect(err).IsNil()
	time.Sleep(1 * time.Second)

	decoded, err := jwt.DecodeTeamdreamClaims(token)
	Expect(err).IsNotNil()
	Expect(decoded).IsNil()
}

func TestJWT_DecodeOAuthClaims(t *testing.T) {
	RegisterT(t)

	claims := &jwt.OAuthClaims{
		OAuthID:       "2",
		OAuthEmail:    "jon.snow@got.com",
		OAuthName:     "Jon Snow",
		OAuthProvider: "facebook",
	}

	token, _ := jwt.Encode(claims)

	decoded, err := jwt.DecodeOAuthClaims(token)
	Expect(err).IsNil()
	Expect(decoded.OAuthID).Equals(claims.OAuthID)
	Expect(decoded.OAuthEmail).Equals(claims.OAuthEmail)
	Expect(decoded.OAuthName).Equals(claims.OAuthName)
	Expect(decoded.OAuthProvider).Equals(claims.OAuthProvider)
}

func TestJWT_DecodeChangedToken(t *testing.T) {
	RegisterT(t)

	claims := &jwt.TeamdreamClaims{
		UserID:    424,
		UserName:  "Jon Snow",
		UserEmail: "jon.snow@got.com",
	}

	token, _ := jwt.Encode(claims)

	decoded, err := jwt.DecodeTeamdreamClaims(token + "foo")
	Expect(err).IsNotNil()
	Expect(decoded).IsNil()
}

func TestJWT_DecodeOAuthClaimsExpired(t *testing.T) {
	RegisterT(t)

	claims := &jwt.OAuthClaims{
		OAuthID:       "2",
		OAuthEmail:    "jon.snow@got.com",
		OAuthName:     "Jon Snow",
		OAuthProvider: "facebook",
		Metadata: jwt.Metadata{
			ExpiresAt: time.Now().Unix(),
		},
	}

	token, err := jwt.Encode(claims)
	Expect(err).IsNil()
	time.Sleep(1 * time.Second)

	decoded, err := jwt.DecodeOAuthClaims(token)
	Expect(err).IsNotNil()
	Expect(decoded).IsNil()
}
