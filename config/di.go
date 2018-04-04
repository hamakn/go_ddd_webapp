package config

import (
	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	dbUser "github.com/hamakn/go_ddd_webapp/src/app/infrastructure/db/user"
)

func injectDependencies() {
	injectRepositoryDependencies()
}

func injectRepositoryDependencies() {
	user.NewRepository = dbUser.NewRepository
}
