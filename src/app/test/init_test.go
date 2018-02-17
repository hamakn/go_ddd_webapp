package test

import "time"

func init() {
	setTimezone()
}

// same as setTimezone on config/server.go
func setTimezone() {
	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}

	time.Local = loc
}
