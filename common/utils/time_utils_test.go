package utils

import (
	"log"
	"testing"
	"time"
)

func TestRangeHourTimestamp(t *testing.T) {
	start := int64(1603555200)
	end := int64(1603720822)

	res, err := RangeHourTimestamp(start, end)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(res)
}

func TestRangeDayTimestamp(t *testing.T) {
	start := int64(1603497600)
	end := int64(1604102400)

	res, err := RangeDayTimestamp(start, end)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(res)
}

func TestBeforeOneDayTime(t *testing.T) {
	res := BeforeOneDayTime()
	t.Log(res)
}

func TestDayLocalStartTimestamp(t *testing.T) {
	l, _ := time.LoadLocation("Asia/Shanghai")
	res, err := DayLocalStartTimestamp(1608566400, l)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(res)
}
