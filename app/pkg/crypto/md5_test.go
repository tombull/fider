package crypto_test

import (
	"testing"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/crypto"
)

func TestMD5Hash(t *testing.T) {
	RegisterT(t)

	hash := crypto.MD5("Teamdream")

	Expect(hash).Equals("3734538c8b650e4f354a55a436566bb6")
}
