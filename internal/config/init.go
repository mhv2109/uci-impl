package config

var Config Configuration = NewConfiguration()

func init() {
	Config.Set(Debug, DebugOff)
}
