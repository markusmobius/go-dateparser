package data

import "regexp"

var loLocale = LocaleData{
	Name:                  "lo",
	DateOrder:             "DMY",
	January:               []string{"ມກ", "ມັງກອນ"},
	February:              []string{"ກພ", "ກຸມພາ"},
	March:                 []string{"ມນ", "ມີນາ"},
	April:                 []string{"ມສ", "ເມສາ"},
	May:                   []string{"ພພ", "ພຶດສະພາ"},
	June:                  []string{"ມິຖ", "ມິຖຸນາ"},
	July:                  []string{"ກລ", "ກໍລະກົດ"},
	August:                []string{"ສຫ", "ສິງຫາ"},
	September:             []string{"ກຍ", "ກັນຍາ"},
	October:               []string{"ຕລ", "ຕຸລາ"},
	November:              []string{"ພຈ", "ພະຈິກ"},
	December:              []string{"ທວ", "ທັນວາ"},
	Monday:                []string{"ຈັນ", "ວັນຈັນ"},
	Tuesday:               []string{"ວັນອັງຄານ", "ອັງຄານ"},
	Wednesday:             []string{"ພຸດ", "ວັນພຸດ"},
	Thursday:              []string{"ພະຫັດ", "ວັນພະຫັດ"},
	Friday:                []string{"ວັນສຸກ", "ສຸກ"},
	Saturday:              []string{"ວັນເສົາ", "ເສົາ"},
	Sunday:                []string{"ວັນອາທິດ", "ອາທິດ"},
	AM:                    []string{"ກ່ອນທ່ຽງ"},
	PM:                    []string{"ຫຼັງທ່ຽງ"},
	Year:                  []string{"ປີ"},
	Month:                 []string{"ດ", "ເດືອນ"},
	Week:                  []string{"ອ", "ອາທິດ"},
	Day:                   []string{"ມື້"},
	Hour:                  []string{"ຊມ", "ຊົ່ວໂມງ"},
	Minute:                []string{"ນທ", "ນາທີ"},
	Second:                []string{"ວິ", "ວິນາທີ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ມື້ນີ້`},
		`0 hour ago`:   {`ຊົ່ວໂມງນີ້`},
		`0 minute ago`: {`ນາທີນີ້`},
		`0 month ago`:  {`ເດືອນນີ້`},
		`0 second ago`: {`ຕອນນີ້`},
		`0 week ago`:   {`ອາທິດນີ້`},
		`0 year ago`:   {`ປີນີ້`},
		`1 day ago`:    {`ມື້ວານ`},
		`1 month ago`:  {`ເດືອນແລ້ວ`},
		`1 week ago`:   {`ອາທິດແລ້ວ`},
		`1 year ago`:   {`ປີກາຍ`},
		`in 1 day`:     {`ມື້ອື່ນ`},
		`in 1 month`:   {`ເດືອນໜ້າ`},
		`in 1 week`:    {`ອາທິດໜ້າ`},
		`in 1 year`:    {`ປີໜ້າ`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ມື້ກ່ອນ`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ຊມ ກ່ອນ`),
			regexp.MustCompile(`(?i)(\d+) ຊົ່ວໂມງກ່ອນ`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) ນທ ກ່ອນ`),
			regexp.MustCompile(`(?i)(\d+) ນາທີກ່ອນ`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ດ ກ່ອນ`),
			regexp.MustCompile(`(?i)(\d+) ເດືອນກ່ອນ`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) ວິ ກ່ອນ`),
			regexp.MustCompile(`(?i)(\d+) ວິນາທີກ່ອນ`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ອທ ກ່ອນ`),
			regexp.MustCompile(`(?i)(\d+) ອາທິດກ່ອນ`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ປີກ່ອນ`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ມື້`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ຊມ`),
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ຊົ່ວໂມງ`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) ໃນອີກ 0 ນາທີ`),
			regexp.MustCompile(`(?i)ໃນ (\d+) ນທ`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ດ`),
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ເດືອນ`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)ໃນ (\d+) ວິ`),
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ວິນາທີ`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ອທ`),
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ອາທິດ`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)ໃນອີກ (\d+) ປີ`),
		},
	},
}
