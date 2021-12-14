// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var km_Locale = LocaleData{
	Name:      "km",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)ក្នុង​រយៈ​ពេល (\d+) ម៉ោង(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) សប្ដាហ៍​មុន(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) វិនាទី​មុន(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) សប្ដាហ៍ទៀត(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ឆ្នាំ​មុន(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ថ្ងៃ​​មុន(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) វិនាទីទៀត(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ឆ្នាំទៀត(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ថ្ងៃ​មុន(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ម៉ោង​មុន(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ថ្ងៃទៀត(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) នាទីទៀត(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ម៉ោងទៀត(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)សប្ដាហ៍​ក្រោយ(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ខែទៀត(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ខែមុន(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ឆ្នាំ​ក្រោយ(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)សប្ដាហ៍​នេះ(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)សប្ដាហ៍​មុន(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ថ្ងៃ​ស្អែក(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ព្រហស្បតិ៍(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ឆ្នាំ​នេះ(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ឆ្នាំ​មុន(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ថ្ងៃស្អែក(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ខែ​ក្រោយ(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ថ្ងៃ​នេះ(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ម្សិលមិញ(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)វិច្ឆិកា(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)នាទីនេះ(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ម៉ោងនេះ(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)សប្ដាហ៍(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)អាទិត្យ(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)កក្កដា(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)កុម្ភៈ(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ខែ​នេះ(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ខែ​មុន(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)មិថុនា(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)វិនាទី(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)អង្គារ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)កញ្ញា(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ច័ន្ទ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ឆ្នាំ(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)សុក្រ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)តុលា(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ថ្ងៃ(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ធ្នូ(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)មករា(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)មីនា(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)មេសា(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ម៉ោង(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)សីហា(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)សៅរ៍(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ឥឡូវ(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ឧសភា(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)ពុធ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ខែ(\z|\W|_)`), "${1}month${2}"},
	},
}
