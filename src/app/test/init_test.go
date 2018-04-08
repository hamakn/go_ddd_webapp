package test

import (
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	"github.com/hamakn/go_ddd_webapp/src/app/internal"
)

func init() {
	internal.MockEnvironments(&environments.Environments{})
	injectDependencies()
}
