package data

import "regexp"

var itLocale = LocaleData{
	Name:                  "it",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "circa", "e", "|", "，"},
	January:               []string{"gen", "gennaio"},
	February:              []string{"feb", "febbraio"},
	March:                 []string{"mar", "marzo"},
	April:                 []string{"apr", "aprile"},
	May:                   []string{"mag", "maggio"},
	June:                  []string{"giu", "giugno"},
	July:                  []string{"lug", "luglio"},
	August:                []string{"ago", "agosto"},
	September:             []string{"set", "settembre"},
	October:               []string{"ott", "ottobre"},
	November:              []string{"nov", "novembre"},
	December:              []string{"dic", "dicembre"},
	Monday:                []string{"lun", "lunedì"},
	Tuesday:               []string{"mar", "martedì"},
	Wednesday:             []string{"mer", "mercoledì"},
	Thursday:              []string{"gio", "giovedì"},
	Friday:                []string{"ven", "venerdì"},
	Saturday:              []string{"sab", "sabato"},
	Sunday:                []string{"dom", "domenica"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"anni", "anno"},
	Month:                 []string{"mese", "mesi"},
	Week:                  []string{"sett", "settimana", "settimane"},
	Day:                   []string{"g", "giorni", "giorno"},
	Hour:                  []string{"h", "ora", "ore"},
	Minute:                []string{"m", "min", "minuti", "minuto"},
	Second:                []string{"s", "sec", "secondi", "secondo"},
	In:                    []string{"in"},
	Ago:                   []string{"fa"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`oggi`},
		`0 hour ago`:   {`quest'ora`},
		`0 minute ago`: {`questo minuto`},
		`0 month ago`:  {`questo mese`},
		`0 second ago`: {`ora`},
		`0 week ago`:   {`questa settimana`},
		`0 year ago`:   {`quest'anno`},
		`1 day ago`:    {`ieri`},
		`1 month ago`:  {`mese scorso`},
		`1 week ago`:   {`settimana scorsa`},
		`1 year ago`:   {`anno scorso`},
		`2 day ago`:    {`altro ieri`},
		`in 1 day`:     {`domani`},
		`in 1 month`:   {`mese prossimo`},
		`in 1 week`:    {`settimana prossima`},
		`in 1 year`:    {`anno prossimo`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) g fa`),
			regexp.MustCompile(`(?i)(\d+) gg fa`),
			regexp.MustCompile(`(?i)(\d+) giorni fa`),
			regexp.MustCompile(`(?i)(\d+) giorno fa`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) h fa`),
			regexp.MustCompile(`(?i)(\d+) ora fa`),
			regexp.MustCompile(`(?i)(\d+) ore fa`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min fa`),
			regexp.MustCompile(`(?i)(\d+) minuti fa`),
			regexp.MustCompile(`(?i)(\d+) minuto fa`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mese fa`),
			regexp.MustCompile(`(?i)(\d+) mesi fa`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) s fa`),
			regexp.MustCompile(`(?i)(\d+) sec fa`),
			regexp.MustCompile(`(?i)(\d+) secondi fa`),
			regexp.MustCompile(`(?i)(\d+) secondo fa`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) sett fa`),
			regexp.MustCompile(`(?i)(\d+) settimana fa`),
			regexp.MustCompile(`(?i)(\d+) settimane fa`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) anni fa`),
			regexp.MustCompile(`(?i)(\d+) anno fa`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)tra (\d+) g`),
			regexp.MustCompile(`(?i)tra (\d+) gg`),
			regexp.MustCompile(`(?i)tra (\d+) giorni`),
			regexp.MustCompile(`(?i)tra (\d+) giorno`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)tra (\d+) h`),
			regexp.MustCompile(`(?i)tra (\d+) ora`),
			regexp.MustCompile(`(?i)tra (\d+) ore`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)tra (\d+) min`),
			regexp.MustCompile(`(?i)tra (\d+) minuti`),
			regexp.MustCompile(`(?i)tra (\d+) minuto`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)tra (\d+) mese`),
			regexp.MustCompile(`(?i)tra (\d+) mesi`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)tra (\d+) s`),
			regexp.MustCompile(`(?i)tra (\d+) sec`),
			regexp.MustCompile(`(?i)tra (\d+) secondi`),
			regexp.MustCompile(`(?i)tra (\d+) secondo`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)tra (\d+) sett`),
			regexp.MustCompile(`(?i)tra (\d+) settimana`),
			regexp.MustCompile(`(?i)tra (\d+) settimane`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)tra (\d+) anni`),
			regexp.MustCompile(`(?i)tra (\d+) anno`),
		},
	},
	Simplifications: map[string]string{
		`(\d+)\s+ora`: `\1 ore`,
	},
	LocaleSpecific: map[string]LocaleData{
		"it-CH": itCHLocale,
		"it-SM": itSMLocale,
		"it-VA": itVALocale,
	},
}

var itCHLocale = LocaleData{
	Name:                  "it-CH",
	DateOrder:             "",
}

var itSMLocale = LocaleData{
	Name:                  "it-SM",
	DateOrder:             "",
}

var itVALocale = LocaleData{
	Name:                  "it-VA",
	DateOrder:             "",
}
