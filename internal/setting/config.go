package setting

type Configuration struct {
	SkipTokens []string
}

var DefaultConfig = Configuration{
	SkipTokens: []string{"t"},
}
