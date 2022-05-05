package dto

type AccountPresenterResponse struct {
	AccountID      uint   `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type CreateAccountInput struct {
	DocumentNumber string `json:"document_number"`
}
