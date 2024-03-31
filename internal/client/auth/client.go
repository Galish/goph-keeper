package auth

type AuthClient struct {
	accessToken string
}

func New() *AuthClient {
	return &AuthClient{
		accessToken: "",
	}
}

func (c *AuthClient) GetToken() string {
	return c.accessToken
}

func (c *AuthClient) SetToken(accessToken string) {
	c.accessToken = accessToken
}

func (c *AuthClient) IsAuthorized() bool {
	return c.accessToken != ""
}
