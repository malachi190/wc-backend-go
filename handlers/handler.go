package handlers

type Handler struct {
	Auth *AuthHandler
}

type AuthHandler struct{}

func New() *Handler {
	return &Handler{
		Auth: &AuthHandler{},
	}
}
