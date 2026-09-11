package Build

import (
	"fmt"
	"strconv"
)

// ConvertToInt converts a string value to an integer and returns 0 if the
// conversion fails.
func ConvertToInt(val string) int {
	c, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return c
}

// ConvertToFloat converts a string value to a float64 and returns 0 if the
// conversion fails.
func ConvertToFloat(val string) float64 {
	c, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Println(err)
		return 0.0
	}
	return c
}
