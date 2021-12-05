package data

import "regexp"

var enLocale = LocaleData{
	Name:                  "en",
	DateOrder:             "MDY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "about", "ad", "and", "at", "by", "just", "m", "nd", "of", "on", "rd", "st", "th", "the", "|", "，"},
	PertainWords:          []string{"of"},
	January:               []string{"jan", "january"},
	February:              []string{"feb", "february"},
	March:                 []string{"mar", "march"},
	April:                 []string{"apr", "april"},
	May:                   []string{"may"},
	June:                  []string{"jun", "june"},
	July:                  []string{"jul", "july"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "sept", "september"},
	October:               []string{"oct", "october"},
	November:              []string{"nov", "november"},
	December:              []string{"dec", "december"},
	Monday:                []string{"mon", "monday"},
	Tuesday:               []string{"tue", "tues", "tuesday"},
	Wednesday:             []string{"wed", "wednesday"},
	Thursday:              []string{"thu", "thursday"},
	Friday:                []string{"fri", "friday"},
	Saturday:              []string{"sat", "saturday"},
	Sunday:                []string{"sun", "sunday"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Decade:                []string{"decade", "decades"},
	Year:                  []string{"year", "years", "yr"},
	Month:                 []string{"mo", "month", "months"},
	Week:                  []string{"week", "weeks", "wk"},
	Day:                   []string{"day", "days"},
	Hour:                  []string{"h", "hour", "hours", "hr", "hrs"},
	Minute:                []string{"m", "min", "mins", "minute", "minutes"},
	Second:                []string{"s", "sec", "second", "seconds", "secs"},
	In:                    []string{"from now", "in"},
	Ago:                   []string{"ago"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`today`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this mo`, `this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`, `this wk`},
		`0 year ago`:   {`this year`, `this yr`},
		`1 day ago`:    {`yesterday`},
		`1 decade ago`: {`last decade`, `this decade`},
		`1 month ago`:  {`last mo`, `last month`},
		`1 week ago`:   {`last week`, `last wk`},
		`1 year ago`:   {`last year`, `last yr`},
		`2 day ago`:    {`day before yesterday`},
		`in 1 day`:     {`tomorrow`},
		`in 1 decade`:  {`next decade`},
		`in 1 month`:   {`next mo`, `next month`},
		`in 1 week`:    {`next week`, `next wk`},
		`in 1 year`:    {`next year`, `next yr`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) day ago`),
			regexp.MustCompile(`(?i)(\d+) days ago`),
		},
		`\1 decade ago`: {
			regexp.MustCompile(`(?i)(\d+) decades? ago`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) hour ago`),
			regexp.MustCompile(`(?i)(\d+) hours ago`),
			regexp.MustCompile(`(?i)(\d+) hr ago`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min ago`),
			regexp.MustCompile(`(?i)(\d+) minute ago`),
			regexp.MustCompile(`(?i)(\d+) minutes ago`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mo ago`),
			regexp.MustCompile(`(?i)(\d+) month ago`),
			regexp.MustCompile(`(?i)(\d+) months ago`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) sec ago`),
			regexp.MustCompile(`(?i)(\d+) second ago`),
			regexp.MustCompile(`(?i)(\d+) seconds ago`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) week ago`),
			regexp.MustCompile(`(?i)(\d+) weeks ago`),
			regexp.MustCompile(`(?i)(\d+) wk ago`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) year ago`),
			regexp.MustCompile(`(?i)(\d+) years ago`),
			regexp.MustCompile(`(?i)(\d+) yr ago`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)in (\d+) day`),
			regexp.MustCompile(`(?i)in (\d+) days`),
		},
		`in \1 decade`: {
			regexp.MustCompile(`(?i)in (\d+) decades?`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)in (\d+) hour`),
			regexp.MustCompile(`(?i)in (\d+) hours`),
			regexp.MustCompile(`(?i)in (\d+) hr`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)in (\d+) min`),
			regexp.MustCompile(`(?i)in (\d+) minute`),
			regexp.MustCompile(`(?i)in (\d+) minutes`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)in (\d+) mo`),
			regexp.MustCompile(`(?i)in (\d+) month`),
			regexp.MustCompile(`(?i)in (\d+) months`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)in (\d+) sec`),
			regexp.MustCompile(`(?i)in (\d+) second`),
			regexp.MustCompile(`(?i)in (\d+) seconds`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)in (\d+) week`),
			regexp.MustCompile(`(?i)in (\d+) weeks`),
			regexp.MustCompile(`(?i)in (\d+) wk`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)in (\d+) year`),
			regexp.MustCompile(`(?i)in (\d+) years`),
			regexp.MustCompile(`(?i)in (\d+) yr`),
		},
	},
	Simplifications: map[string]string{
		`(?:12\s+)?midnight`:     `00:00`,
		`(?:12\s+)?noon`:         `12:00`,
		`(?<=from\s+)now`:        `in`,
		`(\d+)h(\d+)m?`:          `\1:\2`,
		`a`:                      `1`,
		`an`:                     `1`,
		`eight`:                  `8`,
		`eleven`:                 `11`,
		`five`:                   `5`,
		`four`:                   `4`,
		`less than 1 minute ago`: `45 second ago`,
		`nine`:                   `9`,
		`one`:                    `1`,
		`seven`:                  `7`,
		`six`:                    `6`,
		`ten`:                    `10`,
		`three`:                  `3`,
		`twelve`:                 `12`,
		`two`:                    `2`,
	},
	LocaleSpecific: map[string]LocaleData{
		"en-001": en001Locale,
		"en-150": en150Locale,
		"en-AG":  enAGLocale,
		"en-AI":  enAILocale,
		"en-AS":  enASLocale,
		"en-AT":  enATLocale,
		"en-AU":  enAULocale,
		"en-BB":  enBBLocale,
		"en-BE":  enBELocale,
		"en-BI":  enBILocale,
		"en-BM":  enBMLocale,
		"en-BS":  enBSLocale,
		"en-BW":  enBWLocale,
		"en-BZ":  enBZLocale,
		"en-CA":  enCALocale,
		"en-CC":  enCCLocale,
		"en-CH":  enCHLocale,
		"en-CK":  enCKLocale,
		"en-CM":  enCMLocale,
		"en-CX":  enCXLocale,
		"en-CY":  enCYLocale,
		"en-DE":  enDELocale,
		"en-DG":  enDGLocale,
		"en-DK":  enDKLocale,
		"en-DM":  enDMLocale,
		"en-ER":  enERLocale,
		"en-FI":  enFILocale,
		"en-FJ":  enFJLocale,
		"en-FK":  enFKLocale,
		"en-FM":  enFMLocale,
		"en-GB":  enGBLocale,
		"en-GD":  enGDLocale,
		"en-GG":  enGGLocale,
		"en-GH":  enGHLocale,
		"en-GI":  enGILocale,
		"en-GM":  enGMLocale,
		"en-GU":  enGULocale,
		"en-GY":  enGYLocale,
		"en-HK":  enHKLocale,
		"en-IE":  enIELocale,
		"en-IL":  enILLocale,
		"en-IM":  enIMLocale,
		"en-IN":  enINLocale,
		"en-IO":  enIOLocale,
		"en-JE":  enJELocale,
		"en-JM":  enJMLocale,
		"en-KE":  enKELocale,
		"en-KI":  enKILocale,
		"en-KN":  enKNLocale,
		"en-KY":  enKYLocale,
		"en-LC":  enLCLocale,
		"en-LR":  enLRLocale,
		"en-LS":  enLSLocale,
		"en-MG":  enMGLocale,
		"en-MH":  enMHLocale,
		"en-MO":  enMOLocale,
		"en-MP":  enMPLocale,
		"en-MS":  enMSLocale,
		"en-MT":  enMTLocale,
		"en-MU":  enMULocale,
		"en-MW":  enMWLocale,
		"en-MY":  enMYLocale,
		"en-NA":  enNALocale,
		"en-NF":  enNFLocale,
		"en-NG":  enNGLocale,
		"en-NL":  enNLLocale,
		"en-NR":  enNRLocale,
		"en-NU":  enNULocale,
		"en-NZ":  enNZLocale,
		"en-PG":  enPGLocale,
		"en-PH":  enPHLocale,
		"en-PK":  enPKLocale,
		"en-PN":  enPNLocale,
		"en-PR":  enPRLocale,
		"en-PW":  enPWLocale,
		"en-RW":  enRWLocale,
		"en-SB":  enSBLocale,
		"en-SC":  enSCLocale,
		"en-SD":  enSDLocale,
		"en-SE":  enSELocale,
		"en-SG":  enSGLocale,
		"en-SH":  enSHLocale,
		"en-SI":  enSILocale,
		"en-SL":  enSLLocale,
		"en-SS":  enSSLocale,
		"en-SX":  enSXLocale,
		"en-SZ":  enSZLocale,
		"en-TC":  enTCLocale,
		"en-TK":  enTKLocale,
		"en-TO":  enTOLocale,
		"en-TT":  enTTLocale,
		"en-TV":  enTVLocale,
		"en-TZ":  enTZLocale,
		"en-UG":  enUGLocale,
		"en-UM":  enUMLocale,
		"en-VC":  enVCLocale,
		"en-VG":  enVGLocale,
		"en-VI":  enVILocale,
		"en-VU":  enVULocale,
		"en-WS":  enWSLocale,
		"en-ZA":  enZALocale,
		"en-ZM":  enZMLocale,
		"en-ZW":  enZWLocale,
	},
}

var enPRLocale = LocaleData{
	Name:                  "en-PR",
	DateOrder:             "",
}

var en001Locale = LocaleData{
	Name:                  "en-001",
	DateOrder:             "DMY",
}

var enVGLocale = LocaleData{
	Name:                  "en-VG",
	DateOrder:             "DMY",
}

var enDGLocale = LocaleData{
	Name:                  "en-DG",
	DateOrder:             "DMY",
}

var en150Locale = LocaleData{
	Name:                  "en-150",
	DateOrder:             "DMY",
}

var enSZLocale = LocaleData{
	Name:                  "en-SZ",
	DateOrder:             "DMY",
}

var enIELocale = LocaleData{
	Name:                  "en-IE",
	DateOrder:             "DMY",
}

var enMSLocale = LocaleData{
	Name:                  "en-MS",
	DateOrder:             "DMY",
}

var enMOLocale = LocaleData{
	Name:                  "en-MO",
	DateOrder:             "DMY",
}

var enAILocale = LocaleData{
	Name:                  "en-AI",
	DateOrder:             "DMY",
}

var enTKLocale = LocaleData{
	Name:                  "en-TK",
	DateOrder:             "DMY",
}

var enMPLocale = LocaleData{
	Name:                  "en-MP",
	DateOrder:             "",
}

var enUGLocale = LocaleData{
	Name:                  "en-UG",
	DateOrder:             "DMY",
}

var enGBLocale = LocaleData{
	Name:                  "en-GB",
	DateOrder:             "DMY",
}

var enPWLocale = LocaleData{
	Name:                  "en-PW",
	DateOrder:             "DMY",
}

var enCYLocale = LocaleData{
	Name:                  "en-CY",
	DateOrder:             "DMY",
}

var enFKLocale = LocaleData{
	Name:                  "en-FK",
	DateOrder:             "DMY",
}

var enWSLocale = LocaleData{
	Name:                  "en-WS",
	DateOrder:             "DMY",
}

var enMGLocale = LocaleData{
	Name:                  "en-MG",
	DateOrder:             "DMY",
}

var enVILocale = LocaleData{
	Name:                  "en-VI",
	DateOrder:             "",
}

var enGGLocale = LocaleData{
	Name:                  "en-GG",
	DateOrder:             "DMY",
}

var enCMLocale = LocaleData{
	Name:                  "en-CM",
	DateOrder:             "DMY",
}

var enSELocale = LocaleData{
	Name:                  "en-SE",
	DateOrder:             "YMD",
}

var enPNLocale = LocaleData{
	Name:                  "en-PN",
	DateOrder:             "DMY",
}

var enERLocale = LocaleData{
	Name:                  "en-ER",
	DateOrder:             "DMY",
}

var enGULocale = LocaleData{
	Name:                  "en-GU",
	DateOrder:             "",
}

var enMWLocale = LocaleData{
	Name:                  "en-MW",
	DateOrder:             "DMY",
}

var enNFLocale = LocaleData{
	Name:                  "en-NF",
	DateOrder:             "DMY",
}

var enNLLocale = LocaleData{
	Name:                  "en-NL",
	DateOrder:             "DMY",
}

var enTVLocale = LocaleData{
	Name:                  "en-TV",
	DateOrder:             "DMY",
}

var enSBLocale = LocaleData{
	Name:                  "en-SB",
	DateOrder:             "DMY",
}

var enLSLocale = LocaleData{
	Name:                  "en-LS",
	DateOrder:             "DMY",
}

var enCHLocale = LocaleData{
	Name:                  "en-CH",
	DateOrder:             "DMY",
}

var enSDLocale = LocaleData{
	Name:                  "en-SD",
	DateOrder:             "DMY",
}

var enCCLocale = LocaleData{
	Name:                  "en-CC",
	DateOrder:             "DMY",
}

var enJELocale = LocaleData{
	Name:                  "en-JE",
	DateOrder:             "DMY",
}

var enBZLocale = LocaleData{
	Name:                  "en-BZ",
	DateOrder:             "DMY",
}

var enKELocale = LocaleData{
	Name:                  "en-KE",
	DateOrder:             "DMY",
}

var enPHLocale = LocaleData{
	Name:                  "en-PH",
	DateOrder:             "DMY",
}

var enDELocale = LocaleData{
	Name:                  "en-DE",
	DateOrder:             "DMY",
}

var enFMLocale = LocaleData{
	Name:                  "en-FM",
	DateOrder:             "DMY",
}

var enMHLocale = LocaleData{
	Name:                  "en-MH",
	DateOrder:             "",
}

var enMYLocale = LocaleData{
	Name:                  "en-MY",
	DateOrder:             "DMY",
}

var enTCLocale = LocaleData{
	Name:                  "en-TC",
	DateOrder:             "DMY",
}

var enSHLocale = LocaleData{
	Name:                  "en-SH",
	DateOrder:             "DMY",
}

var enCXLocale = LocaleData{
	Name:                  "en-CX",
	DateOrder:             "DMY",
}

var enAGLocale = LocaleData{
	Name:                  "en-AG",
	DateOrder:             "DMY",
}

var enGILocale = LocaleData{
	Name:                  "en-GI",
	DateOrder:             "DMY",
}

var enIOLocale = LocaleData{
	Name:                  "en-IO",
	DateOrder:             "DMY",
}

var enVCLocale = LocaleData{
	Name:                  "en-VC",
	DateOrder:             "DMY",
}

var enBBLocale = LocaleData{
	Name:                  "en-BB",
	DateOrder:             "DMY",
}

var enCALocale = LocaleData{
	Name:                  "en-CA",
	DateOrder:             "YMD",
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) hrs ago`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) mins ago`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mos ago`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) secs ago`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) wks ago`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) yrs ago`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)in (\d+) hrs`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)in (\d+) mins`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)in (\d+) mos`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)in (\d+) secs`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)in (\d+) wks`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)in (\d+) yrs`),
		},
	},
}

var enKYLocale = LocaleData{
	Name:                  "en-KY",
	DateOrder:             "DMY",
}

var enPKLocale = LocaleData{
	Name:                  "en-PK",
	DateOrder:             "DMY",
}

var enNULocale = LocaleData{
	Name:                  "en-NU",
	DateOrder:             "DMY",
}

var enATLocale = LocaleData{
	Name:                  "en-AT",
	DateOrder:             "DMY",
}

var enPGLocale = LocaleData{
	Name:                  "en-PG",
	DateOrder:             "DMY",
}

var enVULocale = LocaleData{
	Name:                  "en-VU",
	DateOrder:             "DMY",
}

var enILLocale = LocaleData{
	Name:                  "en-IL",
	DateOrder:             "DMY",
}

var enTZLocale = LocaleData{
	Name:                  "en-TZ",
	DateOrder:             "DMY",
}

var enIMLocale = LocaleData{
	Name:                  "en-IM",
	DateOrder:             "DMY",
}

var enFILocale = LocaleData{
	Name:                  "en-FI",
	DateOrder:             "DMY",
}

var enZWLocale = LocaleData{
	Name:                  "en-ZW",
	DateOrder:             "DMY",
}

var enGHLocale = LocaleData{
	Name:                  "en-GH",
	DateOrder:             "DMY",
}

var enNZLocale = LocaleData{
	Name:                  "en-NZ",
	DateOrder:             "DMY",
}

var enGDLocale = LocaleData{
	Name:                  "en-GD",
	DateOrder:             "DMY",
}

var enBMLocale = LocaleData{
	Name:                  "en-BM",
	DateOrder:             "DMY",
}

var enBILocale = LocaleData{
	Name:                  "en-BI",
	DateOrder:             "",
}

var enHKLocale = LocaleData{
	Name:                  "en-HK",
	DateOrder:             "DMY",
}

var enDMLocale = LocaleData{
	Name:                  "en-DM",
	DateOrder:             "DMY",
}

var enSGLocale = LocaleData{
	Name:                  "en-SG",
	DateOrder:             "DMY",
	Month:                 []string{"mth"},
	RelativeType: map[string][]string{
		`0 month ago`: {`this mth`},
		`1 month ago`: {`last mth`},
		`in 1 month`:  {`next mth`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mth ago`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)in (\d+) mth`),
		},
	},
}

var enBELocale = LocaleData{
	Name:                  "en-BE",
	DateOrder:             "DMY",
}

var enKILocale = LocaleData{
	Name:                  "en-KI",
	DateOrder:             "DMY",
}

var enKNLocale = LocaleData{
	Name:                  "en-KN",
	DateOrder:             "DMY",
}

var enINLocale = LocaleData{
	Name:                  "en-IN",
	DateOrder:             "DMY",
}

var enSXLocale = LocaleData{
	Name:                  "en-SX",
	DateOrder:             "DMY",
}

var enCKLocale = LocaleData{
	Name:                  "en-CK",
	DateOrder:             "DMY",
}

var enNGLocale = LocaleData{
	Name:                  "en-NG",
	DateOrder:             "DMY",
}

var enSCLocale = LocaleData{
	Name:                  "en-SC",
	DateOrder:             "DMY",
}

var enUMLocale = LocaleData{
	Name:                  "en-UM",
	DateOrder:             "",
}

var enTOLocale = LocaleData{
	Name:                  "en-TO",
	DateOrder:             "DMY",
}

var enSILocale = LocaleData{
	Name:                  "en-SI",
	DateOrder:             "DMY",
}

var enFJLocale = LocaleData{
	Name:                  "en-FJ",
	DateOrder:             "DMY",
}

var enDKLocale = LocaleData{
	Name:                  "en-DK",
	DateOrder:             "DMY",
}

var enGYLocale = LocaleData{
	Name:                  "en-GY",
	DateOrder:             "DMY",
}

var enSLLocale = LocaleData{
	Name:                  "en-SL",
	DateOrder:             "DMY",
}

var enZMLocale = LocaleData{
	Name:                  "en-ZM",
	DateOrder:             "DMY",
}

var enASLocale = LocaleData{
	Name:                  "en-AS",
	DateOrder:             "",
}

var enAULocale = LocaleData{
	Name:                  "en-AU",
	DateOrder:             "DMY",
	Hour:                  []string{"h"},
}

var enNRLocale = LocaleData{
	Name:                  "en-NR",
	DateOrder:             "DMY",
}

var enTTLocale = LocaleData{
	Name:                  "en-TT",
	DateOrder:             "DMY",
}

var enBWLocale = LocaleData{
	Name:                  "en-BW",
	DateOrder:             "DMY",
}

var enGMLocale = LocaleData{
	Name:                  "en-GM",
	DateOrder:             "DMY",
}

var enJMLocale = LocaleData{
	Name:                  "en-JM",
	DateOrder:             "DMY",
}

var enSSLocale = LocaleData{
	Name:                  "en-SS",
	DateOrder:             "DMY",
}

var enMULocale = LocaleData{
	Name:                  "en-MU",
	DateOrder:             "DMY",
}

var enLCLocale = LocaleData{
	Name:                  "en-LC",
	DateOrder:             "DMY",
}

var enNALocale = LocaleData{
	Name:                  "en-NA",
	DateOrder:             "DMY",
}

var enZALocale = LocaleData{
	Name:                  "en-ZA",
	DateOrder:             "YMD",
}

var enLRLocale = LocaleData{
	Name:                  "en-LR",
	DateOrder:             "DMY",
}

var enRWLocale = LocaleData{
	Name:                  "en-RW",
	DateOrder:             "DMY",
}

var enMTLocale = LocaleData{
	Name:                  "en-MT",
	DateOrder:             "DMY",
}

var enBSLocale = LocaleData{
	Name:                  "en-BS",
	DateOrder:             "DMY",
}
