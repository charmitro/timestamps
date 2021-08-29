package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Logic of calculating the periodic tasks
func calculateIntervals(properties *Properties) []string {
	t1 := properties.T1
	t2 := properties.T2

	fmt.Printf("t1.Location(): %v\n", t1.Location())
	var ptlist []string

	switch properties.Period {
	case "1h":
		minutes := 59 - t1.Minute()
		t1 = t1.Add(time.Duration(minutes) * time.Minute)

		seconds := 60 - t1.Second()
		t1 = t1.Add(time.Duration(seconds) * time.Second)

		ptlist = append(ptlist, t1.Format(layout))
		for t1.Before(t2) {
			t1 = t1.Add(time.Duration(time.Hour))
			ptlist = append(ptlist, t1.Format(layout))
		}

	case "1d":
		minutes := 59 - t1.Minute()
		t1 = t1.Add(time.Duration(minutes) * time.Minute)

		seconds := 60 - t1.Second()
		t1 = t1.Add(time.Duration(seconds) * time.Second)

		ptlist = append(ptlist, t1.Format(layout))
		for t1.Before(t2) {
			t1 = t1.Add(time.Duration(24 * time.Hour))
			ptlist = append(ptlist, t1.Format(layout))
		}

	case "1mo":
		minutes := 59 - t1.Minute()
		t1 = t1.Add(time.Duration(minutes) * time.Minute)

		seconds := 60 - t1.Second()
		t1 = t1.Add(time.Duration(seconds) * time.Second)

		ptlist = append(ptlist, t1.Format(layout))
		for t1.Before(t2) {
			t1 = t1.AddDate(0, 1, 0)
			ptlist = append(ptlist, t1.Format(layout))
		}

	case "1y":
		minutes := 59 - t1.Minute()
		t1 = t1.Add(time.Duration(minutes) * time.Minute)

		seconds := 60 - t1.Second()
		t1 = t1.Add(time.Duration(seconds) * time.Second)

		t1 := time.Date(t1.Year(), 13, 0, 22, t1.Minute(), t1.Second(), t1.Nanosecond(), t1.Location())

		ptlist = append(ptlist, t1.Format(layout))
		for t1.Before(t2) {
			t1 = t1.AddDate(1, 0, 0)
			ptlist = append(ptlist, t1.Format(layout))
		}
	}
	return ptlist
}

func ptlist(c *gin.Context) {
	properties := &Properties{}

	for key, param := range c.Request.URL.Query() {
		if err := parseQuery(properties, key, param[0]); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	}

	if err := properties.validateProperties(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	tt := calculateIntervals(properties)

	c.JSON(http.StatusOK, tt)
}

func Server(port string) {
	r := gin.Default()

	r.GET("/ptlist", ptlist)

	r.Run(":" + port)
}
