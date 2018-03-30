package environments

import (
	"github.com/kelseyhightower/envconfig"
)

// Environments holds application environments
type Environments struct {
	AppBaseDir string `split_words:"true"`
}

// for singleton Environments
var environments *Environments

// GetEnvironments is func to return Environments
// DI enable for test
var GetEnvironments = getEnvironmentsDefault

func getEnvironmentsDefault() *Environments {
	if environments != nil {
		return environments
	}

	environments = &Environments{}
	envconfig.Process("myapp", environments)

	return environments
}
