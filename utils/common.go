package utils

import (
	"math"
	"strconv"
)

/*
	Set of utilities to Serialize and Deserialize Monetary Value.
	Serialize Amount into database
	Deserialize Amount for API Response

*/

const AmountPrecision = 1000

const TimeStampPrecision = 1000

func RoundFloat(val float64) float64 {
	return math.Round(val*AmountPrecision) / AmountPrecision
}

func SerializeAmount(amount float64) int64 {
	return int64(amount * AmountPrecision)
}

func DeserializeAmount(amount int64) float64 {
	return RoundFloat(float64(amount) / AmountPrecision)
}

func StringToInt(s string) int64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return int64(num)
}

func GetOffset(page int64, limit int64) int64 {
	return (page - 1) * limit
}
