package auth

type AuthManager struct {
	accessToken string
}

func New() *AuthManager {
	return &AuthManager{
		accessToken: "",
	}
}

func (a *AuthManager) GetToken() string {
	return a.accessToken
}

func (a *AuthManager) SetToken(accessToken string) {
	a.accessToken = accessToken
}

func (a *AuthManager) IsAuthorized() bool {
	return a.accessToken != ""
}
