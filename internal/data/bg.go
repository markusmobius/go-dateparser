// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bg_Locale = LocaleData{
	Name:      "bg",
	DateOrder: "DMY 'Г'.",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "върху", "до", "на", "около", "от", "под"},
	Simplifications: map[*regexp.Regexp]string{
		regexp.MustCompile(`(\A|\W|_)два(\z|\W|_)`):  "${1}2${2}",
		regexp.MustCompile(`(\A|\W|_)един(\z|\W|_)`): "${1}1${2}",
		regexp.MustCompile(`(\A|\W|_)три(\z|\W|_)`):  "${1}3${2}",
	},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) седмица(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) седмици(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) секунда(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) секунди(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)предходната седмица(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) година(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) години(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) месеца(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) седмица(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) седмици(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) секунда(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) секунди(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)след 1 десетилетие(\z|\W|_)`), "${1}in 10 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)следващата седмица(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) месец(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди десетилетие(\z|\W|_)`), "${1}10 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) година(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) години(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) месеца(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) минута(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) минути(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)следващата година(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)миналата седмица(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) седм(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) часа(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) месец(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)миналата година(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) ден(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) дни(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) сек(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) час(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)предходен месец(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) седм(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) часа(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) ден(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) дни(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) мин(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) сек(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) час(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)в тази минута(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)пр (\d+) седм(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) г(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) м(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди (\d+) ч(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди седмица(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) седм(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)следващ месец(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)пр (\d+) сек(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)преди година(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) мин(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) сек(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) г(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) м(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)след (\d+) ч(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)тази седмица(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)тази година(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)в този час(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)понеделник(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пр (\d+) г(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пр (\d+) д(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пр (\d+) м(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пр (\d+) ч(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) г(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) д(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) м(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)сл (\d+) ч(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)следв седм(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)този месец(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)вдругиден(\z|\W|_)`), "${1}in 2 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди ден(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди час(\z|\W|_)`), "${1}1 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)септември(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)следв мес(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)тази седм(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)четвъртък(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)декември(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)мин седм(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)октомври(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)онзи ден(\z|\W|_)`), "${1}2 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)след ден(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)след час(\z|\W|_)`), "${1}in 1 hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)този мес(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)февруари(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)вторник(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)мин мес(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ноември(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)седмица(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)седмици(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунда(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунди(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)сл седм(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)следв г(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)август(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)година(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)години(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)месеци(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)минути(\z|\W|_)`), "${1}minute${2}"},
		{regexp.MustCompile(`(\A|\W|_)неделя(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)събота(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)януари(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)април(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)вчера(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)месец(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)мин г(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)мин м(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)петък(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)подир(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)после(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)преди(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)септм(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)снощи(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)сряда(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)дена(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)днес(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)март(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)проб(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)сега(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)седм(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)септ(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)сл г(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)сл м(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)след(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)слоб(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)утре(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)часа(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)авг(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)апр(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)вто(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)год(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)дек(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ден(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)дни(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)май(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)мес(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ное(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)окт(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)пон(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сед(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)сек(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)сеп(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)сря(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)т г(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)т м(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)фев(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)час(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)юли(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)юни(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)яну(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ап(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)вт(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)нд(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пн(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пт(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сб(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ср(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)фв(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)чт(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)юл(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)юн(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ян(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)г(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)д(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)м(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)с(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)ч(\z|\W|_)`), "${1}hour${2}"},
	},
}
