package config

import (
	"net/http"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/config"
)

func init() {
	injectDependencies()

	http.Handle("/", config.NewRouter())
}
