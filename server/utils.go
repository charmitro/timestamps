package server

import (
	"regexp"
	"time"
)

// Validation via simple regexp to validate 1{h,d,mo,y}
func validatePeriod(period string) bool {
	matched, _ := regexp.MatchString(`^(\d{1,2}[h,d,m,y])`, period)
	return matched
}

// Validates timezone via time.LoadLocation(name string).
// NOTE: Not sure abou this, no time to fix.
func validateTimezone(tz string) (*time.Location, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	time.LoadLocation(tz)

	return loc, nil
}

const layout string = "20060102T150405Z"

// Parse each URL.Query() param and populate `properties` struct object
// if parameter is valid.
func parseQuery(properties *Properties, key, param string) *ErrorResponse {
	switch key {
	case "period":
		if err := validatePeriod(param); !err {
			return &ErrorResponse{
				Status: "error",
				Desc:   "Unsupported period",
			}
		}
		properties.Period = param
	case "tz":
		tz, err := validateTimezone(param)
		if err != nil {
			return &ErrorResponse{
				Status: "error",
				Desc:   "Unsupported timezone",
			}
		}

		properties.Tz = *tz

	case "t1":
		t1, err := time.Parse(layout, param)
		if err != nil {
			return &ErrorResponse{
				Status: "error",
				Desc:   "Unsupported time interval",
			}
		}

		properties.T1 = t1

	case "t2":
		t2, err := time.Parse(layout, param)
		if err != nil {
			return &ErrorResponse{
				Status: "error",
				Desc:   "Unsupported time interval",
			}
		}
		properties.T2 = t2
	}

	return nil
}
