package timestamp

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cfg := &setting.DefaultConfig
	zero := time.Time{}

	// Test milliseconds timestamp
	assert.Equal(t,
		time.Unix(1570308760, 263*1_000_000),
		Parse(cfg, "1570308760263"),
	)

	// Test microseconds timestamp
	assert.Equal(t,
		time.Unix(1570308760, 263_111*1_000),
		Parse(cfg, "1570308760263111"),
	)

	// Test wrong timestamp
	assert.Equal(t, zero, Parse(cfg, "15703087602631"))
	assert.Equal(t, zero, Parse(cfg, "157030876026xx"))
	assert.Equal(t, zero, Parse(cfg, "1570308760263x"))
	assert.Equal(t, zero, Parse(cfg, "157030876026311"))
	assert.Equal(t, zero, Parse(cfg, "15703087602631x"))
	assert.Equal(t, zero, Parse(cfg, "15703087602631xx"))
	assert.Equal(t, zero, Parse(cfg, "15703087602631111"))
	assert.Equal(t, zero, Parse(cfg, "1570308760263111x"))
	assert.Equal(t, zero, Parse(cfg, "1570308760263111xx"))
	assert.Equal(t, zero, Parse(cfg, "1570308760263111222"))
}
