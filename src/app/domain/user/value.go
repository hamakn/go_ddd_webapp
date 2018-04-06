package user

// CreateUserValue is request for CreateUser
type CreateUserValue struct {
	Email      *string `json:"email" validate:"required,email"`
	ScreenName *string `json:"screen_name" validate:"required,printascii,min=3,max=16"`
	Age        *int    `json:"age" validate:"required,min=0,max=120"`
}

// UpdateUserValue is request for UpdateUser
type UpdateUserValue struct {
	Email      *string `json:"email" validate:"omitempty,email"`
	ScreenName *string `json:"screen_name" validate:"omitempty,printascii"`
	Age        *int    `json:"age" validate:"omitempty,min=0,max=120"`
}

// UpdateUser updates user by update request value
func (r *UpdateUserValue) UpdateUser(u *User) bool {
	hasUpdate := false
	if r.Email != nil && u.Email != *r.Email {
		u.Email = *r.Email
		hasUpdate = true
	}
	if r.ScreenName != nil && u.ScreenName != *r.ScreenName {
		u.ScreenName = *r.ScreenName
		hasUpdate = true
	}
	if r.Age != nil && u.Age != *r.Age {
		u.Age = *r.Age
		hasUpdate = true
	}
	return hasUpdate
}
