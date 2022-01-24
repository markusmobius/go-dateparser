package absolute

import (
	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

func Parse(cfg *setting.Configuration, str string) (date.Date, error) {
	parser, err := NewParser(cfg, str)
	if err != nil {
		return date.Date{}, err
	}
	return parser.Parse()
}
