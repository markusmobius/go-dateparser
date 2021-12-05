package data

import "regexp"

var frLocale = LocaleData{
	Name:                  "fr",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "environ", "er", "et", "le", "|", "à", "，"},
	January:               []string{"jan", "janv", "janvier"},
	February:              []string{"fév", "févr", "février"},
	March:                 []string{"mars"},
	April:                 []string{"avr", "avril"},
	May:                   []string{"mai"},
	June:                  []string{"juin", "jun"},
	July:                  []string{"juil", "juillet", "jul"},
	August:                []string{"aoû", "août"},
	September:             []string{"sep", "sept", "septembre"},
	October:               []string{"oct", "octobre"},
	November:              []string{"nov", "novembre"},
	December:              []string{"déc", "décembre"},
	Monday:                []string{"lu", "lun", "lundi"},
	Tuesday:               []string{"ma", "mar", "mardi"},
	Wednesday:             []string{"me", "mer", "mercredi"},
	Thursday:              []string{"je", "jeu", "jeudi"},
	Friday:                []string{"ve", "ven", "vendredi"},
	Saturday:              []string{"sa", "sam", "samedi"},
	Sunday:                []string{"di", "dim", "dimanche"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"a", "an", "année", "années", "ans"},
	Month:                 []string{"m", "mois"},
	Week:                  []string{"sem", "semaine", "semaines"},
	Day:                   []string{"j", "jour", "jours"},
	Hour:                  []string{"h", "heure", "heures"},
	Minute:                []string{"min", "minute", "minutes"},
	Second:                []string{"s", "seconde", "secondes"},
	In:                    []string{"après", "dans", "en"},
	Ago:                   []string{"il y a", "il ya"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`aujourd'hui`},
		`0 hour ago`:   {`cette heure-ci`},
		`0 minute ago`: {`cette minute-ci`},
		`0 month ago`:  {`ce mois-ci`},
		`0 second ago`: {`maintenant`},
		`0 week ago`:   {`cette semaine`},
		`0 year ago`:   {`cette année`},
		`1 day ago`:    {`hier`},
		`1 month ago`:  {`le mois dernier`},
		`1 week ago`:   {`la semaine dernière`},
		`1 year ago`:   {`l'année dernière`},
		`2 day ago`:    {`avant-hier`},
		`in 1 day`:     {`demain`},
		`in 1 month`:   {`le mois prochain`},
		`in 1 week`:    {`la semaine prochaine`},
		`in 1 year`:    {`l'année prochaine`},
		`in 2 day`:     {`après-demain`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) j`),
			regexp.MustCompile(`(?i)il y a (\d+) jour`),
			regexp.MustCompile(`(?i)il y a (\d+) jours`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) h`),
			regexp.MustCompile(`(?i)il y a (\d+) heure`),
			regexp.MustCompile(`(?i)il y a (\d+) heures`),
			regexp.MustCompile(`(?i)il y a (\d+)h`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) min`),
			regexp.MustCompile(`(?i)il y a (\d+) minute`),
			regexp.MustCompile(`(?i)il y a (\d+) minutes`),
			regexp.MustCompile(`(?i)il y a (\d+)min`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) m`),
			regexp.MustCompile(`(?i)il y a (\d+) mois`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) s`),
			regexp.MustCompile(`(?i)il y a (\d+) seconde`),
			regexp.MustCompile(`(?i)il y a (\d+) secondes`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) sem`),
			regexp.MustCompile(`(?i)il y a (\d+) semaine`),
			regexp.MustCompile(`(?i)il y a (\d+) semaines`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)il y a (\d+) a`),
			regexp.MustCompile(`(?i)il y a (\d+) an`),
			regexp.MustCompile(`(?i)il y a (\d+) ans`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dans (\d+) j`),
			regexp.MustCompile(`(?i)dans (\d+) jour`),
			regexp.MustCompile(`(?i)dans (\d+) jours`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)dans (\d+) h`),
			regexp.MustCompile(`(?i)dans (\d+) heure`),
			regexp.MustCompile(`(?i)dans (\d+) heures`),
			regexp.MustCompile(`(?i)dans (\d+)h`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)dans (\d+) min`),
			regexp.MustCompile(`(?i)dans (\d+) minute`),
			regexp.MustCompile(`(?i)dans (\d+) minutes`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)dans (\d+) m`),
			regexp.MustCompile(`(?i)dans (\d+) mois`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)dans (\d+) s`),
			regexp.MustCompile(`(?i)dans (\d+) seconde`),
			regexp.MustCompile(`(?i)dans (\d+) secondes`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)dans (\d+) sem`),
			regexp.MustCompile(`(?i)dans (\d+) semaine`),
			regexp.MustCompile(`(?i)dans (\d+) semaines`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)dans (\d+) a`),
			regexp.MustCompile(`(?i)dans (\d+) an`),
			regexp.MustCompile(`(?i)dans (\d+) ans`),
		},
	},
	Simplifications: map[string]string{
		`(\d+)\s+h\s+(\d+)\s+min`: `\1h\2m`,
		`(\d+)h(\d+)m?`:           `\1:\2`,
		`cinq`:                    `5`,
		`d'un`:                    `1`,
		`d'une`:                   `1`,
		`deux`:                    `2`,
		`dix`:                     `10`,
		`douze`:                   `12`,
		`huit`:                    `8`,
		`moins\s(?:de\s)?(\d+)\s?(minute|seconde|heure)`: `\1 \2`,
		`moins\s(?:de\s)?(\d+)\s?h`:                      `\1 heure`,
		`moins\s(?:de\s)?(\d+)\s?m`:                      `\1 minute`,
		`moins\s(?:de\s)?(\d+)\s?s`:                      `\1 seconde`,
		`neuf`:                                           `9`,
		`onze`:                                           `11`,
		`quatre`:                                         `4`,
		`sept`:                                           `7`,
		`six`:                                            `6`,
		`trois`:                                          `3`,
		`un`:                                             `1`,
		`une`:                                            `1`,
	},
	LocaleSpecific: map[string]LocaleData{
		"fr-BE": frBELocale,
		"fr-BF": frBFLocale,
		"fr-BI": frBILocale,
		"fr-BJ": frBJLocale,
		"fr-BL": frBLLocale,
		"fr-CA": frCALocale,
		"fr-CD": frCDLocale,
		"fr-CF": frCFLocale,
		"fr-CG": frCGLocale,
		"fr-CH": frCHLocale,
		"fr-CI": frCILocale,
		"fr-CM": frCMLocale,
		"fr-DJ": frDJLocale,
		"fr-DZ": frDZLocale,
		"fr-GA": frGALocale,
		"fr-GF": frGFLocale,
		"fr-GN": frGNLocale,
		"fr-GP": frGPLocale,
		"fr-GQ": frGQLocale,
		"fr-HT": frHTLocale,
		"fr-KM": frKMLocale,
		"fr-LU": frLULocale,
		"fr-MA": frMALocale,
		"fr-MC": frMCLocale,
		"fr-MF": frMFLocale,
		"fr-MG": frMGLocale,
		"fr-ML": frMLLocale,
		"fr-MQ": frMQLocale,
		"fr-MR": frMRLocale,
		"fr-MU": frMULocale,
		"fr-NC": frNCLocale,
		"fr-NE": frNELocale,
		"fr-PF": frPFLocale,
		"fr-PM": frPMLocale,
		"fr-RE": frRELocale,
		"fr-RW": frRWLocale,
		"fr-SC": frSCLocale,
		"fr-SN": frSNLocale,
		"fr-SY": frSYLocale,
		"fr-TD": frTDLocale,
		"fr-TG": frTGLocale,
		"fr-TN": frTNLocale,
		"fr-VU": frVULocale,
		"fr-WF": frWFLocale,
		"fr-YT": frYTLocale,
	},
}

