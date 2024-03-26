package auth

type AuthClient struct {
	accessToken string
}

func New() *AuthClient {
	return &AuthClient{
		accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJiMmQzMGViYS05ZjRlLTQwY2UtYTg3Zi01NDUyZWY4Y2U3NWYifQ.ULO0zhbUk-UbdlhPJE5KsqKT_hHHHUBfsr2894FO9jc",
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
