package models

// Wallet ...
type Wallet struct {
	WalletID int `gorm:"column:id; primaryKey; size:10" json:"wallet_id"`
	Balance  float64 `gorm:"column:balance" json:"balance"`
}
