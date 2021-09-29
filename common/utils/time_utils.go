package utils

import (
	"fmt"
	"time"
)

func formatStr(t time.Time) (int, string, string, string, string) {
	year, month, day := t.Date()
	hour := t.Hour()
	min := t.Minute()

	mouthStr := fmt.Sprintf("%d", month)
	if month < 10 {
		mouthStr = fmt.Sprintf("0%d", month)
	}

	dayStr := fmt.Sprintf("%d", day)
	if day < 10 {
		dayStr = fmt.Sprintf("0%d", day)
	}

	hourStr := fmt.Sprintf("%d", hour)
	if hour < 10 {
		hourStr = fmt.Sprintf("0%d", hour)
	}

	minuteStr := fmt.Sprintf("%d", min)
	if min < 10 {
		minuteStr = fmt.Sprintf("0%d", min)
	}

	return year, mouthStr, dayStr, hourStr, minuteStr
}

func ParseLocationTime(t string, location *time.Location) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.ParseInLocation(layout, t, location)
}

func ParseTime(t string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.ParseInLocation(layout, t, time.Local)
}

func NowDayTime() (time.Time, error) {
	now := time.Now()
	year, mouth, day, _, _ := formatStr(now)
	nowTimeStr := fmt.Sprintf("%d-%s-%s 00:00:00", year, mouth, day)
	return ParseTime(nowTimeStr)
}

func BeforeOneDayTime() int64 {
	now := time.Now()

	// 获取当前24小时
	h := fmt.Sprintf("-%dh", 24) //减去24小时（前一天）
	dh, _ := time.ParseDuration(h)

	return now.Add(dh).Unix()
}

func DayLocalStartTimestamp(ts int64, location *time.Location) (int64, error) {
	t := time.Unix(ts, 0).In(location)
	year, mouth, day, _, _ := formatStr(t)
	nowTimeStr := fmt.Sprintf("%d-%s-%s 00:00:00", year, mouth, day)
	dayTime, err := ParseLocationTime(nowTimeStr, location)
	if err != nil {
		return 0, err
	}
	return dayTime.Unix(), nil
}

func DayStartTimestamp(ts int64) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, _, _ := formatStr(t)
	nowTimeStr := fmt.Sprintf("%d-%s-%s 00:00:00", year, mouth, day)
	dayTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return dayTime.Unix(), nil
}

func DayEndTimestamp(ts int64) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, _, _ := formatStr(t)
	nowTimeStr := fmt.Sprintf("%d-%s-%s 23:59:59", year, mouth, day)
	dayTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return dayTime.Unix(), nil
}

func RangeDayTimestamp(start, end int64) ([]int64, error) {
	startDayTimestamp, err := DayStartTimestamp(start)
	if err != nil {
		return nil, err
	}

	endDayTimestamp, err := DayStartTimestamp(end)
	if err != nil {
		return nil, err
	}

	if startDayTimestamp > endDayTimestamp {
		return nil, nil
	}

	if startDayTimestamp == endDayTimestamp {
		return []int64{startDayTimestamp}, nil
	}

	res := make([]int64, 0)
	res = append(res, startDayTimestamp)
	for {
		startDayTimestamp += 86400
		if startDayTimestamp == endDayTimestamp {
			res = append(res, startDayTimestamp)
			break
		}
		if startDayTimestamp > endDayTimestamp {
			break
		}

		res = append(res, startDayTimestamp)
	}

	return res, nil
}

func NowHourTime() (time.Time, error) {
	now := time.Now()
	year, mouth, day, hour, _ := formatStr(now)
	nowTimeStr := fmt.Sprintf("%d-%s-%s %s:00:00", year, mouth, day, hour)
	return ParseTime(nowTimeStr)
}

func HourStartTimestamp(ts int64) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, hour, _ := formatStr(t)

	nowTimeStr := fmt.Sprintf("%d-%s-%s %s:00:00", year, mouth, day, hour)
	hourTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return hourTime.Unix(), nil
}

func HourEndTimestamp(ts int64) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, hour, _ := formatStr(t)

	nowTimeStr := fmt.Sprintf("%d-%s-%s %s:59:59", year, mouth, day, hour)
	hourTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return hourTime.Unix(), nil
}

func RangeHourTimestamp(start, end int64) ([]int64, error) {
	startHourTimestamp, err := HourStartTimestamp(start)
	if err != nil {
		return nil, err
	}

	endHourTimestamp, err := HourStartTimestamp(end)
	if err != nil {
		return nil, err
	}

	if startHourTimestamp > endHourTimestamp {
		return nil, nil
	}

	if startHourTimestamp == endHourTimestamp {
		return []int64{startHourTimestamp}, nil
	}

	res := make([]int64, 0)
	res = append(res, startHourTimestamp)
	for {
		startHourTimestamp += 3600
		if startHourTimestamp == endHourTimestamp {
			res = append(res, startHourTimestamp)
			break
		}
		if startHourTimestamp > endHourTimestamp {
			break
		}

		res = append(res, startHourTimestamp)
	}

	return res, nil
}

func MinuteStartTimestamp(ts int64) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, hour, minute := formatStr(t)

	nowTimeStr := fmt.Sprintf("%d-%s-%s %s:%s:00", year, mouth, day, hour, minute)
	hourTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return hourTime.Unix(), nil
}

func CustomDayTimestamp(ts int64, format string) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, _, _ := formatStr(t)

	nowTimeStr := fmt.Sprintf(format, year, mouth, day)
	hourTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return hourTime.Unix(), nil
}

func CustomHourTimestamp(ts int64, format string) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, hour, _ := formatStr(t)

	nowTimeStr := fmt.Sprintf(format, year, mouth, day, hour)
	hourTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return hourTime.Unix(), nil
}

func CustomMinuteTimestamp(ts int64, format string) (int64, error) {
	t := time.Unix(ts, 0)
	year, mouth, day, hour, minute := formatStr(t)

	nowTimeStr := fmt.Sprintf(format, year, mouth, day, hour, minute)
	hourTime, err := ParseTime(nowTimeStr)
	if err != nil {
		return 0, err
	}
	return hourTime.Unix(), nil
}
