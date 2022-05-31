package model

type SignInResponse struct {
	Username    string `json:"username,omitempty"`
	Role        string `json:"role,omitempty"`
	TokenString string `json:"tokenstring,omitempty"`
}

func NewSignInResponse(username string, role string, tokenstring string) *SignInResponse {
	r := &SignInResponse{username, role, tokenstring}
	return r
}

type ValidateRequest struct {
	// Username    string `json:"username,omitempty"`
	TokenString string `json:"tokenstring,omitempty"`
}

func NewValidateRequest(tokenstring string) *ValidateRequest {
	r := &ValidateRequest{tokenstring}
	return r
}

type ValidateResponse struct {
	Username string `json:"username,omitempty"`
	Role     string `json:"role,omitempty"`
}

func NewValidateResponse(username string, tokenstring string) *ValidateResponse {
	r := &ValidateResponse{username, tokenstring}
	return r
}

type RefreshRequest struct {
	// Username    string `json:"username,omitempty"`
	TokenString string `json:"tokenstring,omitempty"`
}

type RefreshResponse struct {
	// Username    string `json:"username,omitempty"`
	TokenString string `json:"tokenstring,omitempty"`
}

func NewRefreshResponse(tokenstring string) *RefreshResponse {
	r := &RefreshResponse{tokenstring}
	return r
}

type ValidateUserPasswordRequest struct {
	Uname    string `json:"uname"`
	Password string `json:"password"`
}

func NewValidateUserPasswordRequest(uname string, password string) *ValidateUserPasswordRequest {
	r := &ValidateUserPasswordRequest{uname, password}
	return r
}

type ValidatePasswordResponse struct {
	Uname     string `json:"uname"`
	User_role string `json:"user_role"`
}

func NewValidatePasswordResponse(uname string, user_role string) *ValidatePasswordResponse {
	r := &ValidatePasswordResponse{uname, user_role}
	return r
}
