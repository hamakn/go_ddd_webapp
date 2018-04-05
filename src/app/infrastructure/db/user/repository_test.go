package user

import (
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	"github.com/hamakn/go_ddd_webapp/src/app/internal"
	"github.com/stretchr/testify/require"
	"google.golang.org/appengine/aetest"
)

func TestCreateDup(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	defer done()
	require.Nil(t, err)

	f := user.NewFactory()
	u := f.Create("hoge@hamakn.test", "hoge", 24)
	r := NewRepository(ctx)

	// 1st time
	err = r.Create(u)
	require.Nil(t, err)

	// 2nd time
	err = r.Create(u)
	require.Contains(t, err.Error(), "Email cannot take")

	// same email
	u = f.Create("hoge@hamakn.test", "fuga", 42)
	err = r.Create(u)
	require.Contains(t, err.Error(), "Email cannot take")

	// same screen name
	u = f.Create("fuga@hamakn.test", "hoge", 99)
	err = r.Create(u)
	require.Contains(t, err.Error(), "ScreenName cannot take")

	// another user
	u = f.Create("fuga@hamakn.test", "fuga", 24)
	err = r.Create(u)
	require.Nil(t, err)
}

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
