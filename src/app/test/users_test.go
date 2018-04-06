package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/application/request"
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

func TestCreateUser(t *testing.T) {
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
		PostJSON     string
		HasError     bool
		ResponseCode int
	}{
		{
			// OK json
			`{"email":"new@hamakn.test","screen_name":"new_name","age":17}`,
			false,
			200,
		},
		{
			// NG json: broken json
			"{",
			true,
			400,
		},
		{
			// NG json: no required field
			"{}",
			true,
			400,
		},
		{
			// NG json: validation failed
			`{"email":"new@hamakn.test","screen_name":"たろう","age":17}`,
			true,
			400,
		},
		{
			// NG json: email taken user
			`{"email":"foo@hamakn.test","screen_name":"new_foo","age":17}`,
			true,
			422,
		},
	}

	for _, testCase := range testCases {
		req, err := inst.NewRequest("POST", "/users/", strings.NewReader(testCase.PostJSON))
		require.Nil(t, err)

		res := httptest.NewRecorder()
		config.NewRouter().ServeHTTP(res, req)

		require.Equal(t, testCase.ResponseCode, res.Code)

		if !testCase.HasError {
			ctx := appengine.NewContext(req)

			u := &user.User{}
			json.NewDecoder(res.Body).Decode(&u)

			r := request.CreateUserRequest{}
			json.NewDecoder(strings.NewReader(testCase.PostJSON)).Decode(&r)

			dbu, err := user.NewRepository(ctx).GetByID(u.ID)
			require.Nil(t, err)

			require.Equal(t, *r.Email, dbu.Email)
			require.Equal(t, *r.ScreenName, dbu.ScreenName)
			require.Equal(t, *r.Age, dbu.Age)
		}
	}
}

func loadUserFixture(ctx context.Context) error {
	r := user.NewRepository(ctx)
	_, err := r.CreateFixture()
	return err
}
