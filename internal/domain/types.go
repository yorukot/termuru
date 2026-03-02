package domain

type SSHHost struct {
	Name           string
	Host           string
	Port           int
	User           string
	PrivateKeyPath string
	Description    string
	Environment    string
	AuthMethod     string
	LastConnected  string
	Fingerprint    string
	Tags           []string
}

type AppState struct {
	Hosts []SSHHost
}
