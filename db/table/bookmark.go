package table

type Bookmark struct {
	BookmarkID int `gorm:"primary_key"`
	Code       string
	Name       string
	UserID     int
}

func (Bookmark) TableName() string {
	return "bookmark"
}
