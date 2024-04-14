// Package implements the authentication manager.
package auth

// Manager represents authentication logic.
type Manager struct {
	accessToken string
}

// New returns a new authentication manager.
func New() *Manager {
	return &Manager{
		accessToken: "",
	}
}

// GetToken returns the current access token.
func (m *Manager) GetToken() string {
	return m.accessToken
}

// SetToken updates access token value.
func (m *Manager) SetToken(accessToken string) {
	m.accessToken = accessToken
}

// IsAuthorized checks if the user is authorized.
func (m *Manager) IsAuthorized() bool {
	return m.accessToken != ""
}
