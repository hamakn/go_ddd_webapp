package internal

import (
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	"github.com/stretchr/testify/require"
)

func TestMockEnvironments(t *testing.T) {
	MockEnvironments(&environments.Environments{})
	env := environments.GetEnvironments()

	require.Equal(t, env.AppBaseDir, "../../../config")
}
