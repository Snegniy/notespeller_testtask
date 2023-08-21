package api

//go:generate swagger generate spec -o ../public/swagger.json --scan-models

type UserLoginRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:route POST /register auth registerRequest
// Регистрация пользователя.
// responses:
//   200: registerResponse

// swagger:parameters registerRequest
type registerRequest struct {
	// in:body
	Body UserLoginRegisterRequest
}

// swagger:response registerResponse
type registerResponse struct {
	// in:body
	Body UserLoginRegisterRequest
}

// swagger:route POST /login auth loginRequest
// Авторизация пользователя.
// responses:
//   200: loginResponse

// swagger:parameters loginRequest
type loginRequest struct {
	// in:body
	Body UserLoginRegisterRequest
}

// swagger:response loginResponse
type loginResponse struct {
	// in:body
	Body UserLoginRegisterRequest
}
