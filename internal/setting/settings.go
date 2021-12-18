package setting

type Settings struct {
	SkipTokens []string
}

var DefaultSettings = Settings{
	SkipTokens: []string{"t"},
}
