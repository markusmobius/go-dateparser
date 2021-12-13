// Code is generated by script; DO NOT EDIT.

package data

var be_Locale = LocaleData{
	Name:      "be",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "каля", "у", "і", "ў"},
	Simplifications: map[string]string{
		`(\d+)\s*гадзін\s(\d+)\s*хвілін`: "$1:$2",
		`^гадзіна`:        "1 гадзіна",
		`^секунду`:        "1 секунду",
		`^хвіліну`:        "1 хвіліну",
		`гадзіну`:         "1 гадзіну",
		`некалькі секунд`: "44 секунды",
		`некалькі хвілін`: "2 хвіліны",
	},
	Translations: map[string]string{
		`(\A|\W|_)студзеня(\z|\W|_)`:           "$1january$2",
		`(\A|\W|_)сту(\z|\W|_)`:                "$1january$2",
		`(\A|\W|_)студзень(\z|\W|_)`:           "$1january$2",
		`(\A|\W|_)лютага(\z|\W|_)`:             "$1february$2",
		`(\A|\W|_)лют(\z|\W|_)`:                "$1february$2",
		`(\A|\W|_)люты(\z|\W|_)`:               "$1february$2",
		`(\A|\W|_)сакавіка(\z|\W|_)`:           "$1march$2",
		`(\A|\W|_)сак(\z|\W|_)`:                "$1march$2",
		`(\A|\W|_)сакавік(\z|\W|_)`:            "$1march$2",
		`(\A|\W|_)красавіка(\z|\W|_)`:          "$1april$2",
		`(\A|\W|_)кра(\z|\W|_)`:                "$1april$2",
		`(\A|\W|_)красавік(\z|\W|_)`:           "$1april$2",
		`(\A|\W|_)мая(\z|\W|_)`:                "$1may$2",
		`(\A|\W|_)май(\z|\W|_)`:                "$1may$2",
		`(\A|\W|_)чэрвеня(\z|\W|_)`:            "$1june$2",
		`(\A|\W|_)чэр(\z|\W|_)`:                "$1june$2",
		`(\A|\W|_)чэрвень(\z|\W|_)`:            "$1june$2",
		`(\A|\W|_)ліпеня(\z|\W|_)`:             "$1july$2",
		`(\A|\W|_)ліп(\z|\W|_)`:                "$1july$2",
		`(\A|\W|_)ліпень(\z|\W|_)`:             "$1july$2",
		`(\A|\W|_)жніўня(\z|\W|_)`:             "$1august$2",
		`(\A|\W|_)жні(\z|\W|_)`:                "$1august$2",
		`(\A|\W|_)жнівень(\z|\W|_)`:            "$1august$2",
		`(\A|\W|_)верасня(\z|\W|_)`:            "$1september$2",
		`(\A|\W|_)вер(\z|\W|_)`:                "$1september$2",
		`(\A|\W|_)верасень(\z|\W|_)`:           "$1september$2",
		`(\A|\W|_)кастрычніка(\z|\W|_)`:        "$1october$2",
		`(\A|\W|_)кас(\z|\W|_)`:                "$1october$2",
		`(\A|\W|_)кастрычнік(\z|\W|_)`:         "$1october$2",
		`(\A|\W|_)лістапада(\z|\W|_)`:          "$1november$2",
		`(\A|\W|_)ліс(\z|\W|_)`:                "$1november$2",
		`(\A|\W|_)лістапад(\z|\W|_)`:           "$1november$2",
		`(\A|\W|_)снежня(\z|\W|_)`:             "$1december$2",
		`(\A|\W|_)сне(\z|\W|_)`:                "$1december$2",
		`(\A|\W|_)снежань(\z|\W|_)`:            "$1december$2",
		`(\A|\W|_)панядзелак(\z|\W|_)`:         "$1monday$2",
		`(\A|\W|_)пн(\z|\W|_)`:                 "$1monday$2",
		`(\A|\W|_)аўторак(\z|\W|_)`:            "$1tuesday$2",
		`(\A|\W|_)аў(\z|\W|_)`:                 "$1tuesday$2",
		`(\A|\W|_)серада(\z|\W|_)`:             "$1wednesday$2",
		`(\A|\W|_)ср(\z|\W|_)`:                 "$1wednesday$2",
		`(\A|\W|_)чацвер(\z|\W|_)`:             "$1thursday$2",
		`(\A|\W|_)чц(\z|\W|_)`:                 "$1thursday$2",
		`(\A|\W|_)пятніца(\z|\W|_)`:            "$1friday$2",
		`(\A|\W|_)пт(\z|\W|_)`:                 "$1friday$2",
		`(\A|\W|_)субота(\z|\W|_)`:             "$1saturday$2",
		`(\A|\W|_)сб(\z|\W|_)`:                 "$1saturday$2",
		`(\A|\W|_)нядзеля(\z|\W|_)`:            "$1sunday$2",
		`(\A|\W|_)нд(\z|\W|_)`:                 "$1sunday$2",
		`(\A|\W|_)am(\z|\W|_)`:                 "$1am$2",
		`(\A|\W|_)pm(\z|\W|_)`:                 "$1pm$2",
		`(\A|\W|_)год(\z|\W|_)`:                "$1year$2",
		`(\A|\W|_)г(\z|\W|_)`:                  "$1year$2",
		`(\A|\W|_)месяц(\z|\W|_)`:              "$1month$2",
		`(\A|\W|_)мес(\z|\W|_)`:                "$1month$2",
		`(\A|\W|_)тыд(\z|\W|_)`:                "$1week$2",
		`(\A|\W|_)дзень(\z|\W|_)`:              "$1day$2",
		`(\A|\W|_)д(\z|\W|_)`:                  "$1day$2",
		`(\A|\W|_)гадзіна(\z|\W|_)`:            "$1hour$2",
		`(\A|\W|_)гадз(\z|\W|_)`:               "$1hour$2",
		`(\A|\W|_)секунда(\z|\W|_)`:            "$1second$2",
		`(\A|\W|_)с(\z|\W|_)`:                  "$1second$2",
		`(\A|\W|_)у мінулым годзе(\z|\W|_)`:    "$11 year ago$2",
		`(\A|\W|_)у гэтым годзе(\z|\W|_)`:      "$10 year ago$2",
		`(\A|\W|_)у наступным годзе(\z|\W|_)`:  "$1in 1 year$2",
		`(\A|\W|_)у мінулым месяцы(\z|\W|_)`:   "$11 month ago$2",
		`(\A|\W|_)у гэтым месяцы(\z|\W|_)`:     "$10 month ago$2",
		`(\A|\W|_)у наступным месяцы(\z|\W|_)`: "$1in 1 month$2",
		`(\A|\W|_)на мінулым тыдні(\z|\W|_)`:   "$11 week ago$2",
		`(\A|\W|_)на гэтым тыдні(\z|\W|_)`:     "$10 week ago$2",
		`(\A|\W|_)на наступным тыдні(\z|\W|_)`: "$1in 1 week$2",
		`(\A|\W|_)учора(\z|\W|_)`:              "$11 day ago$2",
		`(\A|\W|_)сёння(\z|\W|_)`:              "$10 day ago$2",
		`(\A|\W|_)заўтра(\z|\W|_)`:             "$1in 1 day$2",
		`(\A|\W|_)у гэту гадзіну(\z|\W|_)`:     "$10 hour ago$2",
		`(\A|\W|_)у гэту хвіліну(\z|\W|_)`:     "$10 minute ago$2",
		`(\A|\W|_)цяпер(\z|\W|_)`:              "$10 second ago$2",
		`(\A|\W|_)праз (\d+) год(\z|\W|_)`:     "$1in $2 year$3",
		`(\A|\W|_)праз (\d+) года(\z|\W|_)`:    "$1in $2 year$3",
		`(\A|\W|_)праз (\d+) г(\z|\W|_)`:       "$1in $2 year$3",
		`(\A|\W|_)(\d+) год таму(\z|\W|_)`:     "$1$2 year ago$3",
		`(\A|\W|_)(\d+) года таму(\z|\W|_)`:    "$1$2 year ago$3",
		`(\A|\W|_)(\d+) г таму(\z|\W|_)`:       "$1$2 year ago$3",
		`(\A|\W|_)праз (\d+) месяц(\z|\W|_)`:   "$1in $2 month$3",
		`(\A|\W|_)праз (\d+) месяца(\z|\W|_)`:  "$1in $2 month$3",
		`(\A|\W|_)праз (\d+) мес(\z|\W|_)`:     "$1in $2 month$3",
		`(\A|\W|_)(\d+) месяц таму(\z|\W|_)`:   "$1$2 month ago$3",
		`(\A|\W|_)(\d+) месяца таму(\z|\W|_)`:  "$1$2 month ago$3",
		`(\A|\W|_)(\d+) мес таму(\z|\W|_)`:     "$1$2 month ago$3",
		`(\A|\W|_)праз (\d+) тыдзень(\z|\W|_)`: "$1in $2 week$3",
		`(\A|\W|_)праз (\d+) тыдня(\z|\W|_)`:   "$1in $2 week$3",
		`(\A|\W|_)праз (\d+) тыд(\z|\W|_)`:     "$1in $2 week$3",
		`(\A|\W|_)(\d+) тыдзень таму(\z|\W|_)`: "$1$2 week ago$3",
		`(\A|\W|_)(\d+) тыдня таму(\z|\W|_)`:   "$1$2 week ago$3",
		`(\A|\W|_)(\d+) тыд таму(\z|\W|_)`:     "$1$2 week ago$3",
		`(\A|\W|_)праз (\d+) дзень(\z|\W|_)`:   "$1in $2 day$3",
		`(\A|\W|_)праз (\d+) дня(\z|\W|_)`:     "$1in $2 day$3",
		`(\A|\W|_)праз (\d+) д(\z|\W|_)`:       "$1in $2 day$3",
		`(\A|\W|_)(\d+) дзень таму(\z|\W|_)`:   "$1$2 day ago$3",
		`(\A|\W|_)(\d+) дня таму(\z|\W|_)`:     "$1$2 day ago$3",
		`(\A|\W|_)(\d+) д таму(\z|\W|_)`:       "$1$2 day ago$3",
		`(\A|\W|_)праз (\d+) гадзіну(\z|\W|_)`: "$1in $2 hour$3",
		`(\A|\W|_)праз (\d+) гадзіны(\z|\W|_)`: "$1in $2 hour$3",
		`(\A|\W|_)праз (\d+) гадз(\z|\W|_)`:    "$1in $2 hour$3",
		`(\A|\W|_)(\d+) гадзіну таму(\z|\W|_)`: "$1$2 hour ago$3",
		`(\A|\W|_)(\d+) гадзіны таму(\z|\W|_)`: "$1$2 hour ago$3",
		`(\A|\W|_)(\d+) гадз таму(\z|\W|_)`:    "$1$2 hour ago$3",
		`(\A|\W|_)праз (\d+) хвіліну(\z|\W|_)`: "$1in $2 minute$3",
		`(\A|\W|_)праз (\d+) хвіліны(\z|\W|_)`: "$1in $2 minute$3",
		`(\A|\W|_)праз (\d+) хв(\z|\W|_)`:      "$1in $2 minute$3",
		`(\A|\W|_)праз (\d+) секунду(\z|\W|_)`: "$1in $2 second$3",
		`(\A|\W|_)праз (\d+) секунды(\z|\W|_)`: "$1in $2 second$3",
		`(\A|\W|_)праз (\d+) с(\z|\W|_)`:       "$1in $2 second$3",
		`(\A|\W|_)(\d+) секунду таму(\z|\W|_)`: "$1$2 second ago$3",
		`(\A|\W|_)(\d+) секунды таму(\z|\W|_)`: "$1$2 second ago$3",
		`(\A|\W|_)(\d+) с таму(\z|\W|_)`:       "$1$2 second ago$3",
		`(\A|\W|_)стд(\z|\W|_)`:                "$1january$2",
		`(\A|\W|_)крс(\z|\W|_)`:                "$1april$2",
		`(\A|\W|_)траўня(\z|\W|_)`:             "$1may$2",
		`(\A|\W|_)тра(\z|\W|_)`:                "$1may$2",
		`(\A|\W|_)жнівеня(\z|\W|_)`:            "$1august$2",
		`(\A|\W|_)жнв(\z|\W|_)`:                "$1august$2",
		`(\A|\W|_)врс(\z|\W|_)`:                "$1september$2",
		`(\A|\W|_)кст(\z|\W|_)`:                "$1october$2",
		`(\A|\W|_)ліс(\z|\W|_)`:                "$1november$2",
		`(\A|\W|_)снж(\z|\W|_)`:                "$1december$2",
		`(\A|\W|_)пнд(\z|\W|_)`:                "$1monday$2",
		`(\A|\W|_)аўт(\z|\W|_)`:                "$1tuesday$2",
		`(\A|\W|_)чцв(\z|\W|_)`:                "$1thursday$2",
		`(\A|\W|_)чв(\z|\W|_)`:                 "$1thursday$2",
		`(\A|\W|_)пят(\z|\W|_)`:                "$1friday$2",
		`(\A|\W|_)суб(\z|\W|_)`:                "$1saturday$2",
		`(\A|\W|_)няд(\z|\W|_)`:                "$1sunday$2",
		`(\A|\W|_)гады(\z|\W|_)`:               "$1year$2",
		`(\A|\W|_)года(\z|\W|_)`:               "$1year$2",
		`(\A|\W|_)гадоў(\z|\W|_)`:              "$1year$2",
		`(\A|\W|_)месяца(\z|\W|_)`:             "$1month$2",
		`(\A|\W|_)месяцы(\z|\W|_)`:             "$1month$2",
		`(\A|\W|_)месяцаў(\z|\W|_)`:            "$1month$2",
		`(\A|\W|_)тыдзень(\z|\W|_)`:            "$1week$2",
		`(\A|\W|_)тыдня(\z|\W|_)`:              "$1week$2",
		`(\A|\W|_)тыдні(\z|\W|_)`:              "$1week$2",
		`(\A|\W|_)тыдняў(\z|\W|_)`:             "$1week$2",
		`(\A|\W|_)дні(\z|\W|_)`:                "$1day$2",
		`(\A|\W|_)дзён(\z|\W|_)`:               "$1day$2",
		`(\A|\W|_)дзен(\z|\W|_)`:               "$1day$2",
		`(\A|\W|_)гадзіны(\z|\W|_)`:            "$1hour$2",
		`(\A|\W|_)гадзіну(\z|\W|_)`:            "$1hour$2",
		`(\A|\W|_)гадзін(\z|\W|_)`:             "$1hour$2",
		`(\A|\W|_)хвілін(\z|\W|_)`:             "$1minute$2",
		`(\A|\W|_)хвіліны(\z|\W|_)`:            "$1minute$2",
		`(\A|\W|_)хвіліну(\z|\W|_)`:            "$1minute$2",
		`(\A|\W|_)хвіл(\z|\W|_)`:               "$1minute$2",
		`(\A|\W|_)секунды(\z|\W|_)`:            "$1second$2",
		`(\A|\W|_)секунду(\z|\W|_)`:            "$1second$2",
		`(\A|\W|_)секунд(\z|\W|_)`:             "$1second$2",
		`(\A|\W|_)сек(\z|\W|_)`:                "$1second$2",
		`(\A|\W|_)на працягу(\z|\W|_)`:         "$1in$2",
		`(\A|\W|_)таму назад(\z|\W|_)`:         "$1ago$2",
		`(\A|\W|_)таму(\z|\W|_)`:               "$1ago$2",
		`(\A|\W|_)назад(\z|\W|_)`:              "$1ago$2",
		`(\A|\W|_)сення(\z|\W|_)`:              "$10 day ago$2",
		`(\A|\W|_)пазаўчора(\z|\W|_)`:          "$12 day ago$2",
		`(\A|\W|_)ўчора(\z|\W|_)`:              "$11 day ago$2",
	},
}
