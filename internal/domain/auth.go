package domain

type Auth struct {
	Token        string
	RefreshToken string
	UserID       string
	Email        string
}

func NewAuth(token, refreshToken, userID, email string) *Auth {
	return &Auth{
		Token:        token,
		RefreshToken: refreshToken,
		UserID:       userID,
		Email:        email,
	}
}
