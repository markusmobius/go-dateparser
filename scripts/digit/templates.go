package main

import "text/template"

var tpl = template.Must(template.New("").Parse(codeTemplate))

const codeTemplate = `
// Code is generated by script; DO NOT EDIT.

package digit

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

var rxNumeral = regexp.MustCompile("\\p{Nd}")

type RangeData struct {
	Start rune
	End   rune
}

func NormalizeString(str string) string {
	lastUsedDigit := -1

	return rxNumeral.ReplaceAllStringFunc(str, func(s string) string {
		r, _ := utf8.DecodeRuneInString(s)

		if lastUsedDigit > 0 {
			if number, ok := Ranges[lastUsedDigit].apply(r); ok {
				return strconv.Itoa(number)
			}
		}

		for i, digitRange := range Ranges {
			if number, ok := digitRange.apply(r); ok {
				lastUsedDigit = i
				return strconv.Itoa(number)
			}
		}

		return s
	})
}

func (dr RangeData) apply(r rune) (int, bool) {
	if r >= dr.Start && r <= dr.End {
		number := int(r - dr.Start)
		return number, true
	}

	return -1, false
}

var Ranges = []RangeData{
	{{range $data := . -}}
	{'{{$data.Start}}', '{{$data.End}}'}, // {{$data.Name}}
	{{end -}}
}
`