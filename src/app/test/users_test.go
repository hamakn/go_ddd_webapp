package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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

			r := user.CreateUserValue{}
			json.NewDecoder(strings.NewReader(testCase.PostJSON)).Decode(&r)

			dbu, err := user.NewRepository(ctx).GetByID(u.ID)
			require.Nil(t, err)

			require.Equal(t, *r.Email, dbu.Email)
			require.Equal(t, *r.ScreenName, dbu.ScreenName)
			require.Equal(t, *r.Age, dbu.Age)
		}
	}
}

func TestUpdateUser(t *testing.T) {
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
		PostJSON     string
		HasError     bool
		ResponseCode int
	}{
		{
			// NG1: bad user ID
			"bad",
			`{"email":"new@hamakn.test"}`,
			true,
			400,
		},
		{
			// NG2: broken JSON
			"1",
			`{`,
			true,
			400,
		},
		{
			// NG3: validation error (bad email)
			"1",
			`{"email":"bad_email"}`,
			true,
			400,
		},
		{
			// NG4: nothing to update
			"1",
			`{"extra":"aaa"}`,
			true,
			400,
		},
		{
			// NG5: nothing to update (same old screen_name and new screen_name)
			"1",
			`{"screen_name":"foo"}`,
			true,
			400,
		},
		{
			// NG6: no entity
			"42",
			`{"email":"new@hamakn.test"}`,
			true,
			404,
		},
		{
			// NG7: email taken
			"1",
			`{"email":"bar@hamakn.test"}`,
			true,
			422,
		},
		{
			// OK
			"1",
			`{"email":"new@hamakn.test"}`,
			false,
			200,
		},
	}

	for _, testCase := range testCases {
		req, err := inst.NewRequest("PUT", "/users/"+testCase.UserID, strings.NewReader(testCase.PostJSON))
		require.Nil(t, err)

		res := httptest.NewRecorder()
		config.NewRouter().ServeHTTP(res, req)

		require.Equal(t, testCase.ResponseCode, res.Code)

		if !testCase.HasError {
			ctx := appengine.NewContext(req)

			u := &user.User{}
			json.NewDecoder(res.Body).Decode(&u)

			r := user.UpdateUserValue{}
			json.NewDecoder(strings.NewReader(testCase.PostJSON)).Decode(&r)

			dbu, err := user.NewRepository(ctx).GetByID(u.ID)
			require.Nil(t, err)

			if r.Email != nil {
				require.Equal(t, *r.Email, dbu.Email)
			}
			if r.ScreenName != nil {
				require.Equal(t, *r.ScreenName, dbu.ScreenName)
			}
			if r.Age != nil {
				require.Equal(t, *r.Age, dbu.Age)
			}
		}
	}
}

func TestDeleteUser(t *testing.T) {
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
		{
			// NG1: bad user id
			"bad",
			true,
			400,
		},
		{
			// NG2: not found
			"42",
			true,
			404,
		},
		{
			// OK
			"1",
			false,
			204,
		},
	}

	for _, testCase := range testCases {
		req, err := inst.NewRequest("DELETE", "/users/"+testCase.UserID, nil)
		require.Nil(t, err)

		res := httptest.NewRecorder()
		config.NewRouter().ServeHTTP(res, req)

		require.Equal(t, testCase.ResponseCode, res.Code)

		if !testCase.HasError {
			ctx := appengine.NewContext(req)

			id, err := strconv.ParseInt(testCase.UserID, 10, 64)
			require.Nil(t, err)

			_, err = user.NewRepository(ctx).GetByID(id)
			require.Equal(t, user.ErrNoSuchEntity, err)
		}
	}
}

func loadUserFixture(ctx context.Context) error {
	r := user.NewRepository(ctx)
	_, err := r.CreateFixture()
	return err
}
