package test

import (
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	"github.com/hamakn/go_ddd_webapp/src/app/internal"
)

func init() {
	setTimezone()
	internal.MockEnvironments(&environments.Environments{})
	injectDependencies()
}

// same as setTimezone on config/server.go
func setTimezone() {
	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}

	time.Local = loc
}
