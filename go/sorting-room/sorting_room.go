package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	i := strconv.FormatFloat(f, 'f', 1, 64)
	return fmt.Sprintf("This is the number %s", i)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	floatNum, err := strconv.ParseFloat(strconv.Itoa(nb.Number()), 64)
	if err != nil {
		panic("Couldn't parse the int")
	}
	outputStr := strconv.FormatFloat(floatNum, 'f', 1, 64)
	return fmt.Sprintf("This is a box containing the number %s", outputStr)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		i, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return 0
		}
		return i
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
	case FancyNumber:
		floatNum, err := strconv.ParseFloat(fnb.Value(), 64)
		if err != nil {
			panic("Couldn't parse the string to float")
		}
		outputStr := strconv.FormatFloat(floatNum, 'f', 1, 64)
		return fmt.Sprintf("This is a fancy box containing the number %s", outputStr)
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch val := i.(type) {
	case int:
		return DescribeNumber(float64(val))
	case float64:
		return DescribeNumber(val)
	case NumberBox:
		return DescribeNumberBox(val)
	case FancyNumberBox:
		return DescribeFancyNumberBox(val)
	default:
		return "Return to sender"
	}
	panic("Please implement DescribeAnything")
}
