package user_repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"user_repository"
)

var _ = Describe("InMemoryUserRepository", func() {
	var repo *user_repository.InMemoryUserRepository

	setupRepo := func(users ...user_repository.User) {
		repo = &user_repository.InMemoryUserRepository{Users: users}
	}

	Describe("Test Utilities", func() {
		DescribeTable("repository state inspection",
			func(users []user_repository.User, expectedLen int, expectedPeek string) {
				setupRepo(users...)
				Expect(repo.Len()).To(Equal(expectedLen))
				Expect(repo.Peek()).To(Equal(expectedPeek))
			},
			Entry("empty repository", []user_repository.User{}, 0, ""),
			Entry("repository with users",
				[]user_repository.User{
					{ID: 0, Active: false, Email: "test@test.com"},
					{ID: 1, Active: true, Email: "test@test.com"},
				},
				2,
				"{-0, test@test.com}\n{+1, test@test.com}",
			),
		)
	})

	Describe("New", func() {
		It("should return a new UserRepository instance", func() {
			repo = user_repository.New()
			Expect(repo).To(Not(BeNil()))
			Expect(repo.Len()).To(Equal(0))
			Expect(repo.Peek()).To(Equal(""))
		})
	})

	Describe("NewFromSlice", func() {
		It("should return a new UserRepository instance with predefined users", func() {
			repo = user_repository.NewFromSlice([]user_repository.User{
				{ID: 0, Active: false, Email: "test@test.com"},
				{ID: 1, Active: true, Email: "test@test.com"},
			})
			Expect(repo).To(Not(BeNil()))
			Expect(repo.Len()).To(Equal(2))
			Expect(repo.Peek()).To(ContainSubstring("test@test.com"))
		})
	})

	Describe("FindUser", func() {
		When("the user does not exist", func() {
			It("should return nil", func() {
				user, err := repo.FindUserByEmail("phantom-user@example.com")
				Expect(user).To(BeNil())
				Expect(err).Error()
			})
		})

		When("the user exists", func() {
			It("should return the user's information", func() {
				user, err := repo.FindUserByEmail("user@example.com")
				Expect(user.Email).To(Equal("user@example.com"))
				Expect(err).To(BeNil())
			})
		})
	})
})
