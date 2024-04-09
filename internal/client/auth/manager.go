package auth

type Manager struct {
	accessToken string
}

func New() *Manager {
	return &Manager{
		accessToken: "",
	}
}

func (m *Manager) GetToken() string {
	return m.accessToken
}

func (m *Manager) SetToken(accessToken string) {
	m.accessToken = accessToken
}

func (m *Manager) IsAuthorized() bool {
	return m.accessToken != ""
}
