package fixture

import (
	"os"
	"path/filepath"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	yaml "gopkg.in/yaml.v2"
)

// Load loads fixtures to dst, fixture name is specified by fixtureName
func Load(fixtureName string, dst interface{}) error {
	f, err := os.Open(fixtureFilepath(fixtureName))
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func fixtureFilepath(fixtureName string) string {
	return filepath.Join(environments.GetEnvironments().AppBaseDir, "fixtures", fixtureName+".yml")
}
