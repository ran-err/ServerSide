package auth_service

type Claims interface {
	GetUserID() string
}

type AuthService interface {
	GenerateToken(claims Claims) (string, error)
	ParseToken(tokenStr string) (Claims, error)
}
