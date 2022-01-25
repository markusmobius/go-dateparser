package timestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Prepare variables and helper function
	zero := time.Time{}
	parse := func(s string) time.Time {
		dt := Parse(nil, s)
		return dt.Time
	}

	// Test milliseconds timestamp
	expected := time.Unix(1570308760, 263*1_000_000)
	assert.Equal(t, expected, parse("1570308760263"))

	// Test microseconds timestamp
	expected = time.Unix(1570308760, 263_111*1_000)
	assert.Equal(t, expected, parse("1570308760263111"))

	// Test wrong timestamp
	assert.Equal(t, zero, parse("15703087602631"))
	assert.Equal(t, zero, parse("157030876026xx"))
	assert.Equal(t, zero, parse("1570308760263x"))
	assert.Equal(t, zero, parse("157030876026311"))
	assert.Equal(t, zero, parse("15703087602631x"))
	assert.Equal(t, zero, parse("15703087602631xx"))
	assert.Equal(t, zero, parse("15703087602631111"))
	assert.Equal(t, zero, parse("1570308760263111x"))
	assert.Equal(t, zero, parse("1570308760263111xx"))
	assert.Equal(t, zero, parse("1570308760263111222"))
}
