// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bs_Cyrl_Locale = LocaleData{
	Name:      "bs-Cyrl",
	DateOrder: "DMY.",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) секунди(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) секунди(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) година(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) годину(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) месеци(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) недеља(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) недељу(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) секунд(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) година(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) годину(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) месеци(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) минута(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) недеља(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) недељу(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) секунд(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) месец(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)следећег месеца(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) месец(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) минут(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) дана(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) сати(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)прошлог месеца(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)следеће године(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)следеће недеље(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) дана(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) сати(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) дан(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)пре (\d+) сат(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)прошле године(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)прошле недеље(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) дан(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)за (\d+) сат(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)овог месеца(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ове године(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ове недеље(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)понедељак(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пре подне(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)септембар(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)децембар(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)новембар(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)четвртак(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)октобар(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)поподне(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)сриједа(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)фебруар(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)август(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)година(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)недеља(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)секунд(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)субота(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)уторак(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)јануар(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)април(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)данас(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)месец(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)петак(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сутра(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)март(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)јули(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)јуни(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)јуче(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)авг(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)апр(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)дан(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)дец(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)мар(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)мај(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)нед(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)нов(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)окт(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)пет(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)пон(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)сеп(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)сри(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)суб(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)уто(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)феб(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)час(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)чет(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)јан(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)јул(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)јун(\z|\W|_)`), "${1}june${2}"},
	},
}
