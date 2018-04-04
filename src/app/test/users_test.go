package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/config"
	"github.com/stretchr/testify/require"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestGetUsers(t *testing.T) {
	inst, err := aetest.NewInstance(&aetest.Options{
		StronglyConsistentDatastore: true,
	})
	require.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/users/", nil)
	require.Nil(t, err)

	err = loadUserFixture(appengine.NewContext(req))
	require.Nil(t, err)

	res := httptest.NewRecorder()
	config.NewRouter().ServeHTTP(res, req)
	require.Equal(t, http.StatusOK, res.Code)

	users := []*user.User{}
	err = json.NewDecoder(res.Body).Decode(&users)
	require.Nil(t, err)

	require.Equal(t, 2, len(users))
}

func TestGetUser(t *testing.T) {
	inst, err := aetest.NewInstance(&aetest.Options{
		StronglyConsistentDatastore: true,
	})
	require.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/dummy", nil)
	require.Nil(t, err)

	err = loadUserFixture(appengine.NewContext(req))
	require.Nil(t, err)

	testCases := []struct {
		UserID       string
		HasError     bool
		ResponseCode int
	}{
		{"1", false, 200},
		{"42", true, 404},
		{"bad", true, 400},
	}

	for _, testCase := range testCases {
		req, err := inst.NewRequest("GET", "/users/"+testCase.UserID, nil)
		require.Nil(t, err)

		res := httptest.NewRecorder()
		config.NewRouter().ServeHTTP(res, req)

		require.Equal(t, testCase.ResponseCode, res.Code)
		if !testCase.HasError {
			u := &user.User{}
			json.NewDecoder(res.Body).Decode(u)
			require.Equal(t, testCase.UserID, fmt.Sprint(u.ID))
		}
	}
}

func loadUserFixture(ctx context.Context) error {
	r := user.NewRepository(ctx)
	_, err := r.CreateFixture()
	return err
}
