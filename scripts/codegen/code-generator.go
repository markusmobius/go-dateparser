package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var (
	fnMap = template.FuncMap{
		"localeName": localeName,
		"regex":      regexPattern,
	}

	templates = map[string]*template.Template{
		"locale-map":   template.Must(template.New("").Funcs(fnMap).Parse(localeDataMapTemplate)),
		"locale-data":  template.Must(template.New("").Funcs(fnMap).Parse(localeDataTemplate)),
		"lang-order":   template.Must(template.New("").Parse(languageOrderTemplate)),
		"lang-map":     template.Must(template.New("").Parse(languageMapTemplate)),
		"lang-loc-map": template.Must(template.New("").Parse(languageLocalesMapTemplate)),
	}
)

func generateCode(tplName string, data interface{}, dstPath string) error {
	// Generate code
	b := bytes.NewBuffer(nil)
	err := templates[tplName].Execute(b, data)
	if err != nil {
		return err
	}

	// Format code
	code, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	// Save to file
	return ioutil.WriteFile(dstPath, code, os.ModePerm)
}

func generateLocaleDataCode(data LocaleData, dstPath string) error {
	// Generate code
	b := bytes.NewBuffer(nil)
	b.WriteString(`
	package data
	import "regexp"
	`)

	// Generate locale data
	err := templates["locale-data"].Execute(b, &data)
	if err != nil {
		return err
	}

	// Generate sublocale data
	for _, locale := range data.LocaleSpecific {
		err = templates["locale-data"].Execute(b, &locale)
		if err != nil {
			return err
		}
	}

	// Format code
	code, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	// Remove empty fields
	code = rxGoEmptyField.ReplaceAll(code, []byte(""))
	code = rxGoZeroField.ReplaceAll(code, []byte(""))

	// Save to file
	return ioutil.WriteFile(dstPath, code, os.ModePerm)
}

func localeName(language string) string {
	language = strings.ReplaceAll(language, "-", "")
	language = strings.ReplaceAll(language, " ", "")
	return language + "Locale"
}

func regexPattern(pattern string) string {
	return strings.ReplaceAll(pattern, "{0}", "(\\d+)")
}

const languageOrderTemplate = `
package data

var LanguageOrder = map[string]int{
	{{range $i, $language := . -}}
	"{{$language}}": {{$i}},
	{{end}}
}
`

const languageMapTemplate = `
package data

var LanguageMap = map[string][]string{
	{{range $language, $sublanguage := . -}}
	"{{$language}}": {
		{{- range $sub := $sublanguage}}
			"{{$sub}}",
		{{- end}}
	},
	{{end}}
}
`

const languageLocalesMapTemplate = `
package data

var LanguageLocalesMap = map[string][]string{
	{{range $language, $locales := . -}}
	"{{$language}}": {
		{{- range $locale := $locales}}
			"{{$locale}}",
		{{- end}}
	},
	{{end}}
}
`

const localeDataMapTemplate = `
package data

import "regexp"

type LocaleData struct {
	Name                  string                
	DateOrder             string                
	SkipWords             []string              
	PertainWords          []string              
	NoWordSpacing         bool                  
	SentenceSplitterGroup int                   
	January               []string              
	February              []string              
	March                 []string              
	April                 []string              
	May                   []string              
	June                  []string              
	July                  []string              
	August                []string              
	September             []string              
	October               []string              
	November              []string              
	December              []string              
	Monday                []string              
	Tuesday               []string              
	Wednesday             []string              
	Thursday              []string              
	Friday                []string              
	Saturday              []string              
	Sunday                []string              
	AM                    []string              
	PM                    []string              
	Decade                []string              
	Year                  []string              
	Month                 []string              
	Week                  []string              
	Day                   []string              
	Hour                  []string              
	Minute                []string              
	Second                []string              
	In                    []string              
	Ago                   []string              
	RelativeType          map[string][]string   
	RelativeTypeRegex     map[string][]*regexp.Regexp   
	Simplifications       map[string]string     
	LocaleSpecific        map[string]LocaleData 
}

var LocaleDataMap = map[string]LocaleData {
	{{range $language, $data := . -}}
	"{{$language}}": {{localeName $language}},
	{{end}}
}
`

