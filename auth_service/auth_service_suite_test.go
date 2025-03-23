package auth_service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AuthService Suite")
}
