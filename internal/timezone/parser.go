package timezone

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

func WordIsTz(word string) bool {
	return rxSearch.MatchString(word)
}
