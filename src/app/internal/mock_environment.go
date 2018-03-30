package internal

import (
	"os"
	"path/filepath"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
)

const (
	// AppRootFile is the file on the app root
	AppRootFile = "Gopkg.lock"
)

// MockEnvironments is mocking environments.GetEnvironments
func MockEnvironments(env *environments.Environments) {
	environments.GetEnvironments = func() *environments.Environments {
		env.AppBaseDir = appBaseDir()
		return env
	}
}

func appBaseDir() string {
	dir := filepath.Join(".")
	for {
		if isFileExist(filepath.Join(dir, AppRootFile)) {
			return filepath.Join(dir, "config")
		}
		dir = filepath.Join(dir, "..")
	}
}

// isFileExist returns file exist or not
// refs: https://qiita.com/hnakamur/items/848097aad846d40ae84b
func isFileExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
