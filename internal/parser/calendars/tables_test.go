package calendars

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPythonJalaliConversion(t *testing.T) {
	for _, parts := range [][3]int{{1403, 1, 1}, {1403, 1, 31}, {1356, 12, 30}, {2378, 1, 1}} {
		converted, err := JalaliToGregorian(parts[0], parts[1], parts[2])
		require.NoError(t, err)
		value := time.Date(converted[0], time.Month(converted[1]), converted[2], 0, 0, 0, 0, time.UTC)
		back, err := JalaliFromGregorian(value)
		require.NoError(t, err)
		if parts == [3]int{1356, 12, 30} {
			require.Equal(t, [3]int{1357, 1, 1}, back)
		} else {
			require.Equal(t, parts, back)
		}
	}
	parts, err := JalaliToGregorian(2378, 1, 1)
	require.NoError(t, err)
	require.Equal(t, [3]int{2999, 3, 21}, parts)
	_, err = JalaliMonthLength(2378, 12)
	require.ErrorIs(t, err, errJalaliRange)
	_, err = JalaliToGregorian(2379, 1, 1)
	require.ErrorIs(t, err, errJalaliRange)
	parts, err = JalaliFromGregorian(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	require.Equal(t, [3]int{-621, 10, 11}, parts)
}

func TestPythonUmmAlQuraConversion(t *testing.T) {
	for _, test := range []struct {
		hijri     [3]int
		gregorian [3]int
	}{
		{[3]int{1343, 1, 1}, [3]int{1924, 8, 1}},
		{[3]int{1355, 1, 1}, [3]int{1936, 3, 24}},
		{[3]int{1433, 2, 30}, [3]int{2012, 1, 24}},
		{[3]int{1356, 12, 30}, [3]int{1938, 3, 2}},
	} {
		parts, err := HijriToGregorian(test.hijri[0], test.hijri[1], test.hijri[2])
		require.NoError(t, err)
		require.Equal(t, test.gregorian, parts)
	}
	for index, start := range data.Hijri.MonthStarts[:len(data.Hijri.MonthStarts)-1] {
		parts, err := HijriFromGregorian(time.Unix(int64(start)*86400, 0).UTC())
		require.NoError(t, err)
		require.Equal(t, [3]int{data.Hijri.FirstYear + index/12, index%12 + 1, 1}, parts)
		length, err := HijriMonthLength(parts[0], parts[1])
		require.NoError(t, err)
		require.Equal(t, data.Hijri.MonthStarts[index+1]-start, length)
	}
	for _, year := range []int{0, 1342, 1501, 9999} {
		_, err := HijriToGregorian(year, 1, 1)
		require.ErrorIs(t, err, errHijriRange)
	}
	_, err := HijriFromGregorian(time.Date(1924, 7, 31, 0, 0, 0, 0, time.UTC))
	require.ErrorIs(t, err, errHijriRange)
	_, err = HijriFromGregorian(time.Date(2077, 11, 17, 0, 0, 0, 0, time.UTC))
	require.ErrorIs(t, err, errHijriRange)
}
