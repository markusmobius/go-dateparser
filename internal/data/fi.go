package data

import "regexp"

var fiLocale = LocaleData{
	Name:                  "fi",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ":n", ";", "@", "[", "]", "|", "，"},
	January:               []string{"tammi", "tammik", "tammikuu", "tammikuussa", "tammikuuta"},
	February:              []string{"helmi", "helmik", "helmikuu", "helmikuussa", "helmikuuta"},
	March:                 []string{"maalis", "maalisk", "maaliskuu", "maaliskuussa", "maaliskuuta"},
	April:                 []string{"huhti", "huhtik", "huhtikuu", "huhtikuussa", "huhtikuuta"},
	May:                   []string{"touko", "toukok", "toukokuu", "toukokuussa", "toukokuuta"},
	June:                  []string{"kesä", "kesäk", "kesäkuu", "kesäkuussa", "kesäkuuta"},
	July:                  []string{"heinä", "heinäk", "heinäkuu", "heinäkuussa", "heinäkuuta"},
	August:                []string{"elo", "elok", "elokuu", "elokuussa", "elokuuta"},
	September:             []string{"syys", "syysk", "syyskuu", "syyskuussa", "syyskuuta"},
	October:               []string{"loka", "lokak", "lokakuu", "lokakuussa", "lokakuuta"},
	November:              []string{"marras", "marrask", "marraskuu", "marraskuussa", "marraskuuta"},
	December:              []string{"joulu", "jouluk", "joulukuu", "joulukuussa", "joulukuuta"},
	Monday:                []string{"ma", "maanantai", "maanantaina"},
	Tuesday:               []string{"ti", "tiistai", "tiistaina"},
	Wednesday:             []string{"ke", "keskiviikko", "keskiviikkona"},
	Thursday:              []string{"to", "torstai", "torstaina"},
	Friday:                []string{"pe", "perjantai", "perjantaina"},
	Saturday:              []string{"la", "lauantai", "lauantaina"},
	Sunday:                []string{"su", "sunnuntai", "sunnuntaina"},
	AM:                    []string{"ap"},
	PM:                    []string{"ip"},
	Year:                  []string{"v", "vuoden", "vuonna", "vuosi", "vuotta", "vv"},
	Month:                 []string{"kk", "kuukauden", "kuukausi", "kuukautta"},
	Week:                  []string{"viikko", "viikkoa", "viikon", "vk", "vko"},
	Day:                   []string{"p", "pv", "pvä", "pvää", "päivä", "päivän", "päivää"},
	Hour:                  []string{"t", "tunnin", "tunti", "tuntia"},
	Minute:                []string{"min", "minuutin", "minuutti", "minuuttia"},
	Second:                []string{"s", "sekunnin", "sekunti", "sekuntia", "sekuntin", "sekuntti", "sekunttia"},
	In:                    []string{"kuluttua", "päästä"},
	Ago:                   []string{"sitten"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`tänään`},
		`0 hour ago`:   {`tunnin sisällä`, `tämän tunnin aikana`},
		`0 minute ago`: {`minuutin sisällä`, `tämän minuutin aikana`},
		`0 month ago`:  {`tässä kk`, `tässä kuussa`},
		`0 second ago`: {`nyt`},
		`0 week ago`:   {`tällä viikolla`, `tällä vk`},
		`0 year ago`:   {`tänä v`, `tänä vuonna`},
		`1 day ago`:    {`eilen`},
		`1 month ago`:  {`viime kk`, `viime kuussa`},
		`1 week ago`:   {`viime viikolla`, `viime vk`},
		`1 year ago`:   {`viime v`, `viime vuonna`},
		`2 day ago`:    {`toissa päivänä`},
		`2 month ago`:  {`toissa kuussa`},
		`2 week ago`:   {`toissa viikolla`},
		`2 year ago`:   {`toissa vuonna`},
		`in 1 day`:     {`huom`, `huomenna`},
		`in 1 month`:   {`ensi kk`, `ensi kuussa`},
		`in 1 week`:    {`ensi viikolla`, `ensi vk`},
		`in 1 year`:    {`ensi v`, `ensi vuonna`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) pv sitten`),
			regexp.MustCompile(`(?i)(\d+) päivä sitten`),
			regexp.MustCompile(`(?i)(\d+) päivää sitten`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) t sitten`),
			regexp.MustCompile(`(?i)(\d+) tunti sitten`),
			regexp.MustCompile(`(?i)(\d+) tuntia sitten`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min sitten`),
			regexp.MustCompile(`(?i)(\d+) minuutti sitten`),
			regexp.MustCompile(`(?i)(\d+) minuuttia sitten`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) kk sitten`),
			regexp.MustCompile(`(?i)(\d+) kuukausi sitten`),
			regexp.MustCompile(`(?i)(\d+) kuukautta sitten`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) s sitten`),
			regexp.MustCompile(`(?i)(\d+) sekunti sitten`),
			regexp.MustCompile(`(?i)(\d+) sekuntia sitten`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) viikko sitten`),
			regexp.MustCompile(`(?i)(\d+) viikkoa sitten`),
			regexp.MustCompile(`(?i)(\d+) vk sitten`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) v sitten`),
			regexp.MustCompile(`(?i)(\d+) vuosi sitten`),
			regexp.MustCompile(`(?i)(\d+) vuotta sitten`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) pv päästä`),
			regexp.MustCompile(`(?i)(\d+) päivän päästä`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) t päästä`),
			regexp.MustCompile(`(?i)(\d+) tunnin päästä`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) min päästä`),
			regexp.MustCompile(`(?i)(\d+) minuutin päästä`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) kk päästä`),
			regexp.MustCompile(`(?i)(\d+) kuukauden päästä`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) s päästä`),
			regexp.MustCompile(`(?i)(\d+) sekunnin päästä`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) viikon päästä`),
			regexp.MustCompile(`(?i)(\d+) vk päästä`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) v päästä`),
			regexp.MustCompile(`(?i)(\d+) vuoden päästä`),
		},
	},
	Simplifications: map[string]string{
		`(\d+) (sekunnin|sekuntin|minuutin|tunnin|päivän|viikon|kuukauden|vuoden) (päästä|kuluttua)`: `\3 \1 \2`,
	},
}
