package domain

type Auth struct {
	Token  string
	UserID string
	Email  string
}

func NewAuth(token, userID, email string) *Auth {
	return &Auth{
		Token:  token,
		UserID: userID,
		Email:  email,
	}
}
