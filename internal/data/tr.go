package data

import "regexp"

var trLocale = LocaleData{
	Name:                  "tr",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "ve", "yaklaşık", "|", "，"},
	January:               []string{"oc", "oca", "ocak"},
	February:              []string{"şu", "şub", "şubat"},
	March:                 []string{"mar", "mart"},
	April:                 []string{"ni", "nis", "nisan"},
	May:                   []string{"may", "mayıs"},
	June:                  []string{"ha", "haz", "haziran"},
	July:                  []string{"te", "tem", "temmuz"},
	August:                []string{"ağ", "ağu", "ağustos"},
	September:             []string{"ey", "eyl", "eylül"},
	October:               []string{"ek", "eki", "ekim"},
	November:              []string{"ka", "kas", "kasım"},
	December:              []string{"ar", "ara", "aralık"},
	Monday:                []string{"pazartesi", "pzt"},
	Tuesday:               []string{"sal", "salı"},
	Wednesday:             []string{"çar", "çarşamba", "çrs"},
	Thursday:              []string{"per", "perşembe", "prs"},
	Friday:                []string{"cum", "cuma"},
	Saturday:              []string{"cmt", "cumartesi"},
	Sunday:                []string{"paz", "pazar"},
	AM:                    []string{"öö"},
	PM:                    []string{"ös"},
	Year:                  []string{"sene", "yıl"},
	Month:                 []string{"ay"},
	Week:                  []string{"hafta", "hf"},
	Day:                   []string{"gün"},
	Hour:                  []string{"sa", "saat"},
	Minute:                []string{"dakika", "dk"},
	Second:                []string{"saniye", "sn"},
	In:                    []string{"içerisinde", "içinde", "sonra"},
	Ago:                   []string{"önce"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`bugün`},
		`0 hour ago`:   {`bu saat`},
		`0 minute ago`: {`bu dakika`},
		`0 month ago`:  {`bu ay`},
		`0 second ago`: {`şimdi`},
		`0 week ago`:   {`bu hafta`},
		`0 year ago`:   {`bu yıl`},
		`1 day ago`:    {`dün`, `geçen gün`},
		`1 month ago`:  {`geçen ay`},
		`1 week ago`:   {`geçen hafta`},
		`1 year ago`:   {`geçen yıl`},
		`in 1 day`:     {`yarın`, `önümüzdeki gün`},
		`in 1 month`:   {`gelecek ay`, `önümüzdeki ay`},
		`in 1 week`:    {`gelecek hafta`, `haftaya`, `önümüzdeki hafta`},
		`in 1 year`:    {`gelecek yıl`, `önümüzdeki yıl`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) gün önce`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) sa önce`),
			regexp.MustCompile(`(?i)(\d+) saat önce`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) dakika önce`),
			regexp.MustCompile(`(?i)(\d+) dk önce`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ay önce`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) saniye önce`),
			regexp.MustCompile(`(?i)(\d+) sn önce`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) hafta önce`),
			regexp.MustCompile(`(?i)(\d+) hf önce`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) yıl önce`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) gün sonra`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) sa sonra`),
			regexp.MustCompile(`(?i)(\d+) saat sonra`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) dakika sonra`),
			regexp.MustCompile(`(?i)(\d+) dk sonra`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ay sonra`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) saniye sonra`),
			regexp.MustCompile(`(?i)(\d+) sn sonra`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) hafta sonra`),
			regexp.MustCompile(`(?i)(\d+) hf sonra`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) yıl sonra`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"tr-CY": trCYLocale,
	},
}

var trCYLocale = LocaleData{
	Name:                  "tr-CY",
	DateOrder:             "",
}
