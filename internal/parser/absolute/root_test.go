package absolute

import (
	"testing"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Helper function
	var messagePrefix string
	var cfg setting.Configuration
	expectError := func(s string) {
		_, err := Parse(&cfg, s)
		assert.Error(t, err, messagePrefix+" "+s)
	}

	// Should be error when incomplete dates given
	cfg = setting.DefaultConfig
	cfg.StrictParsing = true
	messagePrefix = "STRICT PARSING"

	expectError("april 2010")
	expectError("11 March")
	expectError("March")
	expectError("31 2010")
	expectError("31/2010")

	// Should be error when partially complete dates given
	cfg = setting.DefaultConfig
	cfg.RequiredParts = []string{"day", "month", "year"}
	messagePrefix = "INCOMPLETE DATES"

	expectError("april 2010")
	expectError("11 March")
	expectError("March")
	expectError("31 2010")
	expectError("31/2010")

	// Should be error when day part missing
	cfg = setting.DefaultConfig
	cfg.RequiredParts = []string{"day"}
	messagePrefix = "MISSING DAY"

	expectError("april 2010")
	expectError("March")
	expectError("2010")

	// Should be error when month part missing
	cfg = setting.DefaultConfig
	cfg.RequiredParts = []string{"month"}
	messagePrefix = "MISSING MONTH"

	expectError("31 2010")
	expectError("31/2010")

	// Should be error when year part missing
	cfg = setting.DefaultConfig
	cfg.RequiredParts = []string{"year"}
	messagePrefix = "MISSING YEAR"

	expectError("11 March")
	expectError("March")

}
