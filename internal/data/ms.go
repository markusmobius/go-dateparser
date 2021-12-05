package data

import "regexp"

var msLocale = LocaleData{
	Name:                  "ms",
	DateOrder:             "DMY",
	January:               []string{"jan", "januari"},
	February:              []string{"feb", "februari"},
	March:                 []string{"mac"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mei"},
	June:                  []string{"jun"},
	July:                  []string{"jul", "julai"},
	August:                []string{"ogo", "ogos"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"dis", "disember"},
	Monday:                []string{"isn", "isnin"},
	Tuesday:               []string{"sel", "selasa"},
	Wednesday:             []string{"rab", "rabu"},
	Thursday:              []string{"kha", "khamis"},
	Friday:                []string{"jum", "jumaat"},
	Saturday:              []string{"sab", "sabtu"},
	Sunday:                []string{"ahad", "ahd"},
	AM:                    []string{"pg"},
	PM:                    []string{"ptg"},
	Year:                  []string{"tahun", "thn"},
	Month:                 []string{"bln", "bulan"},
	Week:                  []string{"mgu", "minggu"},
	Day:                   []string{"hari"},
	Hour:                  []string{"jam"},
	Minute:                []string{"min", "minit"},
	Second:                []string{"saat"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hari ini`},
		`0 hour ago`:   {`jam ini`},
		`0 minute ago`: {`pada minit ini`},
		`0 month ago`:  {`bln ini`, `bulan ini`},
		`0 second ago`: {`sekarang`},
		`0 week ago`:   {`minggu ini`, `mng ini`},
		`0 year ago`:   {`tahun ini`, `thn ini`},
		`1 day ago`:    {`semalam`, `semlm`},
		`1 month ago`:  {`bln lalu`, `bulan lalu`},
		`1 week ago`:   {`minggu lalu`, `mng lepas`},
		`1 year ago`:   {`tahun lalu`, `thn lepas`},
		`in 1 day`:     {`esok`},
		`in 1 month`:   {`bln depan`, `bulan depan`},
		`in 1 week`:    {`minggu depan`, `mng depan`},
		`in 1 year`:    {`tahun depan`, `thn depan`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) hari lalu`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) jam lalu`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min lalu`),
			regexp.MustCompile(`(?i)(\d+) minit lalu`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) bln lalu`),
			regexp.MustCompile(`(?i)(\d+) bulan lalu`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) saat lalu`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) mgu lalu`),
			regexp.MustCompile(`(?i)(\d+) minggu lalu`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) tahun lalu`),
			regexp.MustCompile(`(?i)(\d+) thn lalu`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dalam (\d+) hari`),
			regexp.MustCompile(`(?i)dlm (\d+) hari`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)dalam (\d+) jam`),
			regexp.MustCompile(`(?i)dlm (\d+) jam`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)dalam (\d+) minit`),
			regexp.MustCompile(`(?i)dlm (\d+) min`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)dalam (\d+) bulan`),
			regexp.MustCompile(`(?i)dlm (\d+) bln`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)dalam (\d+) saat`),
			regexp.MustCompile(`(?i)dlm (\d+) saat`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)dalam (\d+) minggu`),
			regexp.MustCompile(`(?i)dlm (\d+) mgu`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)dalam (\d+) saat`),
			regexp.MustCompile(`(?i)dalam (\d+) thn`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ms-BN": msBNLocale,
		"ms-SG": msSGLocale,
	},
}

var msBNLocale = LocaleData{
	Name:                  "ms-BN",
	DateOrder:             "",
}

var msSGLocale = LocaleData{
	Name:                  "ms-SG",
	DateOrder:             "",
}
