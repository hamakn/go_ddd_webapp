package request

// CreateUserRequest is request for CreateUser
type CreateUserRequest struct {
	Email      *string `json:"email" validate:"required,email"`
	ScreenName *string `json:"screen_name" validate:"required,printascii,min=3,max=16"`
	Age        *int    `json:"age" validate:"required,min=0,max=120"`
}
