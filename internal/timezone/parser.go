package timezone

// PopTzOffset extracts timezone data from str then return the
// str with the timezone code removed.
func PopTzOffset(str string) (string, TimezoneOffsetData) {
	if rxSearchIgnoreCase.MatchString(str) {
		for _, offsetData := range timezoneOffsets {
			idxs := offsetData.Regex.FindStringIndex(str)
			if len(idxs) == 0 {
				continue
			}

			str = str[:idxs[0]+1] + str[idxs[1]:]
			return str, offsetData
		}
	}

	return str, TimezoneOffsetData{}
}

// WordIsTz check if the specified word is a timezone code.
func WordIsTz(word string) bool {
	return rxSearch.MatchString(word)
}