const localeDataTemplate = `
var {{localeName .Name}} = LocaleData {
	Name:                  "{{.Name}}",
	DateOrder:             "{{.DateOrder}}",
	NoWordSpacing:         {{.NoWordSpacing}},
	SentenceSplitterGroup: {{.SentenceSplitterGroup}},
	SkipWords:    []string{ {{range $v := .SkipWords}}"{{$v}}", {{end}} },
	PertainWords: []string{ {{range $v := .PertainWords}}"{{$v}}", {{end}} },
	January:      []string{ {{range $v := .January}}"{{$v}}", {{end}} },
	February:     []string{ {{range $v := .February}}"{{$v}}", {{end}} },
	March:        []string{ {{range $v := .March}}"{{$v}}", {{end}} },
	April:        []string{ {{range $v := .April}}"{{$v}}", {{end}} },
	May:          []string{ {{range $v := .May}}"{{$v}}", {{end}} },
	June:         []string{ {{range $v := .June}}"{{$v}}", {{end}} },
	July:         []string{ {{range $v := .July}}"{{$v}}", {{end}} },
	August:       []string{ {{range $v := .August}}"{{$v}}", {{end}} },
	September:    []string{ {{range $v := .September}}"{{$v}}", {{end}} },
	October:      []string{ {{range $v := .October}}"{{$v}}", {{end}} },
	November:     []string{ {{range $v := .November}}"{{$v}}", {{end}} },
	December:     []string{ {{range $v := .December}}"{{$v}}", {{end}} },
	Monday:       []string{ {{range $v := .Monday}}"{{$v}}", {{end}} },
	Tuesday:      []string{ {{range $v := .Tuesday}}"{{$v}}", {{end}} },
	Wednesday:    []string{ {{range $v := .Wednesday}}"{{$v}}", {{end}} },
	Thursday:     []string{ {{range $v := .Thursday}}"{{$v}}", {{end}} },
	Friday:       []string{ {{range $v := .Friday}}"{{$v}}", {{end}} },
	Saturday:     []string{ {{range $v := .Saturday}}"{{$v}}", {{end}} },
	Sunday:       []string{ {{range $v := .Sunday}}"{{$v}}", {{end}} },
	AM:           []string{ {{range $v := .AM}}"{{$v}}", {{end}} },
	PM:           []string{ {{range $v := .PM}}"{{$v}}", {{end}} },
	Decade:       []string{ {{range $v := .Decade}}"{{$v}}", {{end}} },
	Year:         []string{ {{range $v := .Year}}"{{$v}}", {{end}} },
	Month:        []string{ {{range $v := .Month}}"{{$v}}", {{end}} },
	Week:         []string{ {{range $v := .Week}}"{{$v}}", {{end}} },
	Day:          []string{ {{range $v := .Day}}"{{$v}}", {{end}} },
	Hour:         []string{ {{range $v := .Hour}}"{{$v}}", {{end}} },
	Minute:       []string{ {{range $v := .Minute}}"{{$v}}", {{end}} },
	Second:       []string{ {{range $v := .Second}}"{{$v}}", {{end}} },
	In:           []string{ {{range $v := .In}}"{{$v}}", {{end}} },
	Ago:          []string{ {{range $v := .Ago}}"{{$v}}", {{end}} },
	RelativeType: map[string][]string {
		{{range $key, $values := .RelativeType -}}
		` + "`{{$key}}`" + `: { {{range $v := $values}}` + "`{{$v}}`" + `, {{end}} },
		{{end -}}
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp {
		{{range $key, $values := .RelativeTypeRegex -}}
		` + "`{{$key}}`" + `: {
			{{range $v := $values -}}
			regexp.MustCompile(` + "`(?i){{regex $v}}`" + `),
			{{end -}}
		},
		{{end -}}
	},
	Simplifications: map[string]string {
		{{range $key, $value := .Simplifications -}}
		` + "`{{regex $key}}`" + `: ` + "`{{$value}}`" + `,
		{{end -}}
	},
	LocaleSpecific: map[string]LocaleData {
		{{range $key, $data := .LocaleSpecific -}}
		"{{$key}}": {{localeName $data.Name}},
		{{end -}}
	},
}
`
