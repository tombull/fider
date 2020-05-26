package crypto_test

import (
	"testing"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/crypto"
)

func TestMD5Hash(t *testing.T) {
	RegisterT(t)

	hash := crypto.MD5("Teamdream")

	Expect(hash).Equals("6a49434f5d0281f77f236f774e8659df")
}
