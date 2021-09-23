package dto

type UserDto struct {
	Token string `json:"token"`
}

type BookmarkRequestDto struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
