package user_repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"user_repository"
)

var _ = Describe("InMemoryUserRepository", func() {
	var repo *user_repository.InMemoryUserRepository

	createUser := func(id int, active bool, email string) user_repository.User {
		return user_repository.NewUser(id, active, email)
	}
	setupRepo := func(users ...user_repository.User) {
		repo = &user_repository.InMemoryUserRepository{Users: users}
	}
	emptyRepo := []user_repository.User{}
	basicRepo := []user_repository.User{
		createUser(0, false, "test@test.com"),
		createUser(1, true, "test@test.com"),
		createUser(2, true, "user@example.com"),
	}

	Describe("Test Utilities", func() {
		DescribeTable("repository state inspection",
			func(users []user_repository.User, expectedLen int, expectedPeek string) {
				setupRepo(users...)
				Expect(repo.Len()).To(Equal(expectedLen))
				Expect(repo.Peek()).To(Equal(expectedPeek))
			},
			Entry("empty repository", emptyRepo, len(emptyRepo), ""),
			Entry("repository with users", basicRepo, len(basicRepo),
				"{-0, test@test.com}\n{+1, test@test.com}\n{+2, user@example.com}"),
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
			repo = user_repository.NewFromSlice(basicRepo)
			Expect(repo).To(Not(BeNil()))
			Expect(repo.Len()).To(Equal(len(basicRepo)))
			Expect(repo.Peek()).To(ContainSubstring("test@test.com"))
		})
	})

	Describe("FindUser", func() {
		BeforeEach(func() {
			setupRepo(basicRepo...)
		})

		When("the user does not exist", func() {
			It("should return nil", func() {
				user, err := repo.FindUserByEmail("phantom-user@example.com")
				Expect(err).Error()
				Expect(user).To(BeNil())
			})
		})

		When("the user exists", func() {
			It("should return the user's information", func() {
				user, err := repo.FindUserByEmail("user@example.com")
				Expect(err).To(BeNil())
				Expect(user.Email).To(Equal("user@example.com"))
			})
		})

		When("multiple users share the same email, but only one is active", func() {
			It("should return the active user's information", func() {
				user, err := repo.FindUserByEmail("test@test.com")

				Expect(err).To(BeNil())
				Expect(user.Email).To(Equal("test@test.com"))
				Expect(user.Active).To(BeTrue())
			})
		})
	})
})
