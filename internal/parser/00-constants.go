package parser

import "regexp"

var (
	rxDayNumber = regexp.MustCompile(`2.{0,3}`)
	rxTimestamp = regexp.MustCompile(`^(\d{10})(\d{3})?(\d{3})?(?:\.|\s|$)`)
)
