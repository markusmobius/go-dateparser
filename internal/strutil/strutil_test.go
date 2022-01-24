package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeDate(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text     string
		Expected string
	}

	tests := []testScenario{
		{"2005г.", "2005"},
		{"2005г. 20:37", "2005 20:37"},
		{"2005 г.", "2005"},
		{"2005 г. 15:24", "2005 15:24"},
		{"Авг.", "Авг"},
		{"2019:", "2019"},
		{"31/07/2019:", "31/07/2019"},
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" => \"%s\"", test.Text, test.Expected)

		// Sanitize text
		result := SanitizeDate(test.Text)
		passed := assert.Equal(t, test.Expected, result, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
