package users

type CreateUserRequest struct{
		Email string `json:"email"`
		Name string `json:"name"`
}

type CreateUserResponse struct{
	Email string `json:"email"`
	Name string `json:"name"`
}

type CreateGetUserResponse struct{
	Email string `json:"email"`
	Name string `json:"name"`
}

type CreateGetUserRequest struct{
	Email string `json:"email"`
}

