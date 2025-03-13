- [ ] refactor: extract `UserRepository` interface from `InMemoryUserRepository`
- [ ] feat: add `InMemoryUserRepository` for user data storage
- [ ] feat: implement CRUD operations for `InMemoryUserRepository`

- [x] test: add `Len` and `Peek` functions for inspecting internal state
- [x] test: add helper functions and slices for creating `InMemoryUserRepository` 
- [x] test: ensure errors are checked first in tests

- [x] feat: add `New` functions for `InMemoryUserRepository`
- [ ] refactor: move `InMemoryUserRepository` to its own package

- [x] feat: return proper error when user is not found
- [x] feat: implement logic for FindUserByEmail function
- [ ] feat: add `CreateUser` function
- [ ] feat: add `UpdateUser` function
- [ ] feat: add `DeleteUser` function
- [ ] feat: implement data integrity checks

- [x] feat: add `User` struct
- [x] feat: add `ID` field to distinguish users
- [x] feat: add `Active` field to distinguish between user states
- [x] feat: store users in a slice
- [x] feat: wrap `ID`, `Active` and `Email` in dedicated types
- [ ] refactor: move `User` to its own package