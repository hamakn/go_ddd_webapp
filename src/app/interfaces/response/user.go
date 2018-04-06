package response

import (
	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
)

// UsersResponse returns response of users
func UsersResponse(users []*user.User) (*Response, error) {
	return newResponse(users)
}

// UserResponse returns response of user
func UserResponse(user *user.User) (*Response, error) {
	return newResponse(user)
}
