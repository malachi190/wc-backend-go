package types

type RegisterPayload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FcmToken *string `json:"fcm_token,omitempty"`
}

