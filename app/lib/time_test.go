package lib

import (
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func TestCurrentTime(t *testing.T) {
	timeWithFormat := CurrentTime("2006-01-02 15:04:05")
	timeDefault := CurrentTime("")
	utils.AssertEqual(t, timeWithFormat, timeDefault)
}

func TestTimeNow(t *testing.T) {
	TimeNow()
}

func TestRangeDate(t *testing.T) {
	RangeDate("2022-01-02 15:04:05", "2006-01-05 15:04:05", "days")
	RangeDate("2022-01-02 15:04:05", "2006-01-05 15:04:05", "hours")
	RangeDate("2022-01-02 15:04:05", "2006-01-05 15:04:05", "nanoseconds")
	RangeDate("2022-01-02 15:04:05", "2006-01-05 15:04:05", "minutes")
	RangeDate("2022-01-02 15:04:05", "2006-01-05 15:04:05", "seconds")
}

func TestDateTimeAhead(t *testing.T) {
	DateTimeAhead("2022-01-02 15:04:05", "2022-01-02 15:04:05", 0, 1, 1)
	DateTimeAhead("2022-01-02 15:04:05", "", 0, 1, 1)
}
