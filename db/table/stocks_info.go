package table

type StockInfo struct {
	StockInfoID int    `gorm:"primary_key;column:id"`
	Code        string `gorm:"column:code"`
	Name        string `gorm:"column:name"`
}

func (StockInfo) TableName() string {
	return "stocks_info"
}
