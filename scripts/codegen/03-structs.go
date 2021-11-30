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
