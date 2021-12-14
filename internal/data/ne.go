// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ne_Locale = LocaleData{
	Name:      "ne",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) सेकेण्ड पहिले(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) घण्टा पहिले(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) महिना पहिले(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) हप्ता पहिले(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) दिन पहिले(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) सेकेण्डमा(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) वर्ष अघि(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) घण्टामा(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) महिनामा(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) मिनेटमा(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) हप्तामा(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) वर्षमा(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) दिनमा(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)अर्को महिना(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)यही मिनेटमा(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)अर्को वर्ष(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)आउने हप्ता(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)सेप्टेम्बर(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)पूर्वाह्न(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)फेब्रुअरी(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)गत महिना(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)गत हप्ता(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)डिसेम्बर(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)नोभेम्बर(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)मङ्गलबार(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)यो घडीमा(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)यो महिना(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)यो हप्ता(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)शुक्रबार(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)अक्टोबर(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)अपराह्न(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)गत वर्ष(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)बिहिबार(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)यो वर्ष(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)सेकेन्ड(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)अप्रिल(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)आइतबार(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)बुधबार(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)शनिबार(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)सोमबार(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)अगस्ट(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)घण्टा(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)जनवरी(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)जुलाई(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)मङ्गल(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)महिना(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)मार्च(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)शुक्र(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)हप्ता(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)बर्ष(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)बिहि(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)भोलि(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)वर्ष(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)हिजो(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)आइत(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)जुन(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)बार(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)बुध(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)शनि(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)सोम(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)अब(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)आज(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)मई(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)मे(\z|\W|_)`), "${1}may${2}"},
	},
}

var ne_IN_Locale = LocaleData{
	Name:      "ne-IN",
	Parent:    &ne_Locale,
	DateOrder: "YMD",
}
