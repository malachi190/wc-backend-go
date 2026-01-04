package models

type Models struct {
	User *UserModel
	WatchList *WatchList
}

func New() *Models {
	return &Models{
		User: &UserModel{},
		WatchList: &WatchList{},
	}
}
