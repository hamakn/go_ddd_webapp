package request

// CreateUserRequest is request for CreateUser
type CreateUserRequest struct {
	Email      *string `json:"email" validate:"required,email"`
	ScreenName *string `json:"screen_name" validate:"required,printascii,min=3,max=16"`
	Age        *int    `json:"age" validate:"required,min=0,max=120"`
}

// UpdateUserRequest is request for UpdateUser
type UpdateUserRequest struct {
	Email      *string `json:"email" validate:"omitempty,email"`
	ScreenName *string `json:"screen_name" validate:"omitempty,printascii"`
	Age        *int    `json:"age" validate:"omitempty,min=0,max=120"`
}

// IsEmpty returns any value is present or not
func (r *UpdateUserRequest) IsEmpty() bool {
	if r.Email == nil && r.ScreenName == nil && r.Age == nil {
		return true
	}
	return false
}
