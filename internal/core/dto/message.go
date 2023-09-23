package dto

type SendMessageRequestDto struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

type SendMessageResponseDto struct {
	TransactionHash string `json:"transactionHash"`
	GasUsed         string `json:"gasUsed"`
}
