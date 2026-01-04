package types

type RegisterPayload struct {
	Email       string  `json:"email" binding:"required,email"`
	DisplayName string  `json:"display_name" binding:"required,max=20"`
	Username    string  `json:"username" binding:"required"`
	Password    string  `json:"password" binding:"required,min=8"`
	FcmToken    *string `json:"fcm_token,omitempty"`
}

type VerifyEmailPayload struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,max=4"`
}

type ResendEmailPayload struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
