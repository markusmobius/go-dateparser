package nospace

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/dateutil"
	"github.com/markusmobius/go-dateparser/internal/regexp"
)

var (
	rxCompatible = regexp.MustCompile(`\D+`)
	rxEightDigit = regexp.MustCompile(`^\d{8}$`)

	formatMapping = map[string]string{
		"d": "day",
		"m": "month",
		"y": "year",
		"Y": "year",
	}

	dateFormats = []string{
		"Ymd", "Ydm", "mYd",
		"mdY", "dYm", "dmY",
		"ymd", "ydm", "myd",
		"mdy", "dym", "dmy",
	}

	timeFormats = []string{
		"HMS",
		"HM",
		"H",
	}

	preferredFormats = []string{
		"YmdHM", "YmdHMS",
	}

	preferredFormatsOrdered8Digit = []string{
		"01022006", "02012006",
		"20060102", "20060201",
		"01200602", "02200601",
	}

	allFormats = func() []string {
		var mergedFormats []string
		for _, date := range dateFormats {
			for _, time := range timeFormats {
				mergedFormats = append(mergedFormats, date+time)
			}
		}

		formats := append(dateFormats, mergedFormats...)
		formats = append(formats, timeFormats...)
		return formats
	}()

	dateTimeFormats = func() map[string][]string {
		formats := map[string][]string{
			"mdy": createSortedFormats("mdy"),
			"myd": createSortedFormats("myd"),
			"ymd": createSortedFormats("ymd"),
			"ydm": createSortedFormats("ydm"),
			"dmy": createSortedFormats("dmy"),
			"dym": createSortedFormats("dym"),
		}

		formats["mdy"] = append(preferredFormats, formats["mdy"]...)
		return formats
	}()

	partSizeMap = map[string]int{
		"Y": 4, "y": 2, "m": 2, "d": 2,
		"H": 2, "M": 2, "S": 2,
	}
)

func createSortedFormats(prefix string) []string {
	list := append([]string{}, allFormats...)
	sort.SliceStable(list, func(i, j int) bool {
		itemI := strings.ToLower(list[i])
		itemJ := strings.ToLower(list[j])
		iHasPrefix := strings.HasPrefix(itemI, prefix)
		jHasPrefix := strings.HasPrefix(itemJ, prefix)
		return iHasPrefix && !jHasPrefix
	})

	return list
}

// createParseCandidates generate parse candidates and its Go layout from the specified format.
// For example, string `202112` can be parsed as:
// - `2021-1-12`
// - `2021-11-2`
// - `20-2-1112`
// - `2-02-1112`
func createParseCandidates(str string, format string) ([]string, string) {
	// Split format by its part
	var formatParts []string
	for _, part := range strings.Split(format, "") {
		if _, exist := partSizeMap[part]; exist {
			formatParts = append(formatParts, part)
		}
	}

	// Calculate permutation count
	nParts := len(formatParts)
	permutationCount := 1
	partsPermutationCount := map[string]int{}

	for i, part := range formatParts {
		partPermutationCount := 1
		for j := i + 1; j < nParts; j++ {
			partJ := formatParts[j]
			partPermutationCount *= partSizeMap[partJ]
		}

		permutationCount *= partSizeMap[part]
		partsPermutationCount[part] = partPermutationCount
	}

	// Create limit permutations
	limitSums := make([]int, permutationCount)
	limitPermutations := make([][]int, permutationCount)
	for i := 0; i < permutationCount; i++ {
		var limits []int
		var limitSum int

		for _, part := range formatParts {
			partSize := partSizeMap[part]
			partPermutationCount := partsPermutationCount[part]

			limit := int(math.Ceil(float64(i+1) / float64(partPermutationCount)))
			if limit > partSize {
				limit = limit % partSize
				if limit == 0 {
					limit = partSize
				}
			}

			limitSum += limit
			limits = append(limits, limit)
		}

		j := permutationCount - 1 - i
		limitSums[j] = limitSum
		limitPermutations[j] = limits
	}

	// Generate candidates
	strLength := len(str)
	var candidates []string

candidate_loop:
	for i, limits := range limitPermutations {
		limitSum := limitSums[i]
		if limitSum != strLength {
			continue
		}

		var lastIdx int
		var candidateParts []string
		mapCandidateValues := map[string]int{}
		for j, limit := range limits {
			partName := formatParts[j]
			partText := str[lastIdx : lastIdx+limit]
			partValue, err := strconv.Atoi(partText)
			if err != nil || !valueIsValid(partName, partValue) {
				continue candidate_loop
			}

			candidateParts = append(candidateParts, partText)
			mapCandidateValues[partName] = partValue
			lastIdx += limit
		}

		if candidateIsValid(mapCandidateValues) {
			candidate := strings.Join(candidateParts, "-")
			candidates = append(candidates, candidate)
		}
	}

	// Prepare Go format
	format = strings.Join(formatParts, "-")
	format = strings.ReplaceAll(format, "Y", "2006")
	format = strings.ReplaceAll(format, "y", "06")
	format = strings.ReplaceAll(format, "m", "1")
	format = strings.ReplaceAll(format, "d", "2")
	format = strings.ReplaceAll(format, "H", "15")
	format = strings.ReplaceAll(format, "M", "4")
	format = strings.ReplaceAll(format, "S", "5")

	return candidates, format
}

func valueIsValid(partName string, value int) bool {
	switch partName {
	case "Y":
		return value >= 1000
	case "y":
		return value >= 0 && value <= 99
	case "m":
		return value > 0 && value <= 12
	case "d":
		return value > 0 && value <= 31
	case "H":
		return value >= 0 && value < 24
	case "M", "S":
		return value >= 0 && value < 60
	}

	return false
}

func candidateIsValid(mapValues map[string]int) bool {
	// Get year
	year, exist := mapValues["Y"]
	if !exist {
		year, exist = mapValues["y"]
		if !exist {
			return false
		}

		currentYear := time.Now().Year()
		currentCentury := (currentYear / 1000) * 1000
		if year < currentYear-currentCentury {
			year += currentCentury
		} else {
			year += currentCentury - 100
		}
	}

	// Get month and day
	month, exist := mapValues["m"]
	if !exist {
		return false
	}

	day, exist := mapValues["d"]
	if !exist {
		return false
	}

	// Make sure day is valid
	if day < 1 || day > dateutil.GetLastDayOfMonth(year, month) {
		return false
	}

	return true
}
