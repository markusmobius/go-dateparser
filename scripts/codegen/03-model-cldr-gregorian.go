package main

import (
	"encoding/json"
	"fmt"
)

type CldrGregorianContainer struct {
	Data CldrGregorian `json:"data"`
}

func (v *CldrGregorianContainer) UnmarshalJSON(b []byte) error {
	var data struct {
		Main map[string]CldrGregorian `json:"main"`
	}

	if err := json.Unmarshal(b, &data); err == nil {
		for key, greg := range data.Main {
			greg.Name = key
			*&v.Data = greg
			break
		}
		return nil
	}

	return fmt.Errorf("invalid data format")
}

type CldrGregorian struct {
	Name     string `json:"name"`
	Identity struct {
		Language string `json:"language"`
	} `json:"identity"`
	Dates struct {
		Calendars struct {
			Gregorian struct {
				Months struct {
					Format struct {
						Abbreviated CldrGregorianMonths `json:"abbreviated"`
						Wide        CldrGregorianMonths `json:"wide"`
					} `json:"format"`
					StandAlone struct {
						Abbreviated CldrGregorianMonths `json:"abbreviated"`
						Wide        CldrGregorianMonths `json:"wide"`
					} `json:"stand-alone"`
				} `json:"months"`
				Days struct {
					Format struct {
						Abbreviated CldrGregorianDays `json:"abbreviated"`
						Wide        CldrGregorianDays `json:"wide"`
					} `json:"format"`
					StandAlone struct {
						Abbreviated CldrGregorianDays `json:"abbreviated"`
						Wide        CldrGregorianDays `json:"wide"`
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
				DateFormats CldrGregorianDateTimeFormats `json:"dateFormats"`
				TimeFormats CldrGregorianDateTimeFormats `json:"timeFormats"`
			} `json:"gregorian"`
		} `json:"calendars"`
	} `json:"dates"`
}

type CldrGregorianMonths struct {
	Jan string `json:"1"`
	Feb string `json:"2"`
	Mar string `json:"3"`
	Apr string `json:"4"`
	May string `json:"5"`
	Jun string `json:"6"`
	Jul string `json:"7"`
	Aug string `json:"8"`
	Sep string `json:"9"`
	Oct string `json:"10"`
	Nov string `json:"11"`
	Dec string `json:"12"`
}

func (c CldrGregorianMonths) List() []string {
	return []string{
		c.Jan, c.Feb, c.Mar, c.Apr, c.May, c.Jun,
		c.Jul, c.Aug, c.Sep, c.Oct, c.Nov, c.Dec,
	}
}

type CldrGregorianDays struct {
	Sun string `json:"sun"`
	Mon string `json:"mon"`
	Tue string `json:"tue"`
	Wed string `json:"wed"`
	Thu string `json:"thu"`
	Fri string `json:"fri"`
	Sat string `json:"sat"`
}

func (c CldrGregorianDays) List() []string {
	return []string{
		c.Sun, c.Mon, c.Tue, c.Wed,
		c.Thu, c.Fri, c.Sat,
	}
}

type CldrGregorianDayPeriods struct {
	Noon         string `json:"noon"`
	Midnight     string `json:"midnight"`
	Am           string `json:"am"`
	Pm           string `json:"pm"`
	AmAltVariant string `json:"am-alt-variant"`
	PmAltVariant string `json:"pm-alt-variant"`
}

type CldrGregorianDateTimeFormats struct {
	Full      CldrGregorianField `json:"full"`
	Long      CldrGregorianField `json:"long"`
	Medium    CldrGregorianField `json:"medium"`
	Short     CldrGregorianField `json:"short"`
	FullAlt   CldrGregorianField `json:"full-alt-variant"`
	LongAlt   CldrGregorianField `json:"long-alt-variant"`
	MediumAlt CldrGregorianField `json:"medium-alt-variant"`
	ShortAlt  CldrGregorianField `json:"short-alt-variant"`
}

type CldrGregorianField struct {
	Value string `json:"value"`
}

func (v *CldrGregorianField) UnmarshalJSON(b []byte) error {
	// Try unmarshalling as string
	var name string
	if err := json.Unmarshal(b, &name); err == nil {
		v.Value = name
		return nil
	}

	// Try unmarshalling as struct
	var object struct {
		Value string `json:"_value"`
	}

	if err := json.Unmarshal(b, &object); err == nil {
		v.Value = object.Value
		return nil
	}

	return fmt.Errorf("invalid data format")
}
