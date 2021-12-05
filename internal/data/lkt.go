package data

import "regexp"

var lktLocale = LocaleData{
	Name:                  "lkt",
	DateOrder:             "YMD",
	January:               []string{"wiótheȟika wí"},
	February:              []string{"thiyóȟeyuŋka wí"},
	March:                 []string{"ištáwičhayazaŋ wí"},
	April:                 []string{"pȟežítȟo wí"},
	May:                   []string{"čhaŋwápetȟo wí"},
	June:                  []string{"wípazukȟa-wašté wí"},
	July:                  []string{"čhaŋpȟásapa wí"},
	August:                []string{"wasútȟuŋ wí"},
	September:             []string{"čhaŋwápeǧi wí"},
	October:               []string{"čhaŋwápe-kasná wí"},
	November:              []string{"waníyetu wí"},
	December:              []string{"tȟahékapšuŋ wí"},
	Monday:                []string{"aŋpétuwaŋži"},
	Tuesday:               []string{"aŋpétunuŋpa"},
	Wednesday:             []string{"aŋpétuyamni"},
	Thursday:              []string{"aŋpétutopa"},
	Friday:                []string{"aŋpétuzaptaŋ"},
	Saturday:              []string{"owáŋgyužažapi"},
	Sunday:                []string{"aŋpétuwakȟaŋ"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ómakȟa"},
	Month:                 []string{"wí"},
	Week:                  []string{"okó"},
	Day:                   []string{"aŋpétu"},
	Hour:                  []string{"owápȟe"},
	Minute:                []string{"owápȟe oȟ'áŋkȟo"},
	Second:                []string{"okpí"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`lé aŋpétu kiŋ`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`lé wí kiŋ`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`lé okó kiŋ`},
		`0 year ago`:   {`lé ómakȟa kiŋ`},
		`1 day ago`:    {`ȟtálehaŋ`},
		`1 month ago`:  {`wí k'uŋ héhaŋ`},
		`1 week ago`:   {`okó k'uŋ héhaŋ`},
		`1 year ago`:   {`ómakȟa k'uŋ héhaŋ`},
		`in 1 day`:     {`híŋhaŋni kiŋháŋ`},
		`in 1 month`:   {`tȟokáta wí kiŋháŋ`},
		`in 1 week`:    {`tȟokáta okó kiŋháŋ`},
		`in 1 year`:    {`tȟokáta ómakȟa kiŋháŋ`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)hékta (\d+)-čháŋ k'uŋ héhaŋ`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)hékta owápȟe (\d+) k'uŋ héhaŋ`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)hékta oȟ'áŋkȟo (\d+) k'uŋ héhaŋ`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)hékta wíyawapi (\d+) k'uŋ héhaŋ`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)hékta okpí (\d+) k'uŋ héhaŋ`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)hékta okó (\d+) k'uŋ héhaŋ`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)hékta ómakȟa (\d+) k'uŋ héhaŋ`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)letáŋhaŋ (\d+)-čháŋ kiŋháŋ`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)letáŋhaŋ owápȟe (\d+) kiŋháŋ`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)letáŋhaŋ oȟ'áŋkȟo (\d+) kiŋháŋ`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)letáŋhaŋ wíyawapi (\d+) kiŋháŋ`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)letáŋhaŋ okpí (\d+) kiŋháŋ`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)letáŋhaŋ okó (\d+) kiŋháŋ`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)letáŋhaŋ ómakȟa (\d+) kiŋháŋ`),
		},
	},
}
