package dto

type SingleTransferDto struct {
	// FromAddress string `json:"from_address"`
	ToAddress string `json:"to_address"`
	Amount    string `json:"amount"`
}

type MultiTransfer []SingleTransferDto
