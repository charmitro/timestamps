package server

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ErrorResponse is a basic JSON
// structure according to the example.
type ErrorResponse struct {
	Status string `json:"status,omitempty"`
	Desc   string `json:"desc,omitempty"`
}

// Properties store params `period, tz, t1, t2`
type Properties struct {
	Period string
	Tz     time.Location
	T1     time.Time
	T2     time.Time
}

// We iterate through the values of p Properties object
// and find if IsZero() https://pkg.go.dev/reflect#Value.IsZero
// and if yes, return an ErrorResponse saying that the parameter
// missing
func (p Properties) validateProperties() *ErrorResponse {
	v := reflect.ValueOf(p)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			return &ErrorResponse{
				Status: "error",
				Desc:   fmt.Sprintf("Parameter '%s' is missing.", strings.ToLower(typeOfS.Field(i).Name)),
			}

		}
	}

	return nil
}
