package timezone

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPopTzOffset(t *testing.T) {
	// Helper function
	validOffset := func(str string, offsetHour float64) {
		_, offset := PopTzOffset(str)
		expected := time.Duration(offsetHour*3_600) * time.Second
		assert.Equal(t, expected, offset.Offset)
	}

	emptyOffset := func(str string) {
		_, offset := PopTzOffset(str)
		assert.True(t, offset.IsZero())
	}

	leftoverString := func(str, expectedLeftover string) {
		leftover, _ := PopTzOffset(str)
		assert.Equal(t, expectedLeftover, leftover)
	}

	// Test extracting valid offset
	validOffset("Sep 03 2014 | 4:32 pm EDT", -4)
	validOffset("17th October, 2034 @ 01:08 am PDT", -7)
	validOffset("17th October, 2034 @ 01:08 am (PDT)", -7)
	validOffset("October 17, 2014 at 7:30 am PST", -8)
	validOffset("20 Oct 2014 13:08 CET", +1)
	validOffset("20 Oct 2014 13:08cet", +1)
	validOffset("Nov 25 2014 | 10:17 pm EST", -5)
	validOffset("Nov 25 2014 | 10:17 pm +0600", +6)
	validOffset("Nov 25 2014 | 10:17 pm -0930", -9.5)
	validOffset("20 Oct 2014 | 05:17 am -1200", -12)
	validOffset("20 Oct 2014 | 05:17 am +0000", 0)
	validOffset("20 Oct 2014 | 05:17 am -0000", 0)
	validOffset("Wed Aug 05 12:00:00 EDT 2015", -4)
	validOffset("April 10, 2016 at 12:00:00 UTC", 0)
	validOffset("April 10, 2016 at 12:00:00 MEZ", 1)
	validOffset("April 10, 2016 at 12:00:00 MESZ", 2)
	validOffset("April 10, 2016 at 12:00:00 GMT+2", 2)
	validOffset("April 10, 2016 at 12:00:00 UTC+2:00", 2)
	validOffset("April 10, 2016 at 12:00:00 GMT+02:00", 2)
	validOffset("April 10, 2016 at 12:00:00 UTC+5:30", 5.5)
	validOffset("April 10, 2016 at 12:00:00 GMT+05:30", 5.5)
	validOffset("April 10, 2016 at 12:00:00 UTC-2", -2)
	validOffset("April 10, 2016 at 12:00:00 GMT-2:00", -2)
	validOffset("April 10, 2016 at 12:00:00 UTC-02:00", -2)
	validOffset("April 10, 2016 at 12:00:00 GMT-9:30", -9.5)
	validOffset("April 10, 2016 at 12:00:00 UTC-09:30", -9.5)
	validOffset("Thu, 24 Nov 2016 16:03:00 UT", 0)
	validOffset("Fri Sep 23 2016 10:34:51 GMT+0800 (CST)", 8)
	validOffset("Fri Sep 23 2016 10:34:51 GMT+12", 12)
	validOffset("Fri Sep 23 2016 10:34:51 UTC+13", 13)
	validOffset("Fri Sep 23 2016 10:34:51 GMT+1245 (CST)", 12.75)
	validOffset("Fri Sep 23 2016 10:34:51 UTC+1245", 12.75)
	validOffset("2019-07-17T12:30:00.000-03:30", -3.5)
	validOffset("2019-07-17T12:30:00.000-02:30", -2.5)
	validOffset("16. srpna 2021 9:59:44 SELČ", 2)
	validOffset("16. srpna 2021 9:59:44 SEČ", 1)
	validOffset("16. srpna 2021 9:59:44 ZEČ", 0)
	validOffset("16. srpna 2021 9:59:44 VEČ", 2)

	// Test extracting string without timezone
	emptyOffset("15 May 2004")
	emptyOffset("Wed Aug 05 12:00:00 EDTERR 2015")

	// Valid timezone must be removed from string
	leftoverString("Sep 03 2014 | 4:32 pm EDT", "Sep 03 2014 | 4:32 pm ")
	leftoverString("17th October, 2034 @ 01:08 am PDT", "17th October, 2034 @ 01:08 am ")
	leftoverString("October 17, 2014 at 7:30 am PST", "October 17, 2014 at 7:30 am ")
	leftoverString("20 Oct 2014 13:08 CET", "20 Oct 2014 13:08 ")
	leftoverString("20 Oct 2014 13:08cet", "20 Oct 2014 13:08")
	leftoverString("Nov 25 2014 | 10:17 pm EST", "Nov 25 2014 | 10:17 pm ")
	leftoverString("17th October, 2034 @ 01:08 am +0700", "17th October, 2034 @ 01:08 am ")
	leftoverString("Sep 03 2014 4:32 pm +0630", "Sep 03 2014 4:32 pm ")

	// String won't be changed if there are no valid timezone
	leftoverString("15 May 2004", "15 May 2004")
}
