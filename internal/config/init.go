package config

// Initialize a globa Configuratin instance
var Config Configuration = NewConfiguration()

func init() {
	Config.Set(Debug, DebugOff)
}