var frHTLocale = LocaleData{
	Name:                  "fr-HT",
	DateOrder:             "",
	Day:                   []string{"jr"},
	Hour:                  []string{"hr"},
	Second:                []string{"sec"},
}

var frRELocale = LocaleData{
	Name:                  "fr-RE",
	DateOrder:             "",
}

var frYTLocale = LocaleData{
	Name:                  "fr-YT",
	DateOrder:             "",
}

var frGNLocale = LocaleData{
	Name:                  "fr-GN",
	DateOrder:             "",
}

var frCGLocale = LocaleData{
	Name:                  "fr-CG",
	DateOrder:             "",
}

var frMFLocale = LocaleData{
	Name:                  "fr-MF",
	DateOrder:             "",
}

var frDZLocale = LocaleData{
	Name:                  "fr-DZ",
	DateOrder:             "",
}

var frMCLocale = LocaleData{
	Name:                  "fr-MC",
	DateOrder:             "",
}

var frTGLocale = LocaleData{
	Name:                  "fr-TG",
	DateOrder:             "",
}

var frPMLocale = LocaleData{
	Name:                  "fr-PM",
	DateOrder:             "",
}

var frGALocale = LocaleData{
	Name:                  "fr-GA",
	DateOrder:             "",
}

var frCMLocale = LocaleData{
	Name:                  "fr-CM",
	DateOrder:             "",
	AM:                    []string{"mat", "matin"},
	PM:                    []string{"soir"},
}

