package data

import "regexp"

var taLocale = LocaleData{
	Name:                  "ta",
	DateOrder:             "DMY",
	January:               []string{"ஜன", "ஜனவரி"},
	February:              []string{"பிப்", "பிப்ரவரி"},
	March:                 []string{"மார்", "மார்ச்"},
	April:                 []string{"ஏப்", "ஏப்ரல்"},
	May:                   []string{"மே"},
	June:                  []string{"ஜூன்"},
	July:                  []string{"ஜூலை"},
	August:                []string{"ஆக", "ஆகஸ்ட்"},
	September:             []string{"செப்", "செப்டம்பர்"},
	October:               []string{"அக்", "அக்டோபர்"},
	November:              []string{"நவ", "நவம்பர்"},
	December:              []string{"டிச", "டிசம்பர்"},
	Monday:                []string{"திங்", "திங்கள்"},
	Tuesday:               []string{"செவ்", "செவ்வாய்"},
	Wednesday:             []string{"புத", "புதன்"},
	Thursday:              []string{"வியா", "வியாழன்"},
	Friday:                []string{"வெள்", "வெள்ளி"},
	Saturday:              []string{"சனி"},
	Sunday:                []string{"ஞாயி", "ஞாயிறு"},
	AM:                    []string{"முற்பகல்"},
	PM:                    []string{"பிற்பகல்"},
	Year:                  []string{"ஆ", "ஆண்டு"},
	Month:                 []string{"மா", "மாத", "மாதம்"},
	Week:                  []string{"வா", "வாரம்"},
	Day:                   []string{"நா", "நாள்"},
	Hour:                  []string{"ம", "மணி"},
	Minute:                []string{"நிமி", "நிமிடம்"},
	Second:                []string{"வி", "விநா", "விநாடி"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`இன்று`},
		`0 hour ago`:   {`இந்த ஒரு மணிநேரத்தில்`},
		`0 minute ago`: {`இந்த ஒரு நிமிடத்தில்`},
		`0 month ago`:  {`இந்த மாதம்`},
		`0 second ago`: {`இப்போது`},
		`0 week ago`:   {`இந்த வாரம்`},
		`0 year ago`:   {`இந்த ஆண்டு`},
		`1 day ago`:    {`நேற்று`},
		`1 month ago`:  {`கடந்த மாதம்`},
		`1 week ago`:   {`கடந்த வாரம்`},
		`1 year ago`:   {`கடந்த ஆண்டு`},
		`in 1 day`:     {`நாளை`},
		`in 1 month`:   {`அடுத்த மாதம்`},
		`in 1 week`:    {`அடுத்த வாரம்`},
		`in 1 year`:    {`அடுத்த ஆண்டு`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) நா முன்`),
			regexp.MustCompile(`(?i)(\d+) நாட்களுக்கு முன்`),
			regexp.MustCompile(`(?i)(\d+) நாளுக்கு முன்`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ம முன்`),
			regexp.MustCompile(`(?i)(\d+) மணி முன்`),
			regexp.MustCompile(`(?i)(\d+) மணிநேரம் முன்`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) நி முன்`),
			regexp.MustCompile(`(?i)(\d+) நிமி முன்`),
			regexp.MustCompile(`(?i)(\d+) நிமிடங்களுக்கு முன்`),
			regexp.MustCompile(`(?i)(\d+) நிமிடத்திற்கு முன்`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) மா முன்`),
			regexp.MustCompile(`(?i)(\d+) மாத முன்`),
			regexp.MustCompile(`(?i)(\d+) மாதங்களுக்கு முன்`),
			regexp.MustCompile(`(?i)(\d+) மாதத்துக்கு முன்`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) வி முன்`),
			regexp.MustCompile(`(?i)(\d+) விநா முன்`),
			regexp.MustCompile(`(?i)(\d+) விநாடிகளுக்கு முன்`),
			regexp.MustCompile(`(?i)(\d+) விநாடிக்கு முன்`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) வா முன்`),
			regexp.MustCompile(`(?i)(\d+) வார முன்`),
			regexp.MustCompile(`(?i)(\d+) வாரங்களுக்கு முன்`),
			regexp.MustCompile(`(?i)(\d+) வாரத்திற்கு முன்பு`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ஆ முன்`),
			regexp.MustCompile(`(?i)(\d+) ஆண்டிற்கு முன்`),
			regexp.MustCompile(`(?i)(\d+) ஆண்டுகளுக்கு முன்`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) நா`),
			regexp.MustCompile(`(?i)(\d+) நாட்களில்`),
			regexp.MustCompile(`(?i)(\d+) நாளில்`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ம`),
			regexp.MustCompile(`(?i)(\d+) மணி`),
			regexp.MustCompile(`(?i)(\d+) மணிநேரத்தில்`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) நி`),
			regexp.MustCompile(`(?i)(\d+) நிமி`),
			regexp.MustCompile(`(?i)(\d+) நிமிடங்களில்`),
			regexp.MustCompile(`(?i)(\d+) நிமிடத்தில்`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) மா`),
			regexp.MustCompile(`(?i)(\d+) மாத`),
			regexp.MustCompile(`(?i)(\d+) மாதங்களில்`),
			regexp.MustCompile(`(?i)(\d+) மாதத்தில்`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) வி`),
			regexp.MustCompile(`(?i)(\d+) விநா`),
			regexp.MustCompile(`(?i)(\d+) விநாடிகளில்`),
			regexp.MustCompile(`(?i)(\d+) விநாடியில்`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) வா`),
			regexp.MustCompile(`(?i)(\d+) வார`),
			regexp.MustCompile(`(?i)(\d+) வாரங்களில்`),
			regexp.MustCompile(`(?i)(\d+) வாரத்தில்`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) ஆ`),
			regexp.MustCompile(`(?i)(\d+) ஆண்டில்`),
			regexp.MustCompile(`(?i)(\d+) ஆண்டுகளில்`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ta-LK": taLKLocale,
		"ta-MY": taMYLocale,
		"ta-SG": taSGLocale,
	},
}

var taLKLocale = LocaleData{
	Name:                  "ta-LK",
	DateOrder:             "",
}

var taMYLocale = LocaleData{
	Name:                  "ta-MY",
	DateOrder:             "",
}

var taSGLocale = LocaleData{
	Name:                  "ta-SG",
	DateOrder:             "",
}
