package crypto_test

import (
	"testing"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/crypto"
)

func TestSHA512Hash(t *testing.T) {
	RegisterT(t)

	hash := crypto.SHA512("Teamdream")

	Expect(hash).Equals("2692e75882eeb5dc18756979a6f6266734bd0a61744c9083dc7c5693af42f54b68c2f58dcdf2ac93e7ee5cb1251581d6735c8793817a17f80ef91e92cade3a89")
}
