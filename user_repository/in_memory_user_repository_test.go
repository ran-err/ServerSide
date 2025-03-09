package user_repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"user_repository"
)

var _ = Describe("InMemoryUserRepository", func() {
	var repo *user_repository.InMemoryUserRepository

	Describe("FindUser", func() {
		When("the user does not exist", func() {
			It("should return nil", func() {
				user, err := repo.FindUserByEmail("user@example.com")
				Expect(user).To(BeNil())
				Expect(err).Error()
			})
		})
	})
})
