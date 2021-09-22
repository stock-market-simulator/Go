package table

type User struct {
	UserID   uint       //기본적으로 ID가 primaryKey
	Token    string     `json:"Token" gorm:"column:Token"`
	Bookmark []Bookmark `gorm:"ForeignKey:ID"`
}

type Bookmark struct {
	BookmarkID uint
	Code       string
	Name       string
}

// 테이블 이름을 따로 설정해주지 않으면 default로 구조체이름을 소문자로 만든 후 s가 붙은 테이블을 참조함
// ex) User구조체면 users의 테이블을 참조함. 따라서 밑에 TableName처럼 테이블 이름을 설정
func (User) TableName() string {
	return "user"
}

func (Bookmark) TableName() string {
	return "bookmark"
}
