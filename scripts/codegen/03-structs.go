package main

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
					Months struct {
						Format struct {
							Abbreviated map[int]string `json:"abbreviated"`
							Wide        map[int]string `json:"wide"`
						} `json:"format"`
						StandAlone struct {
							Abbreviated map[int]string `json:"abbreviated"`
							Wide        map[int]string `json:"wide"`
						} `json:"stand-alone"`
					} `json:"months"`
					Days struct {
						Format struct {
							Abbreviated CldrGregorianDayFormat `json:"abbreviated"`
							Wide        CldrGregorianDayFormat `json:"wide"`
						} `json:"format"`
						StandAlone struct {
							Abbreviated CldrGregorianDayFormat `json:"abbreviated"`
							Wide        CldrGregorianDayFormat `json:"wide"`
						} `json:"stand-alone"`
					} `json:"days"`
					DayPeriods struct {
						Format struct {
							Abbreviated CldrGregorianDayPeriods `json:"abbreviated"`
							Wide        CldrGregorianDayPeriods `json:"wide"`
						} `json:"format"`
						StandAlone struct {
							Abbreviated CldrGregorianDayPeriods `json:"abbreviated"`
							Wide        CldrGregorianDayPeriods `json:"wide"`
						} `json:"stand-alone"`
					} `json:"dayPeriods"`
					DateFormats struct {
						ShortItf interface{} `json:"short"`
						Short    string      `json:"-"`
					} `json:"dateFormats"`
				} `json:"gregorian"`
			} `json:"calendars"`
		} `json:"dates"`
	} `json:"main"`
}

type CldrGregorianDayFormat struct {
	Sun string `json:"sun"`
	Mon string `json:"mon"`
	Tue string `json:"tue"`
	Wed string `json:"wed"`
	Thu string `json:"thu"`
	Fri string `json:"fri"`
	Sat string `json:"sat"`
}

type CldrGregorianDayPeriods struct {
	AM string `json:"am"`
	PM string `json:"pm"`
}

type CldrDateFieldsData struct {
	Main map[string]struct {
		Dates struct {
			Fields struct {
				Year         CldrDateFieldsPart `json:"year"`
				YearShort    CldrDateFieldsPart `json:"year-short"`
				YearNarrow   CldrDateFieldsPart `json:"year-narrow"`
				Month        CldrDateFieldsPart `json:"month"`
				MonthShort   CldrDateFieldsPart `json:"month-short"`
				MonthNarrow  CldrDateFieldsPart `json:"month-narrow"`
				Week         CldrDateFieldsPart `json:"week"`
				WeekShort    CldrDateFieldsPart `json:"week-short"`
				WeekNarrow   CldrDateFieldsPart `json:"week-narrow"`
				Day          CldrDateFieldsPart `json:"day"`
				DayShort     CldrDateFieldsPart `json:"day-short"`
				DayNarrow    CldrDateFieldsPart `json:"day-narrow"`
				Hour         CldrDateFieldsPart `json:"hour"`
				HourShort    CldrDateFieldsPart `json:"hour-short"`
				HourNarrow   CldrDateFieldsPart `json:"hour-narrow"`
				Minute       CldrDateFieldsPart `json:"minute"`
				MinuteShort  CldrDateFieldsPart `json:"minute-short"`
				MinuteNarrow CldrDateFieldsPart `json:"minute-narrow"`
				Second       CldrDateFieldsPart `json:"second"`
				SecondShort  CldrDateFieldsPart `json:"second-short"`
				SecondNarrow CldrDateFieldsPart `json:"second-narrow"`
			} `json:"fields"`
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
