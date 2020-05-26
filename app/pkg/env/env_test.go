package env_test

import (
	"testing"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/env"
)

var envs = []struct {
	go_env string
	env    string
	isEnv  func() bool
}{
	{"test", "test", env.IsTest},
	{"development", "development", env.IsDevelopment},
	{"production", "production", env.IsProduction},
	{"anything", "production", env.IsProduction},
}

func TestIsEnvironment(t *testing.T) {
	RegisterT(t)

	current := env.Config.Environment
	defer func() {
		env.Config.Environment = current
	}()

	for _, testCase := range envs {
		env.Config.Environment = testCase.go_env
		actual := testCase.isEnv()
		Expect(actual).IsTrue()
	}
}

func TestHasLegal(t *testing.T) {
	RegisterT(t)

	Expect(env.HasLegal()).IsTrue()
}

func TestMultiTenantDomain(t *testing.T) {
	RegisterT(t)

	env.Config.HostDomain = "test.fider.io"
	Expect(env.MultiTenantDomain()).Equals(".test.fider.io")
	env.Config.HostDomain = "dev.fider.io"
	Expect(env.MultiTenantDomain()).Equals(".dev.fider.io")
	env.Config.HostDomain = "fider.io"
	Expect(env.MultiTenantDomain()).Equals(".fider.io")
}

func TestIsBillingEnbled(t *testing.T) {
	RegisterT(t)

	env.Config.Stripe.SecretKey = ""
	env.Config.Stripe.PublicKey = "pk_111"
	Expect(env.IsBillingEnabled()).IsFalse()
	env.Config.Stripe.SecretKey = "sk_1234"
	Expect(env.IsBillingEnabled()).IsTrue()
}

func TestSubdomain(t *testing.T) {
	RegisterT(t)

	Expect(env.Subdomain("demo.test.assets-fider.io")).Equals("")

	env.Config.CDN.Host = "test.assets-fider.io:3000"

	Expect(env.Subdomain("demo.test.fider.io")).Equals("demo")
	Expect(env.Subdomain("demo.test.assets-fider.io")).Equals("demo")
	Expect(env.Subdomain("test.fider.io")).Equals("")
	Expect(env.Subdomain("test.assets-fider.io")).Equals("")
	Expect(env.Subdomain("helloworld.com")).Equals("")

	env.Config.HostMode = "single"

	Expect(env.Subdomain("demo.test.fider.io")).Equals("")
	Expect(env.Subdomain("demo.test.assets-fider.io")).Equals("")
	Expect(env.Subdomain("test.fider.io")).Equals("")
	Expect(env.Subdomain("test.assets-fider.io")).Equals("")
	Expect(env.Subdomain("helloworld.com")).Equals("")
}
