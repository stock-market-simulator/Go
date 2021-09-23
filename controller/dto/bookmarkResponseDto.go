package dto

type BookmarkResponseDto struct {
	Name          string `json:"name"`
	CurrentPrice  int    `json:"currentPrice"`
	PreviousPrice int    `json:"previousPrice"`
}
