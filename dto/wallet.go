package dto

// WalletRequest ...
type WalletRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

// BalanceResponse ...
type BalanceResponse struct {
	Balance float64 `json:"balance"`
}