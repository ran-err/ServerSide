package auth_service_test

import (
	"auth_service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type mockClaims struct {
	userID string
}

func (c *mockClaims) GetUserID() string { return c.userID }

var _ = Describe("AuthJWT", func() {
	var (
		authService auth_service.AuthService
		secretKey   = "secretKey"
		userID      = "userID"
	)

	BeforeEach(func() {
		authService = auth_service.NewAuthServiceJWT(secretKey)
	})

	Describe("GenerateToken and ParseToken", func() {
		It("should generate a valid token and parse it correctly", func() {
			claims := &mockClaims{userID}
			token, err := authService.GenerateToken(claims)
			Expect(err).To(BeNil())
			Expect(token).NotTo(BeEmpty())

			parsedClaims, err := authService.ParseToken(token)
			Expect(err).To(BeNil())
			Expect(parsedClaims.GetUserID()).To(Equal(userID))
		})

		It("should fail to parse an invalid token", func() {
			_, err := authService.ParseToken("not.a.valid.token")
			Expect(err).ToNot(BeNil())
		})
	})
})
