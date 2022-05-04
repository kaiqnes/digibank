package dto

type AccountPresenterResponse struct {
	AccountID      uint
	DocumentNumber string `json:"document_number"`
}

type CreateAccountInput struct {
	DocumentNumber string `json:"document_number"`
}
