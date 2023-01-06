//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeStamp struct {
	hour   int
	minute int
	second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (err *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", err.msg, err.input)
}

func ParseTime(input string) (TimeStamp, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return TimeStamp{}, &TimeParseError{"Invalid number of time components", input}
	}

	hour, err := strconv.Atoi(components[0])
	if err != nil {
		return TimeStamp{}, &TimeParseError{"Error parsing hour", input}
	}

	if hour > 23 || hour < 0 {
		return TimeStamp{}, &TimeParseError{"Hour out of range", input}
	}

	minute, err := strconv.Atoi(components[1])
	if err != nil {
		return TimeStamp{}, &TimeParseError{"Error parsing minute", input}
	}

	if minute > 59 || minute < 0 {
		return TimeStamp{}, &TimeParseError{"Minute out of range", input}
	}

	second, err := strconv.Atoi(components[2])
	if err != nil {
		return TimeStamp{}, &TimeParseError{"Error parsing second", input}
	}

	if second > 59 || second < 0 {
		return TimeStamp{}, &TimeParseError{"Second out of range", input}
	}

	return TimeStamp{hour, minute, second}, nil
}
