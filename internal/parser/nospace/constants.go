package nospace

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	rxCompatible = regexp.MustCompile(`\D+`)
	rxEightDigit = regexp.MustCompile(`^\d{8}$`)

	formatMapping = map[string]string{
		"%d": "day",
		"%m": "month",
		"%y": "year",
		"%Y": "year",
	}

	dateOrderMap = map[string]string{
		"DMY": "%d%m%y",
		"DYM": "%d%y%m",
		"MDY": "%m%d%y",
		"MYD": "%m%y%d",
		"YDM": "%y%d%m",
		"YMD": "%y%m%d",
	}

	dateFormats = []string{
		"%Y%m%d", "%Y%d%m", "%m%Y%d",
		"%m%d%Y", "%d%Y%m", "%d%m%Y",
		"%y%m%d", "%y%d%m", "%m%y%d",
		"%m%d%y", "%d%y%m", "%d%m%y",
	}

	timeFormats = []string{
		"%H%M%S",
		"%H%M%S",
		"%H%M",
		"%H",
	}

	preferredFormats = []string{
		"%Y%m%d%H%M", "%Y%m%d%H%M%S", "%Y%m%d%H%M%S.%f",
	}

	preferredFormatsOrdered8Digit = []string{
		"%m%d%Y", "%d%m%Y", "%Y%m%d", "%Y%d%m", "%m%Y%d", "%d%Y%m",
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
			"%m%d%y": createSortedFormats("%m%d%y"),
			"%m%y%d": createSortedFormats("%m%y%d"),
			"%y%m%d": createSortedFormats("%y%m%d"),
			"%y%d%m": createSortedFormats("%y%d%m"),
			"%d%m%y": createSortedFormats("%d%m%y"),
			"%d%y%m": createSortedFormats("%d%y%m"),
		}

		formats["%m%d%y"] = append(preferredFormats, formats["%m%d%y"]...)
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

func createParseCandidates(str string, format string) ([]string, string) {
	// Split format by its part
	var formatParts []string
	for _, part := range strings.Split(strings.Trim(format, "%"), "%") {
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
		for j, limit := range limits {
			partName := formatParts[j]
			partText := str[lastIdx : lastIdx+limit]
			partValue, err := strconv.Atoi(partText)
			if err != nil || !valueIsValid(partName, partValue) {
				continue candidate_loop
			}

			candidateParts = append(candidateParts, partText)
			lastIdx += limit
		}

		candidate := strings.Join(candidateParts, "-")
		candidates = append(candidates, candidate)
	}

	// Prepare new format
	format = strings.ReplaceAll(format, "%", "-%")
	format = strings.Trim(format, "-")

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
