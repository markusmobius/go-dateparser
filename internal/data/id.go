// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var idLocale = LocaleData{
	Name:                  "id",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "pukul", "tanggal", "yang", "|", "，"},
	January:               []string{"jan", "januari"},
	February:              []string{"feb", "februari"},
	March:                 []string{"mar", "maret"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mei"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"agt", "agu", "agustus"},
	September:             []string{"sep", "sept", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"des", "desember"},
	Monday:                []string{"sen", "senin"},
	Tuesday:               []string{"sel", "selasa"},
	Wednesday:             []string{"rab", "rabu"},
	Thursday:              []string{"kam", "kamis"},
	Friday:                []string{"jum", "jumat"},
	Saturday:              []string{"sab", "sabtu"},
	Sunday:                []string{"ahad", "min", "minggu"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"tahun", "thn"},
	Month:                 []string{"bln", "bulan"},
	Week:                  []string{"mgg", "minggu"},
	Day:                   []string{"h", "hari"},
	Hour:                  []string{"j", "jam"},
	Minute:                []string{"m", "menit", "mnt"},
	Second:                []string{"d", "detik", "dtk"},
	In:                    []string{"dalam"},
	Ago:                   []string{"lalu"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hari ini`},
		`0 hour ago`:   {`jam ini`},
		`0 minute ago`: {`menit ini`},
		`0 month ago`:  {`bulan ini`},
		`0 second ago`: {`baru saja`, `sekarang`},
		`0 week ago`:   {`minggu ini`},
		`0 year ago`:   {`tahun ini`},
		`1 day`:        {`sehari`},
		`1 day ago`:    {`kemarin`},
		`1 month`:      {`sebulan`},
		`1 month ago`:  {`bulan lalu`},
		`1 week`:       {`seminggu`},
		`1 week ago`:   {`minggu lalu`},
		`1 year`:       {`setahun`},
		`1 year ago`:   {`tahun lalu`},
		`2 day ago`:    {`kemarin lusa`},
		`in 1 day`:     {`besok`},
		`in 1 month`:   {`bulan berikutnya`},
		`in 1 week`:    {`minggu depan`},
		`in 1 year`:    {`tahun depan`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) h lalu`),
			regexp.MustCompile(`(?i)(\d+) hari yang lalu`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) jam lalu`),
			regexp.MustCompile(`(?i)(\d+) jam yang lalu`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) menit yang lalu`),
			regexp.MustCompile(`(?i)(\d+) mnt lalu`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) bln lalu`),
			regexp.MustCompile(`(?i)(\d+) bulan yang lalu`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) detik yang lalu`),
			regexp.MustCompile(`(?i)(\d+) dtk lalu`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) mgg lalu`),
			regexp.MustCompile(`(?i)(\d+) minggu yang lalu`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) tahun yang lalu`),
			regexp.MustCompile(`(?i)(\d+) thn lalu`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dalam (\d+) h`),
			regexp.MustCompile(`(?i)dalam (\d+) hari`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)dalam (\d+) jam`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)dalam (\d+) menit`),
			regexp.MustCompile(`(?i)dlm (\d+) mnt`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)dalam (\d+) bulan`),
			regexp.MustCompile(`(?i)dlm (\d+) bln`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)dalam (\d+) detik`),
			regexp.MustCompile(`(?i)dlm (\d+) dtk`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)dalam (\d+) minggu`),
			regexp.MustCompile(`(?i)dlm (\d+) mgg`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)dalam (\d+) tahun`),
			regexp.MustCompile(`(?i)dlm (\d+) thn`),
		},
	},
}
