package responses

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func GetErrorUserResponse(statusCode int, err error) *UserResponse {
	return &UserResponse{
		Status:  statusCode,
		Message: "error",
		Data:    map[string]interface{}{"data": err.Error()},
	}
}
