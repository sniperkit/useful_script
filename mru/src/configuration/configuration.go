package configuration

type Configuration struct {
	Device         string
	IP             string
	Port           string
	Username       string
	Password       string
	EnablePrompt   string
	LoginPrompt    string
	PasswordPrompt string
	BasePrompt     string
	Prompt         string
	ModeDB         map[string]string
	Name           string
}
