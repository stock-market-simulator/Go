package dto

type UserDto struct {
	Token string `json:"token"`
}

type BookmarkDto struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
