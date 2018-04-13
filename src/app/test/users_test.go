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
		{"1", false, http.StatusOK},
		{"42", true, http.StatusNotFound},
		{"bad", true, http.StatusBadRequest},
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
			http.StatusOK,
		},
		{
			// NG json: broken json
			"{",
			true,
			http.StatusBadRequest,
		},
		{
			// NG json: no required field
			"{}",
			true,
			http.StatusBadRequest,
		},
		{
			// NG json: validation failed
			`{"email":"new@hamakn.test","screen_name":"ｂａｄｎａｍｅ","age":17}`,
			true,
			http.StatusBadRequest,
		},
		{
			// NG json: email taken user
			`{"email":"foo@hamakn.test","screen_name":"new_foo","age":17}`,
			true,
			http.StatusUnprocessableEntity,
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

			dbu, err := user.NewRepository().GetByID(ctx, u.ID)
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
			http.StatusBadRequest,
		},
		{
			// NG2: broken JSON
			"1",
			`{`,
			true,
			http.StatusBadRequest,
		},
		{
			// NG3: validation error (bad email)
			"1",
			`{"email":"bad_email"}`,
			true,
			http.StatusBadRequest,
		},
		{
			// NG4: nothing to update
			"1",
			`{"extra":"aaa"}`,
			true,
			http.StatusBadRequest,
		},
		{
			// NG5: nothing to update (same old screen_name and new screen_name)
			"1",
			`{"screen_name":"foo"}`,
			true,
			http.StatusBadRequest,
		},
		{
			// NG6: no entity
			"42",
			`{"email":"new@hamakn.test"}`,
			true,
			http.StatusNotFound,
		},
		{
			// NG7: email taken
			"1",
			`{"email":"bar@hamakn.test"}`,
			true,
			http.StatusUnprocessableEntity,
		},
		{
			// OK
			"1",
			`{"email":"new@hamakn.test"}`,
			false,
			http.StatusOK,
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

			dbu, err := user.NewRepository().GetByID(ctx, u.ID)
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
			http.StatusBadRequest,
		},
		{
			// NG2: not found
			"42",
			true,
			http.StatusNotFound,
		},
		{
			// OK
			"1",
			false,
			http.StatusNoContent,
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

			_, err = user.NewRepository().GetByID(ctx, id)
			require.Equal(t, user.ErrNoSuchEntity, err)
		}
	}
}

func loadUserFixture(ctx context.Context) error {
	r := user.NewRepository()
	_, err := r.CreateFixture(ctx)
	return err
}
