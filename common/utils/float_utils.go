package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"strconv"
)

func FormatFloatByString(fs string) string {
	f, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return fs
	}

	return fmt.Sprint(FormatFloat(f))
}

func FormatFloatByStringToF(fs string) (float64, error) {
	f, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return 0, err
	}

	return FormatFloat(f), nil
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

// FormatFloat 保留指定位数有效数字
func FormatFloat(f float64, digits ...int) float64 {
	// 默认保留4为有效数字
	curDigits := 4
	if len(digits) > 0 {
		curDigits = digits[0]
	}

	if f == 0 {
		return 0
	}

	if f == 1 {
		return 1
	}

	fr, err := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", curDigits), f), 64)
	if err != nil {
		log.WithField("f", f).WithError(err).Error("parse float64 error")
		return f
	}
	return fr
}

func FloatToBigInt(val float64, decimalCount int) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(int64(math.Pow10(decimalCount))))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

func ComposeMax(a, b string) string {
	first, _ := strconv.ParseFloat(a, 64)
	second, _ := strconv.ParseFloat(b, 64)
	max := math.Max(first, second)
	out := strconv.FormatFloat(max, 'f', -1, 64)
	return out
}

func ComposeMin(a, b string) string {
	first, _ := strconv.ParseFloat(a, 64)
	second, _ := strconv.ParseFloat(b, 64)
	min := math.Min(first, second)
	out := strconv.FormatFloat(min, 'f', -1, 64)
	return out
}

func StringAdd(a, b string) string {
	first, _ := strconv.ParseFloat(a, 64)
	second, _ := strconv.ParseFloat(b, 64)
	addition := first + second
	out := strconv.FormatFloat(addition, 'f', -1, 64)
	return out
}
