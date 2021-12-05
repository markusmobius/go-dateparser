package data

import "regexp"

var arLocale = LocaleData{
	Name:                  "ar",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 6,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "الساعة", "في", "مساءً", "هـ", "，"},
	January:               []string{"يناير"},
	February:              []string{"فبراير"},
	March:                 []string{"مارس"},
	April:                 []string{"أبريل"},
	May:                   []string{"مايو"},
	June:                  []string{"يونيو"},
	July:                  []string{"يوليو"},
	August:                []string{"أغسطس"},
	September:             []string{"سبتمبر"},
	October:               []string{"أكتوبر"},
	November:              []string{"نوفمبر"},
	December:              []string{"ديسمبر"},
	Monday:                []string{"الإثنين", "الاثنين"},
	Tuesday:               []string{"الثلاثاء"},
	Wednesday:             []string{"الأربعاء"},
	Thursday:              []string{"الخميس"},
	Friday:                []string{"الجمعة"},
	Saturday:              []string{"السبت"},
	Sunday:                []string{"الأحد"},
	AM:                    []string{"ص", "صباحاً", "صباحًا"},
	PM:                    []string{"م", "مساءً"},
	Year:                  []string{"السنة", "سنة", "عام"},
	Month:                 []string{"الشهر", "شهر"},
	Week:                  []string{"أسبوع", "الأسبوع"},
	Day:                   []string{"أيام", "يوم"},
	Hour:                  []string{"الساعات", "ساعات", "ساعة"},
	Minute:                []string{"الدقائق", "دقائق", "دقيقة"},
	Second:                []string{"الثواني", "ثانية"},
	In:                    []string{"خلال"},
	Ago:                   []string{"منذ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`اليوم`},
		`0 hour ago`:   {`الساعة الحالية`},
		`0 minute ago`: {`هذه الدقيقة`},
		`0 month ago`:  {`هذا الشهر`},
		`0 second ago`: {`الآن`},
		`0 week ago`:   {`هذا الأسبوع`},
		`0 year ago`:   {`السنة الحالية`},
		`1 day ago`:    {`أمس`, `الأمس`, `اليوم السابق`},
		`1 hour ago`:   {`ساعة واحدة`},
		`1 month ago`:  {`الشهر الماضي`},
		`1 week ago`:   {`الأسبوع الماضي`},
		`1 year ago`:   {`السنة الماضية`},
		`2 day`:        {`يومين`},
		`2 hour`:       {`ساعتين`},
		`in 1 day`:     {`غدًا`},
		`in 1 month`:   {`الشهر القادم`},
		`in 1 week`:    {`الأسبوع القادم`},
		`in 1 year`:    {`السنة القادمة`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) يوم`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) ساعة`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) دقيقة`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) شهر`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) ثانية`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) أسبوع`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)قبل (\d+) سنة`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)خلال (\d+) يوم`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)خلال (\d+) ساعة`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)خلال (\d+) دقيقة`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)خلال (\d+) شهر`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)خلال (\d+) ثانية`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)خلال (\d+) أسبوع`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)خلال (\d+) سنة`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ar-AE": arAELocale,
		"ar-BH": arBHLocale,
		"ar-DJ": arDJLocale,
		"ar-DZ": arDZLocale,
		"ar-EG": arEGLocale,
		"ar-EH": arEHLocale,
		"ar-ER": arERLocale,
		"ar-IL": arILLocale,
		"ar-IQ": arIQLocale,
		"ar-JO": arJOLocale,
		"ar-KM": arKMLocale,
		"ar-KW": arKWLocale,
		"ar-LB": arLBLocale,
		"ar-LY": arLYLocale,
		"ar-MA": arMALocale,
		"ar-MR": arMRLocale,
		"ar-OM": arOMLocale,
		"ar-PS": arPSLocale,
		"ar-QA": arQALocale,
		"ar-SA": arSALocale,
		"ar-SD": arSDLocale,
		"ar-SO": arSOLocale,
		"ar-SS": arSSLocale,
		"ar-SY": arSYLocale,
		"ar-TD": arTDLocale,
		"ar-TN": arTNLocale,
		"ar-YE": arYELocale,
	},
}

var arAELocale = LocaleData{
	Name:                  "ar-AE",
	DateOrder:             "",
	RelativeType: map[string][]string{
		`0 year ago`: {`هذه السنة`},
		`in 1 year`:  {`السنة التالية`},
	},
}

var arSDLocale = LocaleData{
	Name:                  "ar-SD",
	DateOrder:             "",
}

var arIQLocale = LocaleData{
	Name:                  "ar-IQ",
	DateOrder:             "",
	January:               []string{"كانون الثاني"},
	February:              []string{"شباط"},
	March:                 []string{"آذار"},
	April:                 []string{"نيسان"},
	May:                   []string{"أيار"},
	June:                  []string{"حزيران"},
	July:                  []string{"تموز"},
	August:                []string{"آب"},
	September:             []string{"أيلول"},
	October:               []string{"تشرين الأول", "تشرین الأول"},
	November:              []string{"تشرين الثاني"},
	December:              []string{"كانون الأول"},
}

