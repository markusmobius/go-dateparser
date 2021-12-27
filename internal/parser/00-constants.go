package parser

import "regexp"

var (
	rxDayNumber = regexp.MustCompile(`2.{0,3}`)
)
