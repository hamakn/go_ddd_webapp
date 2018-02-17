package config

import (
	"net/http"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/config"
)

func init() {
	setTimezone()

	http.Handle("/", config.NewRouter())
}

func setTimezone() {
	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}

	time.Local = loc
}
