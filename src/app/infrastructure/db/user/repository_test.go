package user

import (
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	"github.com/hamakn/go_ddd_webapp/src/app/internal"
	"github.com/stretchr/testify/require"
	"google.golang.org/appengine/aetest"
)

func TestCreateFixture(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	defer done()
	require.Nil(t, err)

	r := NewRepository(ctx)
	users, err := r.CreateFixture()

	require.Nil(t, err)
	require.Equal(t, 2, len(users))
}

func init() {
	internal.MockEnvironments(&environments.Environments{})
}