var frSNLocale = LocaleData{
	Name:                  "fr-SN",
	DateOrder:             "",
}

var frMLLocale = LocaleData{
	Name:                  "fr-ML",
	DateOrder:             "",
}

var frMRLocale = LocaleData{
	Name:                  "fr-MR",
	DateOrder:             "",
}

var frCDLocale = LocaleData{
	Name:                  "fr-CD",
	DateOrder:             "",
}

var frTNLocale = LocaleData{
	Name:                  "fr-TN",
	DateOrder:             "",
}

var frBFLocale = LocaleData{
	Name:                  "fr-BF",
	DateOrder:             "",
}

var frGPLocale = LocaleData{
	Name:                  "fr-GP",
	DateOrder:             "",
}

var frPFLocale = LocaleData{
	Name:                  "fr-PF",
	DateOrder:             "",
}

var frRWLocale = LocaleData{
	Name:                  "fr-RW",
	DateOrder:             "",
}

var frBJLocale = LocaleData{
	Name:                  "fr-BJ",
	DateOrder:             "",
}

var frGFLocale = LocaleData{
	Name:                  "fr-GF",
	DateOrder:             "",
}

var frBELocale = LocaleData{
	Name:                  "fr-BE",
	DateOrder:             "",
}

var frKMLocale = LocaleData{
	Name:                  "fr-KM",
	DateOrder:             "",
}

var frCALocale = LocaleData{
	Name:                  "fr-CA",
	DateOrder:             "YMD",
	July:                  []string{"juill"},
}

var frNCLocale = LocaleData{
	Name:                  "fr-NC",
	DateOrder:             "",
}

var frDJLocale = LocaleData{
	Name:                  "fr-DJ",
	DateOrder:             "",
}

var frSYLocale = LocaleData{
	Name:                  "fr-SY",
	DateOrder:             "",
}

var frVULocale = LocaleData{
	Name:                  "fr-VU",
	DateOrder:             "",
}

var frGQLocale = LocaleData{
	Name:                  "fr-GQ",
	DateOrder:             "",
}

var frSCLocale = LocaleData{
	Name:                  "fr-SC",
	DateOrder:             "",
}

var frMQLocale = LocaleData{
	Name:                  "fr-MQ",
	DateOrder:             "",
}

var frMGLocale = LocaleData{
	Name:                  "fr-MG",
	DateOrder:             "",
}

var frBILocale = LocaleData{
	Name:                  "fr-BI",
	DateOrder:             "",
}

var frMULocale = LocaleData{
	Name:                  "fr-MU",
	DateOrder:             "",
}

var frCILocale = LocaleData{
	Name:                  "fr-CI",
	DateOrder:             "",
}

var frTDLocale = LocaleData{
	Name:                  "fr-TD",
	DateOrder:             "",
}

var frLULocale = LocaleData{
	Name:                  "fr-LU",
	DateOrder:             "",
}

var frNELocale = LocaleData{
	Name:                  "fr-NE",
	DateOrder:             "",
}

var frBLLocale = LocaleData{
	Name:                  "fr-BL",
	DateOrder:             "",
}

var frCHLocale = LocaleData{
	Name:                  "fr-CH",
	DateOrder:             "",
}

var frCFLocale = LocaleData{
	Name:                  "fr-CF",
	DateOrder:             "",
}

var frMALocale = LocaleData{
	Name:                  "fr-MA",
	DateOrder:             "",
	January:               []string{"jan"},
	February:              []string{"fév"},
	March:                 []string{"mar"},
	June:                  []string{"jui"},
}

var frWFLocale = LocaleData{
	Name:                  "fr-WF",
	DateOrder:             "",
}