var arJOLocale = LocaleData{
	Name:                  "ar-JO",
	DateOrder:             "",
	January:               []string{"كانون الثاني"},
	February:              []string{"شباط"},
	March:                 []string{"آذار"},
	April:                 []string{"نيسان"},
	May:                   []string{"أيار"},
	June:                  []string{"حزيران"},
	July:                  []string{"تموز"},
	August:                []string{"آب"},
	September:             []string{"أيلول"},
	October:               []string{"تشرين الأول"},
	November:              []string{"تشرين الثاني"},
	December:              []string{"كانون الأول"},
}

var arLYLocale = LocaleData{
	Name:                  "ar-LY",
	DateOrder:             "",
}

var arQALocale = LocaleData{
	Name:                  "ar-QA",
	DateOrder:             "",
}

var arSSLocale = LocaleData{
	Name:                  "ar-SS",
	DateOrder:             "",
}

var arPSLocale = LocaleData{
	Name:                  "ar-PS",
	DateOrder:             "",
	January:               []string{"كانون الثاني"},
	February:              []string{"شباط"},
	March:                 []string{"آذار"},
	April:                 []string{"نيسان"},
	May:                   []string{"أيار"},
	June:                  []string{"حزيران"},
	July:                  []string{"تموز"},
	August:                []string{"آب"},
	September:             []string{"أيلول"},
	October:               []string{"تشرين الأول"},
	November:              []string{"تشرين الثاني"},
	December:              []string{"كانون الأول"},
}

var arSALocale = LocaleData{
	Name:                  "ar-SA",
	DateOrder:             "",
}

var arSOLocale = LocaleData{
	Name:                  "ar-SO",
	DateOrder:             "",
}

var arDZLocale = LocaleData{
	Name:                  "ar-DZ",
	DateOrder:             "",
	January:               []string{"جانفي"},
	February:              []string{"فيفري"},
	April:                 []string{"أفريل"},
	May:                   []string{"ماي"},
	June:                  []string{"جوان"},
	July:                  []string{"جويلية"},
	August:                []string{"أوت"},
}

var arEGLocale = LocaleData{
	Name:                  "ar-EG",
	DateOrder:             "",
}

var arILLocale = LocaleData{
	Name:                  "ar-IL",
	DateOrder:             "",
}

var arEHLocale = LocaleData{
	Name:                  "ar-EH",
	DateOrder:             "",
}

var arLBLocale = LocaleData{
	Name:                  "ar-LB",
	DateOrder:             "",
	January:               []string{"كانون الثاني"},
	February:              []string{"شباط"},
	March:                 []string{"آذار"},
	April:                 []string{"نيسان"},
	May:                   []string{"أيار"},
	June:                  []string{"حزيران"},
	July:                  []string{"تموز"},
	August:                []string{"آب"},
	September:             []string{"أيلول"},
	October:               []string{"تشرين الأول"},
	November:              []string{"تشرين الثاني"},
	December:              []string{"كانون الأول"},
}

var arKMLocale = LocaleData{
	Name:                  "ar-KM",
	DateOrder:             "",
}

var arDJLocale = LocaleData{
	Name:                  "ar-DJ",
	DateOrder:             "",
}

var arTDLocale = LocaleData{
	Name:                  "ar-TD",
	DateOrder:             "",
}

var arKWLocale = LocaleData{
	Name:                  "ar-KW",
	DateOrder:             "",
}

var arERLocale = LocaleData{
	Name:                  "ar-ER",
	DateOrder:             "",
}

var arYELocale = LocaleData{
	Name:                  "ar-YE",
	DateOrder:             "",
}

var arOMLocale = LocaleData{
	Name:                  "ar-OM",
	DateOrder:             "",
}

var arMRLocale = LocaleData{
	Name:                  "ar-MR",
	DateOrder:             "",
	April:                 []string{"إبريل"},
	August:                []string{"أغشت"},
	September:             []string{"شتمبر"},
	December:              []string{"دجمبر"},
}

var arSYLocale = LocaleData{
	Name:                  "ar-SY",
	DateOrder:             "",
	January:               []string{"كانون الثاني"},
	February:              []string{"شباط"},
	March:                 []string{"آذار"},
	April:                 []string{"نيسان"},
	May:                   []string{"أيار"},
	June:                  []string{"حزيران"},
	July:                  []string{"تموز"},
	August:                []string{"آب"},
	September:             []string{"أيلول"},
	October:               []string{"تشرين الأول"},
	November:              []string{"تشرين الثاني"},
	December:              []string{"كانون الأول"},
}

var arTNLocale = LocaleData{
	Name:                  "ar-TN",
	DateOrder:             "",
	January:               []string{"جانفي"},
	February:              []string{"فيفري"},
	April:                 []string{"أفريل"},
	May:                   []string{"ماي"},
	June:                  []string{"جوان"},
	July:                  []string{"جويلية"},
	August:                []string{"أوت"},
}

var arMALocale = LocaleData{
	Name:                  "ar-MA",
	DateOrder:             "",
	May:                   []string{"ماي"},
	July:                  []string{"يوليوز"},
	August:                []string{"غشت"},
	September:             []string{"شتنبر"},
	November:              []string{"نونبر"},
	December:              []string{"دجنبر"},
}

var arBHLocale = LocaleData{
	Name:                  "ar-BH",
	DateOrder:             "",
}
