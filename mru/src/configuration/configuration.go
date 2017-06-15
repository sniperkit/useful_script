package configuration

type Configuration struct {
	DeviceName     string
	IP             string
	Port           string
	UserName       string
	Password       string
	EnablePrompt   string
	LoginPrompt    string
	PasswordPrompt string
	BasePrompt     string
	Prompt         string
	ModeDB         map[string]string
}
