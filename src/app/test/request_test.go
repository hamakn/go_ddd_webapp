package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"

	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/config"
	"github.com/stretchr/testify/require"
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

	res := httptest.NewRecorder()
	config.NewRouter().ServeHTTP(res, req)
	require.Equal(t, http.StatusOK, res.Code)

	users := &[]*user.User{}
	err = json.NewDecoder(res.Body).Decode(users)
	require.Nil(t, err)

	require.Equal(t, 42, int((*users)[0].ID))
}

func TestGetUser(t *testing.T) {
	inst, err := aetest.NewInstance(&aetest.Options{
		StronglyConsistentDatastore: true,
	})
	require.Nil(t, err)
	defer inst.Close()

	key := "42"
	req, err := inst.NewRequest("GET", "/users/"+key, nil)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	config.NewRouter().ServeHTTP(res, req)
	require.Equal(t, http.StatusNotFound, res.Code)

	require.Equal(t, res.Body.String(), "{\"error\":\"key: "+key+" not found\"}")
}
