package user

import (
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/environments"
	"github.com/hamakn/go_ddd_webapp/src/app/internal"
	"github.com/stretchr/testify/require"
	"google.golang.org/appengine/aetest"
)

func TestCreate(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	defer done()
	require.Nil(t, err)

	f := user.NewFactory()
	r := NewRepository(ctx)
	_, err = r.CreateFixture()
	require.Nil(t, err)

	testCases := []struct {
		email      string
		screenName string
		age        int
		hasError   bool
		err        error
	}{
		{
			// taken email and screen name
			"foo@hamakn.test",
			"foo",
			24,
			true,
			user.ErrEmailCannotTake,
		},
		{
			// taken email
			"foo@hamakn.test",
			"new_name",
			25,
			true,
			user.ErrEmailCannotTake,
		},
		{
			// taken screen name
			"new@hamakn.test",
			"foo",
			26,
			true,
			user.ErrScreenNameCannotTake,
		},
		{
			// ok
			"new@hamakn.test",
			"new",
			17,
			false,
			nil,
		},
	}

	for _, testCase := range testCases {
		u := f.Create(testCase.email, testCase.screenName, testCase.age)
		err := r.Create(u)

		if testCase.hasError {
			require.NotNil(t, err)
			require.Equal(t, testCase.err, err)

		} else {
			require.Nil(t, err)
		}
	}
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
