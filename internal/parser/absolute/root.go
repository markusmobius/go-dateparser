package absolute

import (
	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

func Parse(cfg *setting.Configuration, str string) (date.Date, error) {
	parser := NewParser(cfg, str)
	return parser.Parse()
}
