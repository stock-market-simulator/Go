package dto

type BookmarkResponseDto struct {
	BookMark []BookmarkInfo `json:"bookmark"`
}

type BookmarkInfo struct {
	Name          string `json:"name"`
	CurrentPrice  int    `json:"currentPrice"`
	PreviousPrice int    `json:"previousPrice"`
}
