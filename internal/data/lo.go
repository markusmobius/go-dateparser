// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lo_Locale = LocaleData{
	Name:      "lo",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ຊົ່ວໂມງ(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ໃນອີກ 0 ນາທີ(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ວິນາທີ(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ຊົ່ວໂມງກ່ອນ(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ອາທິດ(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ເດືອນ(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ວິນາທີກ່ອນ(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ອາທິດກ່ອນ(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ເດືອນກ່ອນ(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ມື້(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ຊມ(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ປີ(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ອທ(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ຊມ ກ່ອນ(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ມື້ກ່ອນ(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ວິ ກ່ອນ(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ອທ ກ່ອນ(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນອີກ (\d+) ດ(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ດ ກ່ອນ(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ປີກ່ອນ(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນ (\d+) ນທ(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)ໃນ (\d+) ວິ(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)ຊົ່ວໂມງນີ້(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນອັງຄານ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ອາທິດຫນ້າ(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ອາທິດແລ້ວ(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ເດືອນຫນ້າ(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ເດືອນແລ້ວ(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກ່ອນທ່ຽງ(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນພະຫັດ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນອາທິດ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຫຼັງທ່ຽງ(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ອາທິດນີ້(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ເດືອນນີ້(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກໍລະກົດ(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຊົ່ວໂມງ(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)ນາທີນີ້(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ພຶດສະພາ(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມື້ອື່ນ(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນເສົາ(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຕອນນີ້(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ປີຫນ້າ(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມັງກອນ(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມິຖຸນາ(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມື້ນີ້(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມື້ວານ(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນຈັນ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນພຸດ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວັນສຸກ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວິນາທີ(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ອັງຄານ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກັນຍາ(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກຸມພາ(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ທັນວາ(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ປີກາຍ(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ປີນີ້(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ພະຈິກ(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ພະຫັດ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ສິງຫາ(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ອາທິດ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ເດືອນ(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຕຸລາ(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມີນາ(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)ເມສາ(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ເສົາ(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຈັນ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ພຸດ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມິຖ(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມື້(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ສຸກ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກຍ(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກພ(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ກລ(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຊມ(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)ຕລ(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ທວ(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ປີ(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ພຈ(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ພພ(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມກ(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມນ(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)ມສ(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ວິ(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ສຫ(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ດ(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ອ(\z|\W|_)`), "${1}week${2}"},
	},
}
