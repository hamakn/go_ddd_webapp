package response

import (
	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
)

// GetUsersResponse returns response of users
func GetUsersResponse(users []*user.User) (*Response, error) {
	return newResponse(users)
}

// GetUserResponse returns response of user
func GetUserResponse(user *user.User) (*Response, error) {
	return newResponse(user)
}
