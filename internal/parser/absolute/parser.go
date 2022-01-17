package absolute

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/tokenizer"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

var (
	rxHourMinute  = regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`)
	rxMicrosecond = regexp.MustCompile(`\d{1,6}`)
	rxMeridian    = regexp.MustCompile(`(?i)am|pm`)

	alphaDirectives = map[string][]string{
		"weekday": {"monday", "mon"},
		"month":   {"january", "jan"},
	}

	weekdayNames = []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
)

type FilteredToken struct {
	tokenizer.Token
	OriginalIndex int
}

type TokenParseResult struct {
	Component string
	Value     int
}

type Parser struct {
	Config *setting.Configuration
	Now    time.Time

	Tokens         []tokenizer.Token
	FilteredTokens []FilteredToken

	Components      []string
	ComponentValues map[string]int
	ComponentTokens map[string]tokenizer.Token
	ParsedTime      time.Time

	AutoOrder   []string
	UnsetTokens []tokenizer.Token

	SkippedComponent string
	SkippedIndexes   map[int]struct{}
	SkippedTokens    strutil.Dict

	NumberDirectives [][]string
}

func NewParser(cfg *setting.Configuration, str string) Parser {
	// Prepare variables
	p := Parser{
		Config:          cfg,
		Now:             time.Now(),
		Tokens:          tokenizer.Tokenize(str),
		ComponentValues: map[string]int{},
		ComponentTokens: map[string]tokenizer.Token{},
		SkippedIndexes:  map[int]struct{}{},
		SkippedTokens:   strutil.NewDict("t", "year", "hour", "minute"),
	}
	print("TOKENS:", jsonify(&p.Tokens))

	// Set current time if specified on config
	if cfg != nil && !cfg.CurrentTime.IsZero() {
		p.Now = cfg.CurrentTime
	}

	// Filter tokens to exclude non-letter and non-digit
	for i, token := range p.Tokens {
		if token.Type <= tokenizer.Letter {
			p.FilteredTokens = append(p.FilteredTokens, FilteredToken{
				Token:         token,
				OriginalIndex: i,
			})
		}
	}
	print("FILTERED TOKENS:", jsonify(&p.FilteredTokens))

	// Prepare list of number directives
	if cfg != nil {
		for _, char := range strings.ToLower(cfg.DateOrder) {
			switch char {
			case 'd':
				p.Components = append(p.Components, "day")
				p.NumberDirectives = append(p.NumberDirectives, []string{"2"})
			case 'm':
				p.Components = append(p.Components, "month")
				p.NumberDirectives = append(p.NumberDirectives, []string{"1"})
			case 'y':
				p.Components = append(p.Components, "year")
				p.NumberDirectives = append(p.NumberDirectives, []string{"2006", "06"})
			}
		}
	}

	// Process each filtered token
	for index, filteredToken := range p.FilteredTokens {
		print("INDEX:", index, "TTOI:", jsonify(&filteredToken))

		// Check if this token should be skipped
		if _, exist := p.SkippedIndexes[index]; exist {
			continue
		}

		if p.SkippedTokens.Contain(filteredToken.Text) {
			continue
		}

		// Try to parse time
		if p.ParsedTime.IsZero() {
			meridianIndex := index + 1

			// Try case where hours and minutes are separated by a period. Example: 13.20.
			var isBeforePeriod bool
			if t, exist := p.getToken(filteredToken.OriginalIndex + 1); exist {
				isBeforePeriod = t.Text == "."
			}

			var isAfterPeriod bool
			if t, exist := p.getToken(filteredToken.OriginalIndex - 1); exist {
				isAfterPeriod = t.Text == "."
			}

			nextFilteredToken, exist := p.getFilteredToken(index + 1)
			if exist && isBeforePeriod && !isAfterPeriod {
				nextTokenIsLast := index+1 == len(p.FilteredTokens)-1
				nextOriginalToken, exist := p.getToken(nextFilteredToken.OriginalIndex + 1)
				nextOriginalTokenIsNotPeriod := !exist || nextOriginalToken.Text != "."

				if nextTokenIsLast || nextOriginalTokenIsNotPeriod {
					newToken := filteredToken.Text + ":" + nextOriginalToken.Text
					if rxHourMinute.MatchString(newToken) {
						filteredToken.Text = newToken
						p.SkippedIndexes[index+1] = struct{}{}
						meridianIndex++
					}
				}
			}

			// Capture microsecond
			var microsecond string
			if nextFilteredToken, exist := p.getFilteredToken(index + 1); exist {
				nextOriginalIndex := -2
				for i, ot := range p.Tokens {
					if ot.Text == filteredToken.Text && ot.Type == tokenizer.Digit {
						nextOriginalIndex = i
						break
					}
				}

				nextOriginalToken, _ := p.getToken(nextOriginalIndex + 1)
				nextOriginalToken, _ = p.getToken(filteredToken.OriginalIndex + 1) // DIFF

				if strings.Contains(filteredToken.Text, ":") &&
					strings.Contains(nextOriginalToken.Text, ".") {
					microsecond = rxMicrosecond.FindString(nextFilteredToken.Text)
				}
			}

			if microsecond != "" {
				meridianIndex++
			}

			// Capture meridian
			var meridian string
			if meridianToken, exist := p.getFilteredToken(meridianIndex); exist {
				meridian = rxMeridian.FindString(meridianToken.Text)
			}

			// Capture time
			if strings.Contains(filteredToken.Text, ":") ||
				strings.Contains(meridian, ":") ||
				strings.Contains(microsecond, ":") {
				var strTime string
				tokenText := filteredToken.Text

				if meridian != "" && microsecond == "" {
					strTime = fmt.Sprintf("%s %s", tokenText, meridian)
					p.SkippedIndexes[meridianIndex] = struct{}{}
				} else if microsecond != "" && meridian == "" {
					strTime = fmt.Sprintf("%s.%s", tokenText, microsecond)
					p.SkippedIndexes[index+1] = struct{}{}
				} else if meridian != "" && microsecond != "" {
					strTime = fmt.Sprintf("%s.%s %s", tokenText, microsecond, meridian)
					p.SkippedIndexes[index+1] = struct{}{}
					p.SkippedIndexes[meridianIndex] = struct{}{}
				} else {
					strTime = tokenText
				}

				p.ParsedTime, _ = common.ParseTime(strTime)
				p.ComponentTokens["time"] = tokenizer.Token{Text: strTime, Type: tokenizer.Digit}
				continue
			}
		}

		var results []TokenParseResult
		if filteredToken.Type == tokenizer.Digit {
			results, _ = p.parseDigitToken(filteredToken.Text, p.SkippedComponent)
		} else if filteredToken.Type == tokenizer.Letter {
			results, _ = p.parseLetterToken(filteredToken.Text, p.SkippedComponent)
		}

		// Save the initial token parsing result
		print("RESULTS:", jsonify(&results))
		for _, result := range results {
			if len(filteredToken.Text) == 4 && result.Component == "year" {
				p.SkippedComponent = "year"
			}
			p.ComponentValues[result.Component] = result.Value
		}
	}

	// Use fallback for unknown component
	for _, component := range []string{"day", "month", "year"} {
		if _, exist := p.ComponentValues[component]; !exist {
			for _, ut := range p.UnsetTokens {
				if ut.Type == tokenizer.Digit {
					newValue, _ := strconv.Atoi(ut.Text)
					p.ComponentTokens[component] = ut
					p.ComponentValues[component] = newValue
				}
			}
		}
	}

	return p
}

func (p *Parser) Parse() (date.Date, error) {
	// Find missing date fields
	missingParts := strutil.NewDict()
	for _, field := range []string{"day", "month", "year"} {
		if _, exist := p.ComponentValues[field]; !exist {
			missingParts.Add(field)
		}
	}

	print("MISSING PARTS:", jsonify(&missingParts))
	err := common.CheckStrictParsing(p.Config, missingParts)
	if err != nil {
		return date.Date{}, err
	}

	// Generate time result
	result := p.createTimeFromComponents()
	if !p.ParsedTime.IsZero() {
		result = time.Date(result.Year(), result.Month(), result.Day(),
			p.ParsedTime.Hour(), p.ParsedTime.Minute(), p.ParsedTime.Second(),
			p.ParsedTime.Nanosecond(), p.ParsedTime.Location())
	}

	// Apply correction for past and future
	result = p.correctForTimeFrame(result)

	// Apply correction for preference of day: beginning, current, end
	result = p.correctForDay(result)

	return date.Date{
		Time:   result,
		Period: p.getPeriod(),
	}, nil
}

func (p *Parser) getToken(i int) (tokenizer.Token, bool) {
	if i >= 0 && i < len(p.Tokens) {
		return p.Tokens[i], true
	}
	return tokenizer.Token{}, false
}

func (p *Parser) getFilteredToken(i int) (FilteredToken, bool) {
	if i >= 0 && i < len(p.FilteredTokens) {
		return p.FilteredTokens[i], true
	}
	return FilteredToken{}, false
}

func (p *Parser) setAndReturn(component string, token tokenizer.Token, date time.Time, skipDateOrder bool) ([]TokenParseResult, error) {
	if !skipDateOrder {
		p.AutoOrder = append(p.AutoOrder, component)
	}

	p.ComponentTokens[component] = token
	return []TokenParseResult{{
		Component: component,
		Value:     p.ComponentValues[component],
	}}, nil
}

func (p *Parser) parseDigitToken(tokenText string, skippedComponent string) ([]TokenParseResult, error) {
	// Prepare token data for return later
	token := tokenizer.Token{
		Text: tokenText,
		Type: tokenizer.Digit,
	}

	for i, directives := range p.NumberDirectives {
		// Check if this component need to be skipped
		component := p.Components[i]
		if component == skippedComponent {
			continue
		}

		// Fetch previous value
		prevValue := p.ComponentValues[component]
		prevToken := p.ComponentTokens[component]

		// Try each directive
		for _, directive := range directives {
			date, _ := time.Parse(directive, tokenText)
			if date.IsZero() {
				continue
			}

			if prevValue == 0 {
				return p.setAndReturn(component, token, date, false)
			} else if prevToken.Type == tokenizer.Digit {
				prevDate, _ := time.Parse(directive, prevToken.Text)
				if prevDate.IsZero() {
					p.UnsetTokens = append(p.UnsetTokens, prevToken)
					return p.setAndReturn(component, token, date, false)
				}
			}
		}
	}

	return nil, fmt.Errorf("unable to parse %s", tokenText)
}

func (p *Parser) parseLetterToken(tokenText string, skippedComponent string) ([]TokenParseResult, error) {
	// Prepare token data for return later
	token := tokenizer.Token{
		Text: tokenText,
		Type: tokenizer.Letter,
	}

	tokenText = strings.ToLower(tokenText)
	for component, directives := range alphaDirectives {
		// Check if this component need to be skipped
		if component == skippedComponent {
			continue
		}

		// Fetch previous value
		prevValue := p.ComponentValues[component]

		// Try each directive
		for _, directive := range directives {
			date, _ := time.Parse(directive, tokenText)
			if date.IsZero() {
				continue
			}

			if prevValue == 0 {
				return p.setAndReturn(component, token, date, true)
			} else if component == "month" {
				for j := range p.AutoOrder {
					if p.AutoOrder[j] == "month" {
						p.AutoOrder[j] = "day"
						break
					}
				}

				p.ComponentTokens["day"] = p.ComponentTokens["month"]
				p.ComponentTokens["month"] = token
				return []TokenParseResult{
					{Component: "month", Value: int(date.Month())},
					{Component: "day", Value: prevValue},
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to parse %s", tokenText)
}

func (p *Parser) createTimeFromComponents() time.Time {
	day, dayExist := p.ComponentValues["day"]
	if !dayExist || day == 0 {
		day = p.Now.Day()
	}

	month, monthExist := p.ComponentValues["month"]
	if !monthExist || month == 0 {
		month = int(p.Now.Month())
	}

	year, yearExist := p.ComponentValues["year"]
	if !yearExist || year == 0 {
		year = p.Now.Year()
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, p.Now.Location())
}

func (p *Parser) correctForTimeFrame(t time.Time) time.Time {
	_, yearExist := p.ComponentValues["year"]
	_, monthExist := p.ComponentValues["month"]
	_, tokenTimeExist := p.ComponentTokens["time"]
	_, tokenDayExist := p.ComponentTokens["day"]
	_, tokenMonthExist := p.ComponentTokens["month"]
	tokenYear, tokenYearExist := p.ComponentTokens["year"]
	tokenWeekday, tokenWeekdayExist := p.ComponentTokens["weekday"]

	var dateSource setting.PreferredDateSource
	if p.Config != nil {
		dateSource = p.Config.PreferredDateSource
	}

	if tokenWeekdayExist && !tokenYearExist && !tokenMonthExist && !tokenDayExist {
		dayIndex := int(t.Weekday())
		day := strings.ToLower(tokenWeekday.Text[:3])

		var steps int
		if dateSource == setting.Future {
			if weekdayNames[dayIndex] == day {
				steps = 7
			} else {
				for weekdayNames[dayIndex] != day {
					dayIndex = (dayIndex + 1) % 7
					steps++
				}
			}
		} else {
			if weekdayNames[dayIndex] == day {
				if dateSource == setting.Past {
					steps = 7
				} else {
					steps = 0
				}
			} else {
				for weekdayNames[dayIndex] != day {
					dayIndex--
					if dayIndex < 0 {
						dayIndex = 6
					}
					steps++
				}
			}
			steps *= -1
		}

		t = t.AddDate(0, 0, steps)
	}

	if monthExist && !yearExist {
		if p.Now.Before(t) {
			if dateSource == setting.Past {
				t = t.AddDate(-1, 0, 0)
			}
		} else {
			if dateSource == setting.Future {
				t = t.AddDate(1, 0, 0)
			}
		}
	}

	if tokenYearExist && len(tokenYear.Text) == 2 {
		if p.Now.Before(t) {
			if dateSource == setting.Past {
				t = t.AddDate(-100, 0, 0)
			}
		} else {
			if dateSource == setting.Future {
				t = t.AddDate(100, 0, 0)
			}
		}
	}

	if tokenTimeExist && !tokenYearExist && !tokenMonthExist && !tokenDayExist && !tokenWeekdayExist {
		if dateSource == setting.Past && p.Now.Before(t) {
			t = t.AddDate(0, 0, -1)
		}

		if dateSource == setting.Future && p.Now.After(t) {
			t = t.AddDate(0, 0, 1)
		}
	}

	return t
}

func (p *Parser) correctForDay(t time.Time) time.Time {
	_, tokenDayExist := p.ComponentTokens["day"]
	_, tokenTimeExist := p.ComponentTokens["time"]
	_, tokenWeekdayExist := p.ComponentTokens["weekday"]
	if tokenDayExist || tokenTimeExist || tokenWeekdayExist {
		return t
	}

	t = common.ApplyDayFromConfig(p.Config, t, p.Now.Day())
	return t
}

func (p *Parser) getPeriod() date.Period {
	_, timeExist := p.ComponentValues["time"]
	_, dayExist := p.ComponentValues["day"]
	_, monthExist := p.ComponentValues["month"]
	_, yearExist := p.ComponentValues["year"]
	returnTimeAsPeriod := p.Config != nil && p.Config.ReturnTimeAsPeriod

	switch {
	case returnTimeAsPeriod && timeExist:
		return date.Time
	case timeExist || dayExist:
		return date.Day
	case monthExist:
		return date.Month
	case yearExist:
		return date.Year
	}

	return date.Day
}
