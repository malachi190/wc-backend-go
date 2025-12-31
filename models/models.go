package models

type Models struct {
	User *UserModel
}

func New() *Models {
	return &Models{
		User: &UserModel{},
	}
}
