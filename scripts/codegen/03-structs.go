package main

type LocaleData struct {
	Name              string                `json:"name,omitempty"`
	DateOrder         string                `json:"date_order,omitempty"`
	January           []string              `json:"january,omitempty"`
	February          []string              `json:"february,omitempty"`
	March             []string              `json:"march,omitempty"`
	April             []string              `json:"april,omitempty"`
	May               []string              `json:"may,omitempty"`
	June              []string              `json:"june,omitempty"`
	July              []string              `json:"july,omitempty"`
	August            []string              `json:"august,omitempty"`
	September         []string              `json:"september,omitempty"`
	October           []string              `json:"october,omitempty"`
	November          []string              `json:"november,omitempty"`
	December          []string              `json:"december,omitempty"`
	Monday            []string              `json:"monday,omitempty"`
	Tuesday           []string              `json:"tuesday,omitempty"`
	Wednesday         []string              `json:"wednesday,omitempty"`
	Thursday          []string              `json:"thursday,omitempty"`
	Friday            []string              `json:"friday,omitempty"`
	Saturday          []string              `json:"saturday,omitempty"`
	Sunday            []string              `json:"sunday,omitempty"`
	AM                []string              `json:"am,omitempty"`
	PM                []string              `json:"pm,omitempty"`
	Year              []string              `json:"year,omitempty"`
	Month             []string              `json:"month,omitempty"`
	Week              []string              `json:"week,omitempty"`
	Day               []string              `json:"day,omitempty"`
	Hour              []string              `json:"hour,omitempty"`
	Minute            []string              `json:"minute,omitempty"`
	Second            []string              `json:"second,omitempty"`
	RelativeType      map[string][]string   `json:"relative-type,omitempty"`
	RelativeTypeRegex map[string][]string   `json:"relative-type-regex,omitempty"`
	LocaleSpecific    map[string]LocaleData `json:"locale_specific,omitempty"`
}

type CldrTerritoryData struct {
	Supplemental struct {
		TerritoryInfo map[string]struct {
			Gdp                string `json:"_gdp"`
			LiteracyPercent    string `json:"_literacyPercent"`
			Population         string `json:"_population"`
			LanguagePopulation map[string]struct {
				PopulationPercent string `json:"_populationPercent"`
				OfficialStatus    string `json:"_officialStatus"`
			} `json:"languagePopulation"`
		} `json:"territoryInfo"`
	} `json:"supplemental"`
}

type CldrGregorianData struct {
	Main map[string]struct {
		Dates struct {
			Calendars struct {
				Gregorian struct {
					Months      CldrGregorianDataPart `json:"months"`
					Days        CldrGregorianDataPart `json:"days"`
					DayPeriods  CldrGregorianDataPart `json:"dayPeriods"`
					DateFormats struct {
						ShortItf interface{} `json:"short"`
						Short    string      `json:"-"`
					} `json:"dateFormats"`
				} `json:"gregorian"`
			} `json:"calendars"`
		} `json:"dates"`
	} `json:"main"`
}

type CldrGregorianDataPart struct {
	Format struct {
		Abbreviated map[string]string `json:"abbreviated"`
		Wide        map[string]string `json:"wide"`
	} `json:"format"`
	StandAlone struct {
		Abbreviated map[string]string `json:"abbreviated"`
		Wide        map[string]string `json:"wide"`
	} `json:"stand-alone"`
}

type CldrDateFieldsData struct {
	Main map[string]struct {
		Dates struct {
			Fields map[string]CldrDateFieldsPart `json:"fields"`
		} `json:"dates"`
	} `json:"main"`
}

type CldrDateFieldsPart struct {
	DisplayName            string                       `json:"displayName"`
	RelativeTypeMin1       string                       `json:"relative-type--1"`
	RelativeType0          string                       `json:"relative-type-0"`
	RelativeType1          string                       `json:"relative-type-1"`
	RelativeTimeTypeFuture CldrDateFieldsRelTimePattern `json:"relativeTime-type-future"`
	RelativeTimeTypePast   CldrDateFieldsRelTimePattern `json:"relativeTime-type-past"`
}

type CldrDateFieldsRelTimePattern struct {
	CountOne   string `json:"relativeTimePattern-count-one"`
	CountOther string `json:"relativeTimePattern-count-other"`
}
