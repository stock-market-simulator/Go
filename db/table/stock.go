package table

type Stock struct {
	StockID           int `gorm:"primary_key;column:id"`
	TradingVolume     int `gorm:"column:trading_volume"`
	TransactionAmount int `gorm:"column:transaction_amount"`
	Date              int `gorm:"column:data"`
	MarketPrice       int `gorm:"column:market_price"`
	HighPrice         int `gorm:"column:high_price"`
	LowPrice          int `gorm:"column:low_price"`
}
