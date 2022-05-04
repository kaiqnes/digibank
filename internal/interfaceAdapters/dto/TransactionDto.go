package dto

type CreateTransactionInput struct {
	AccountID       uint    `json:"account_id"`
	OperationTypeID uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
