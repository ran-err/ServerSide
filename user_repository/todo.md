- [ ] refactor: extract `UserRepository` interface from `InMemoryUserRepository`
- [ ] feat: add `InMemoryUserRepository` for user data storage
- [ ] feat: implement CRUD operations for `InMemoryUserRepository`

- [x] test: add `Len` and `Peek` functions for inspecting internal state
-
- [ ] feat: add `New` functions for `InMemoryUserRepository`

- [x] feat: return proper error when user is not found

- [x] feat: add `User` struct
- [x] feat: add `ID` field to distinguish users
- [x] feat: add `Active` field to distinguish between user states
- [x] feat: store users in a slice
- [ ] refactor: move `User` to its own package